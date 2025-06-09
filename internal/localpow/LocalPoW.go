package localpow

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"infinity/pool/internal/problemprovider"
	"infinity/pool/internal/submitter"
	"infinity/pool/internal/utils"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/redis/go-redis/v9"
)

var (
	magic                 = new(big.Int).SetBytes(common.FromHex("0x8888888888888888888888888888888888888888"))
	submitsPrefix         = "submits:"
	difficultyKey         = "infinityPool.difficulty"
	sharesKey             = "infinityPool.shares"
	speedNumSubmits       = int64(100)
	speedTime             = int64(100)
	maxDifficultyIncrease = int64(8)
)

type LocalPow struct {
	problemProvider *problemprovider.ProblemProvider
	submitter       *submitter.Submitter
	privateKeyPool  *ecdsa.PrivateKey

	rdb *redis.Client

	globalLock sync.RWMutex
	minerLocks map[common.Address]*sync.Mutex
}

func NewLocalPow(
	rdb *redis.Client,
	problemProvider *problemprovider.ProblemProvider,
	submitter *submitter.Submitter,
) *LocalPow {
	privateKeyPool, err := utils.ParsePrivateKey(big.NewInt(int64(submitter.PoolId)))
	if err != nil {
		log.Fatal(errors.Join(errors.New("can't parse poolId"), err))
	}

	return &LocalPow{
		problemProvider: problemProvider,
		submitter:       submitter,
		privateKeyPool:  privateKeyPool,
		rdb:             rdb,
		minerLocks:      make(map[common.Address]*sync.Mutex),
	}
}

func (pow *LocalPow) Submit(minerAddress common.Address, privateKeyB *ecdsa.PrivateKey) error {
	unlock := pow.lockFor(minerAddress)
	defer unlock()

	difficulty, err := pow.getMinerDifficulty(minerAddress)
	if err != nil {
		return errors.Join(fmt.Errorf("can't get difficulty"), err)
	}

	err = pow.processSubmit(minerAddress, difficulty, privateKeyB)
	if err != nil {
		return err
	}
	pow.rewardMiner(minerAddress, difficulty)
	pow.adjustDifficulty(minerAddress, difficulty)
	return nil
}

func (pow *LocalPow) GetProblem(minerAddress common.Address) (*big.Int, *big.Int, error) {
	unlock := pow.lockFor(minerAddress)
	defer unlock()

	privateKeyMiner, err := utils.ParsePrivateKey(minerAddress.Big())
	if err != nil {
		return nil, nil, errors.Join(errors.New("can't parse minerAddress"), err)
	}

	privateKeyA_, err1 := utils.EcAdd(pow.problemProvider.Problem.PrivateKeyA, privateKeyMiner)
	privateKeyA_, err2 := utils.EcAdd(privateKeyA_, pow.privateKeyPool)
	if err1 != nil || err2 != nil {
		return nil, nil, errors.Join(errors.New("can't construct privateKeyA'"), err1, err2)
	}

	difficulty, err := pow.getMinerDifficulty(minerAddress)
	if err != nil {
		return nil, nil, err
	}
	return difficulty, privateKeyA_.D, nil
}

func (pow *LocalPow) lockFor(minerAddress common.Address) (unlock func()) {
	pow.globalLock.RLock()
	m, ok := pow.minerLocks[minerAddress]
	pow.globalLock.RUnlock()

	if !ok {
		pow.globalLock.Lock()
		defer pow.globalLock.Unlock()

		m, ok = pow.minerLocks[minerAddress]
		if !ok {
			m = &sync.Mutex{}
			pow.minerLocks[minerAddress] = m
		}
	}

	m.Lock()
	return func() { m.Unlock() }
}

func (pow *LocalPow) getMinerDifficulty(minerAddress common.Address) (*big.Int, error) {
	return loadBigInt(pow.rdb.HGet(context.Background(), difficultyKey, minerAddress.String()), common.MaxAddress.Big())
}

func (pow *LocalPow) processSubmit(minerAddress common.Address, difficulty *big.Int, privateKeyB *ecdsa.PrivateKey) error {
	privateKeyMiner, err := utils.ParsePrivateKey(minerAddress.Big())
	if err != nil {
		return errors.Join(errors.New("can't parse minerAddress"), err)
	}

	privateKeyB_, err := utils.EcAdd(privateKeyB, privateKeyMiner)
	if err != nil {
		return errors.Join(errors.New("can't construct privateKeyB'"), err)
	}

	privateKeyAB_, err := utils.EcAdd(pow.problemProvider.Problem.PrivateKeyA, privateKeyB_)
	if err != nil {
		return errors.Join(errors.New("can't construct privateKeyAB'"), err)
	}

	privateKeyAB, err := utils.EcAdd(privateKeyAB_, pow.privateKeyPool)
	if err != nil {
		return errors.Join(errors.New("can't construct privateKeyAB"), err)
	}

	addressAB := crypto.PubkeyToAddress(privateKeyAB.PublicKey)
	solution := new(big.Int).Xor(addressAB.Big(), magic)
	if solution.Cmp(difficulty) > 0 {
		return fmt.Errorf("not a solution - %s, 0x%s", addressAB.Hex(), difficulty.Text(16))
	}
	if solution.Cmp(pow.problemProvider.Problem.Difficulty) < 0 {
		go pow.submitter.Submit(privateKeyB_, privateKeyAB)
	}

	return nil
}

