package submitter

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"errors"
	"fmt"
	"infinity/pool/internal"
	"infinity/pool/internal/contracts/PoW"
	"infinity/pool/internal/contracts/PoolRegistry"
	"infinity/pool/internal/utils"
	"log"
	"log/slog"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type Submitter struct {
	PoolId uint64
	pdb    *sql.DB

	// submitter info
	Address     common.Address
	nonce       uint64
	privateKey  *ecdsa.PrivateKey
	chainClient *ethclient.Client
	mu          sync.RWMutex

	chainId *big.Int
	// submit contract info
	powAddress  common.Address
	pow         PoW.PoW
	powInstance *bind.BoundContract

	// claim contract info
	poolRegistryAddress  common.Address
	poolRegistry         PoolRegistry.PoolRegistry
	poolRegistryInstance *bind.BoundContract
	poolRegistryDomain   apitypes.TypedDataDomain
}

const gasLimit = uint64(1_000_000)
const createTables = `
CREATE TABLE IF NOT EXISTS submits (
    tx_hash VARCHAR(66) PRIMARY KEY,
    tx_cost NUMERIC(78, 0) NOT NULL,
    reward NUMERIC(78, 0) NOT NULL,
    block BIGINT NOT NULL,
    finalized BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS idx_submits_block
ON submits (block);

CREATE INDEX IF NOT EXISTS idx_submits_finalized
ON submits (finalized);

CREATE TABLE IF NOT EXISTS rewards (
    miner VARCHAR(42) PRIMARY KEY,
    amount       NUMERIC(78, 0) NOT NULL DEFAULT 0,
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_rewards_miner
ON submits (block);
`
const insertSubmit = `
	INSERT INTO submits (tx_hash, tx_cost, reward, block)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (tx_hash) DO NOTHING
`

const unfinalizedStats = `
SELECT
	COALESCE(SUM(tx_cost), 0) cost,
	COUNT(*) count
FROM submits
WHERE finalized = FALSE;
`

const finalizeTxs = `
UPDATE submits SET finalized = TRUE WHERE finalized = FALSE;
`

func NewSubmitter(chainClient *ethclient.Client, pdb *sql.DB, poolPrivateKey []byte) *Submitter {
	_, err := pdb.Exec(createTables)
	if err != nil {
		log.Fatalf("Can't init db, %v", err)
	}

	chainId, err := chainClient.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Can't get chain id, %v", err)
	}

	privateKey, err := crypto.ToECDSA(poolPrivateKey)
	if err != nil {
		log.Fatalf("Can't parse pool private key, %v", err)
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := chainClient.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatalf("Can't get pool nonce, %v", err)
	}

	pow := *PoW.NewPoW()
	powAddress := common.HexToAddress(internal.PoWAddress)
	powInstance := pow.Instance(chainClient, powAddress)

	poolRegistry := *PoolRegistry.NewPoolRegistry()
	poolRegistryAddress, err := bind.Call(powInstance, nil, pow.PackPoolRegistry(), pow.UnpackPoolRegistry)
	if err != nil {
		log.Fatalf("Can't get pool registy address, %v", err)
	}
	poolRegistryInstance := poolRegistry.Instance(chainClient, poolRegistryAddress)

	poolId, err := bind.Call(poolRegistryInstance, nil, poolRegistry.PackGetPoolId(address), poolRegistry.UnpackGetPoolId)
	if err != nil {
		log.Fatalf("Can't get pool id, %v", err)
	}

	poolRegistryDomain, err := bind.Call(poolRegistryInstance, nil, poolRegistry.PackEip712Domain(), poolRegistry.UnpackEip712Domain)
	if err != nil {
		log.Fatalf("Can't get pool registry domain, %v", err)
	}

	submitter := Submitter{
		PoolId:               poolId.Uint64(),
		pdb:                  pdb,
		Address:              address,
		nonce:                nonce,
		privateKey:           privateKey,
		chainClient:          chainClient,
		chainId:              chainId,
		powAddress:           powAddress,
		poolRegistryAddress:  poolRegistryAddress,
		poolRegistry:         poolRegistry,
		poolRegistryInstance: poolRegistryInstance,
		pow:                  pow,
		powInstance:          powInstance,
		poolRegistryDomain: apitypes.TypedDataDomain{
			Name:              poolRegistryDomain.Name,
			Version:           poolRegistryDomain.Version,
			ChainId:           (*math.HexOrDecimal256)(poolRegistryDomain.ChainId),
			VerifyingContract: poolRegistryDomain.VerifyingContract.Hex(),
		},
	}

	go func() {
		queryTicker := time.NewTicker(time.Second)
		defer queryTicker.Stop()

		for range queryTicker.C {
			submitter.nonce, _ = chainClient.PendingNonceAt(context.Background(), address)
		}
	}()

	return &submitter
}

func (s *Submitter) GetBalance() (*big.Int, error) {
	return s.chainClient.PendingBalanceAt(context.Background(), s.Address)
}

