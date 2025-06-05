package submitter

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"infinity/pool/internal"
	"infinity/pool/internal/contracts/PoW"
	"infinity/pool/internal/contracts/PoolRegistry"
	"infinity/pool/internal/utils"
	"log"
	"log/slog"
	"math/big"
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
	powInstance bind.BoundContract

	// claim contract info
	poolRegistryAddress common.Address
	poolRegistryDomain  apitypes.TypedDataDomain
}

const gasLimit = uint64(1_000_000)
const createSubmitsTable = `
CREATE TABLE IF NOT EXISTS submits (
    tx_hash TEXT PRIMARY KEY,
    tx_cost NUMERIC(38, 18) NOT NULL,
    reward NUMERIC(38, 18) NOT NULL,
    block BIGINT NOT NULL,
    finalized BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS idx_submits_block
ON submits (block);

CREATE INDEX IF NOT EXISTS idx_submits_finalized
ON submits (finalized);
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
	_, err := pdb.Exec(createSubmitsTable)
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

	return &Submitter{
		PoolId:              poolId.Uint64(),
		pdb:                 pdb,
		Address:             address,
		nonce:               nonce,
		privateKey:          privateKey,
		chainClient:         chainClient,
		chainId:             chainId,
		powAddress:          powAddress,
		poolRegistryAddress: poolRegistryAddress,
		pow:                 pow,
		powInstance:         *powInstance,
		poolRegistryDomain: apitypes.TypedDataDomain{
			Name:              poolRegistryDomain.Name,
			Version:           poolRegistryDomain.Version,
			ChainId:           (*math.HexOrDecimal256)(poolRegistryDomain.ChainId),
			VerifyingContract: poolRegistryDomain.VerifyingContract.Hex(),
		},
	}
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

	// signature
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
	_, err = s.pdb.Exec(insertSubmit, tx.Hash().Hex(), utils.WeiToEther(txCost).String(), utils.WeiToEther(reward).String(), receipt.BlockNumber.String())
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (s *Submitter) FinalizeRewards(shares map[string]*big.Float) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	totalShare := new(big.Float) // Defaults to 0
	for _, share := range shares {
		totalShare.Add(totalShare, share)
	}

	var costStr string
	var count int

	stats := s.pdb.QueryRow(unfinalizedStats)
	if err := stats.Scan(&costStr, &count); err != nil {
		return err
	}
	if count == 0 {
		slog.Info("Skip FinalizeRewards - nothing to submit")
		return nil
	}

	return nil
}

func (s *Submitter) GetClaimInfo(minerAddress common.Address) (*big.Int, []byte, error) {
	totalReward := big.NewInt(0)
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