func (pow *LocalPow) rewardMiner(minerAddress common.Address, difficulty *big.Int) error {
	ctx := context.Background()

	share := new(big.Int).Div(common.MaxAddress.Big(), difficulty)
	shares, err := loadBigInt(pow.rdb.HGet(ctx, sharesKey, minerAddress.String()), big.NewInt(0))
	if err != nil {
		return err
	}
	return pow.rdb.HSet(ctx, sharesKey, minerAddress.String(), shares.Add(shares, share).String()).Err()
}

func (pow *LocalPow) adjustDifficulty(minerAddress common.Address, currentDifficulty *big.Int) error {
	ctx := context.Background()
	key := submitsPrefix + minerAddress.String()

	numSubmits, err := pow.rdb.LPush(ctx, key, time.Now().Unix()).Result()
	go pow.rdb.Expire(ctx, key, time.Duration(speedTime)*time.Second)
	if err != nil {
		return errors.Join(fmt.Errorf("can't save submit"), err)
	}
	if numSubmits < int64(speedNumSubmits) {
		return nil
	}

	firstSubmit, _ := pow.rdb.RPop(ctx, key).Int64()
	lastSubmit, _ := pow.rdb.LPop(ctx, key).Int64()
	pow.rdb.Del(ctx, key)

	duration := lastSubmit - firstSubmit
	adjustedDifficulty := new(big.Int).Set(currentDifficulty)
	adjustedDifficulty.Mul(currentDifficulty, big.NewInt(duration))
	adjustedDifficulty.Div(adjustedDifficulty, big.NewInt(speedTime))

	if adjustedDifficulty.Cmp(common.MaxAddress.Big()) > 0 {
		adjustedDifficulty = common.MaxAddress.Big()
	}

	adjustedDifficulty = max(
		adjustedDifficulty,
		new(big.Int).Div(currentDifficulty, big.NewInt(maxDifficultyIncrease)),
	)

	adjustedDifficulty = max(
		adjustedDifficulty,
		pow.problemProvider.Problem.Difficulty,
	)

	err = pow.rdb.HSetEXWithArgs(
		ctx,
		difficultyKey,
		&redis.HSetEXOptions{
			ExpirationType: redis.HSetEXExpirationEX,
			ExpirationVal:  speedTime,
		},
		minerAddress.String(),
		adjustedDifficulty.String(),
	).Err()
	if err != nil {
		return err
	}
	return nil
}

func (pow *LocalPow) DistributeRewards() error {
	ctx := context.Background()

	pow.globalLock.Lock()
	sharesRaw, err := pow.rdb.HGetAll(ctx, sharesKey).Result()
	pow.rdb.Del(ctx, sharesKey)
	pow.globalLock.Unlock()

	if err != nil || len(sharesRaw) == 0 {
		return err
	}

	shares := make(map[string]*big.Int, len(sharesRaw))
	for minerAddress, sharesRaw := range sharesRaw {
		share, ok := new(big.Int).SetString(sharesRaw, 10)
		if ok {
			shares[minerAddress] = share
		}
	}

	finalizeRewardsErr := pow.submitter.FinalizeRewards(shares)
	if finalizeRewardsErr == nil {
		return nil
	}

	pow.globalLock.Lock()
	defer pow.globalLock.Unlock()
	newSharesRaw, err := pow.rdb.HGetAll(ctx, sharesKey).Result()
	if err != nil {
		return err
	}

	for minerAddress, share := range shares {
		newShareRaw, ok := newSharesRaw[minerAddress]

		var newShare *big.Int
		if ok {
			newShare, _ = new(big.Int).SetString(newShareRaw, 10)
			newShare.Add(newShare, share)
		} else {
			newShare = share
		}
		pow.rdb.HSet(ctx, sharesKey, minerAddress, newShare.String())
	}

	return finalizeRewardsErr
}

func (pow *LocalPow) Hashrate() float64 {
	ctx := context.Background()
	difficultiesRaw, err := pow.rdb.HGetAll(ctx, difficultyKey).Result()
	if err != nil {
		return 0
	}

	hashrate := big.NewFloat(0)
	maxDifficulty := new(big.Float).SetInt(common.MaxAddress.Big())
	for _, difficultyRaw := range difficultiesRaw {
		difficulty, _ := new(big.Float).SetString(difficultyRaw)
		hashrate.Add(hashrate, new(big.Float).Quo(maxDifficulty, difficulty))
	}

	h, _ := hashrate.Float64()
	return h
}

func max(a *big.Int, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

func loadBigInt(cmd *redis.StringCmd, defaultValue *big.Int) (*big.Int, error) {
	valueRaw, err := cmd.Result()

	switch {
	case err == redis.Nil:
		return defaultValue, nil
	case err != nil:
		return nil, err
	default:
		value, ok := new(big.Int).SetString(valueRaw, 10)
		if ok {
			return value, nil
		} else {
			return nil, fmt.Errorf("can't parse %s", valueRaw)
		}
	}

}