func waitForTransactionReceipt(ctx context.Context, c *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second / 10)
	defer queryTicker.Stop()

	for {
		receipt, err := c.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func (s *Submitter) Submit(privateKeyB *ecdsa.PrivateKey, privateKeyAB *ecdsa.PrivateKey) error {
	gasPrice, err := s.chainClient.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	data := common.Hex2Bytes(internal.Data)
	var packed []byte
	packed = append(packed, s.poolRegistryAddress.Bytes()...)
	packed = append(packed, data...)
	digest := crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), crypto.Keccak256(packed))
	signature, err := crypto.Sign(digest, privateKeyAB)
	if err != nil {
		return err
	}
	signature[crypto.RecoveryIDOffset] += 27

	s.mu.Lock()
	tx, err := s.powInstance.Transact(
		&bind.TransactOpts{
			Nonce: big.NewInt(int64(s.nonce)),
			Signer: func(_ common.Address, t *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(t, types.NewEIP155Signer(s.chainId), s.privateKey)
			},
			GasLimit: gasLimit,
			GasPrice: gasPrice,
		},
		"poolSubmit",
		PoW.ECCPoint{X: privateKeyB.PublicKey.X, Y: privateKeyB.PublicKey.Y},
		signature,
		data,
	)
	if err != nil {
		s.mu.Unlock()
		return err
	}

	s.nonce++
	slog.Debug("Submit transaction sended", "tx", tx.Hash())
	s.mu.Unlock()

	receipt, err := waitForTransactionReceipt(context.Background(), s.chainClient, tx.Hash())
	if err != nil {
		return err
	}

	txCost := new(big.Int)
	txCost.Mul(receipt.EffectiveGasPrice, big.NewInt(int64(receipt.GasUsed)))

	var reward *big.Int
	if receipt.Status == types.ReceiptStatusSuccessful {
		for _, log := range receipt.Logs {
			if log.Address != s.powAddress {
				continue
			}
			submit, _ := s.pow.UnpackSubmissionEvent(log)
			if submit != nil {
				reward = submit.Reward
				break
			}
		}
	} else {
		reward = common.Big0
	}
	_, err = s.pdb.Exec(insertSubmit, tx.Hash().Hex(), txCost.String(), reward.String(), receipt.BlockNumber.String())
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (s *Submitter) FinalizeRewards(shares map[string]*big.Int) error {
	if len(shares) == 0 {
		slog.Debug("Skip FinalizeRewards - nothing to submit")
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	var costStr string
	var count int

	err := s.pdb.QueryRow(unfinalizedStats).Scan(&costStr, &count)
	if err != nil {
		return err
	}
	if count == 0 {
		slog.Debug("Skip FinalizeRewards - nothing to submit")
		return nil
	}

	auth := bind.NewKeyedTransactor(s.privateKey, s.chainId)
	tx, err := bind.Transact(s.poolRegistryInstance, auth, s.poolRegistry.PackFinalizeReward(big.NewInt(0)))
	if err != nil {
		return err
	}
	s.nonce++

	receipt, err := waitForTransactionReceipt(context.Background(), s.chainClient, tx.Hash())
	if err != nil {
		return err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("finalize rewards transaction failed")
	}

	s.pdb.Query(finalizeTxs)
	minersReward := big.NewInt(0)
	for _, log := range receipt.Logs {
		rewardsFinalized, _ := s.poolRegistry.UnpackRewardsFinalizedEvent(log)
		if rewardsFinalized != nil {
			minersReward.Add(minersReward, rewardsFinalized.Reward)
			minersReward.Sub(minersReward, rewardsFinalized.PoolFee)
			minersReward.Sub(minersReward, rewardsFinalized.SubmitsCost)
			break
		}
	}

	totalShare := big.NewInt(0)
	for _, share := range shares {
		totalShare.Add(totalShare, share)
	}

	type MinerReward struct {
		minerAddress string
		amount       *big.Int
	}

	rewards := []MinerReward{}
	for minerAddress, share := range shares {
		reward := new(big.Int)
		reward.Mul(minersReward, share)
		reward.Div(reward, totalShare)
		rewards = append(rewards, MinerReward{minerAddress: minerAddress, amount: reward})
	}

	var query strings.Builder
	query.WriteString("INSERT INTO rewards (miner, amount, updated_at) VALUES ")
	args := []any{}

	for i, r := range rewards {
		query.WriteString(fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
		if (i + 1) != len(rewards) {
			query.WriteString(", ")
		}
		args = append(args, r.minerAddress, r.amount.String(), time.Now())
	}
	query.WriteString(" ON CONFLICT (miner) DO UPDATE SET amount = rewards.amount + EXCLUDED.amount, updated_at = EXCLUDED.updated_at;")
	s.pdb.Exec(query.String(), args...)

	slog.Debug("FinalizeRewards finished", "submitsCount", count, "minersReward", utils.WeiToEther(minersReward).String())
	return nil
}

func (s *Submitter) GetClaimInfo(minerAddress common.Address) (*big.Int, []byte, error) {
	var totalRewardRaw string
	totalReward := new(big.Int)
	err := s.pdb.QueryRow("SELECT amount FROM rewards WHERE miner = $1", minerAddress.Hex()).Scan(&totalRewardRaw)
	if err != nil {
		if err == sql.ErrNoRows {
			totalReward.SetInt64(0)
		} else {
			return nil, nil, err
		}
	} else {
		totalReward.SetString(totalRewardRaw, 10)
	}

	digest, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Claim": {
				{Name: "poolId", Type: "uint256"},
				{Name: "miner", Type: "address"},
				{Name: "totalReward", Type: "uint256"},
			},
		},
		PrimaryType: "Claim",
		Domain:      s.poolRegistryDomain,
		Message: apitypes.TypedDataMessage{
			"poolId":      big.NewInt(int64(s.PoolId)),
			"miner":       minerAddress.Hex(),
			"totalReward": totalReward,
		},
	})
	if err != nil {
		return nil, nil, err
	}

	signature, err := crypto.Sign(digest[:], s.privateKey)
	if err != nil {
		return nil, nil, err
	}
	signature[crypto.RecoveryIDOffset] += 27

	return totalReward, signature, nil
}
