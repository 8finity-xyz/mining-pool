package problemprovider

import (
	"crypto/ecdsa"
	"infinity/pool/internal"
	"infinity/pool/internal/contracts/PoW"
	"infinity/pool/internal/utils"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Problem struct {
	Nonce       uint64
	PrivateKeyA *ecdsa.PrivateKey
	Difficulty  *big.Int
}

type ProblemProvider struct {
	Problem *Problem
}

func NewProblemProvider(chainClient *ethclient.Client, ws string) *ProblemProvider {
	problemProvider := ProblemProvider{}

	problems := make(chan Problem)

	err := pollProblem(chainClient, problems)
	if err != nil {
		log.Fatal(err)
	}

	err = listenProblem(ws, problems)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for newProblem := range problems {
			if problemProvider.Problem == nil || problemProvider.Problem.Nonce < newProblem.Nonce {
				problemProvider.Problem = &newProblem
			}
		}
	}()

	return &problemProvider
}

func pollProblem(chainClient *ethclient.Client, problems chan<- Problem) error {
	pow := PoW.NewPoW()
	instance := pow.Instance(chainClient, common.HexToAddress(internal.PoWAddress))

	go func() {
		queryTicker := time.NewTicker(time.Second)
		defer queryTicker.Stop()

		for range queryTicker.C {
			currentProblem, err := bind.Call(instance, nil, pow.PackCurrentProblem(), pow.UnpackCurrentProblem)
			if err == nil {
				privateKeyA, _ := utils.ParsePrivateKey(currentProblem.Arg1)
				problems <- Problem{
					Nonce:       (*currentProblem.Arg0).Uint64(),
					PrivateKeyA: privateKeyA,
					Difficulty:  currentProblem.Arg2,
				}
			}
		}
	}()

	return nil
}

func listenProblem(ws string, problems chan<- Problem) error {
	conn, err := ethclient.Dial(ws)
	if err != nil {
		return err
	}

	pow := PoW.NewPoW()
	instance := pow.Instance(conn, common.HexToAddress(internal.PoWAddress))
	logs, sub, err := instance.WatchLogs(nil, "NewProblem")
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal("Subscription error:", err)
			case newPorblemLog := <-logs:
				newProblem, err := pow.UnpackNewProblemEvent(&newPorblemLog)
				if err != nil {
					log.Println("Parse error:", err)
					continue
				}
				privateKeyA, _ := utils.ParsePrivateKey(newProblem.PrivateKeyA)

				problems <- Problem{
					Nonce:       (*newProblem.Nonce).Uint64(),
					PrivateKeyA: privateKeyA,
					Difficulty:  newProblem.Difficulty,
				}
			}
		}
	}()

	return nil
}
