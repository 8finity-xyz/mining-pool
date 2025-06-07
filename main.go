package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"infinity/pool/internal/localpow"
	"infinity/pool/internal/problemprovider"
	"infinity/pool/internal/submitter"
	"log"
	"log/slog"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/valyala/fasthttp"
)

func Getenv(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	} else {
		return defaultValue
	}
}

var (
	Rpc            = Getenv("INFINITY_POOL_RPC", "https://rpc.soniclabs.com")
	Ws             = Getenv("INFINITY_POOL_WS", "wss://rpc.soniclabs.com")
	PoolPrivateKey = common.FromHex(Getenv("INFINITY_POOL_PRIVATE_KEY", ""))
	RedisUri       = Getenv("INFINITY_POOL_REDIS_URI", "")
	PostgresUri    = Getenv("INFINITY_POOL_POSTGRES_URI", "")
	PoolPort       = Getenv("INFINITY_POOL_PORT", ":18888")
	LogLevel       = Getenv("LOGLEVEL", "")
)

func main() {
	l := slog.LevelInfo
	l.UnmarshalText([]byte(LogLevel))
	slog.SetLogLoggerLevel(l)

	redisOptions, err := redis.ParseURL(RedisUri)
	if err != nil {
		log.Fatalf("Can't parse redis uri, %v", err)
	}
	rdb := redis.NewClient(redisOptions)

	pdb, err := sql.Open("postgres", PostgresUri)
	if err != nil {
		log.Fatalf("Can't create postgress client, %v", err)
	}

	chainClient, err := ethclient.Dial(Rpc)
	if err != nil {
		log.Fatalf("Can't create chain client, %v", err)
	}

	problemProvider := problemprovider.NewProblemProvider(chainClient, Ws)
	submitter := submitter.NewSubmitter(chainClient, pdb, PoolPrivateKey)
	pow := localpow.NewLocalPow(rdb, problemProvider, submitter)

	go func() {
		queryTicker := time.NewTicker(time.Minute)
		defer queryTicker.Stop()

		for range queryTicker.C {
			err := pow.FinalizeRewards()
			if err != nil {
				slog.Error("FinalizeRewards failed", "err", err)
			}
		}
	}()

	submitterBalance, _ := submitter.GetBalance()
	log.Printf("∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞")
	log.Printf("Pool id: #%d", submitter.PoolId)
	log.Printf("Pool operator address: %s", submitter.Address)
	log.Printf("Pool operator balance: %f $S", new(big.Float).Quo(new(big.Float).SetInt(submitterBalance), big.NewFloat(params.Ether)))
	log.Printf("Ensure, that it have enough funds")
	log.Printf("∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞")

	h := fasthttp.CompressHandler(requestHandler(pow, submitter))
	if err := fasthttp.ListenAndServe("0.0.0.0"+PoolPort, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}

func requestHandler(pow *localpow.LocalPow, submitter *submitter.Submitter) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/healthcheck":
			ctx.Response.StatusCode()
			ctx.SuccessString("text/html; charset=utf-8", "ok")
		case "/submit":
			submitHandler(ctx, pow)
		case "/problem":
			problemHandler(ctx, pow)
		case "/claim":
			claimHandler(ctx, submitter)
		case "/hashrate":
			hashrateHandler(ctx, pow)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}

		slog.Debug(string(ctx.Path()), "status", ctx.Response.StatusCode())
	}
}

func submitHandler(
	ctx *fasthttp.RequestCtx,
	pow *localpow.LocalPow,
) {
	miner := string(ctx.QueryArgs().Peek("miner"))
	if !common.IsHexAddress(miner) {
		ctx.Error("miner param - not an address", fasthttp.StatusBadRequest)
		return
	}
	minerAddress := common.HexToAddress(miner)

	privateKeyB, err := crypto.ToECDSA(common.FromHex(string(ctx.QueryArgs().Peek("private_key_b"))))
	if err != nil {
		ctx.Error(fmt.Sprintf("private_key_b param - %v", err), fasthttp.StatusBadRequest)
		return
	}

	err = pow.Submit(minerAddress, privateKeyB)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	} else {
		problemHandler(ctx, pow)
	}
}

func problemHandler(ctx *fasthttp.RequestCtx, pow *localpow.LocalPow) {
	miner := string(ctx.QueryArgs().Peek("miner"))
	if !common.IsHexAddress(miner) {
		ctx.Error("miner param - not an address", fasthttp.StatusBadRequest)
		return
	}
	minerAddress := common.HexToAddress(miner)

	difficulty, privateKeyA, err := pow.GetProblem(minerAddress)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	} else {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(struct {
			Difficulty  string `json:"difficulty"`
			PrivateKeyA string `json:"private_key_a"`
		}{difficulty.String(), privateKeyA.String()})
	}
}

func claimHandler(ctx *fasthttp.RequestCtx, submitter *submitter.Submitter) {
	miner := string(ctx.QueryArgs().Peek("miner"))
	if !common.IsHexAddress(miner) {
		ctx.Error("miner param - not an address", fasthttp.StatusBadRequest)
		return
	}
	minerAddress := common.HexToAddress(miner)

	totalReward, signature, err := submitter.GetClaimInfo(minerAddress)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	} else {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(struct {
			PoolId      uint64 `json:"pool_id"`
			TotalReward string `json:"total_reward"`
			Signature   string `json:"signature"`
		}{submitter.PoolId, totalReward.String(), "0x" + common.Bytes2Hex(signature)})
	}
}

func hashrateHandler(ctx *fasthttp.RequestCtx, pow *localpow.LocalPow) {
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.SetBytesV("Access-Control-Allow-Origin", ctx.Request.Header.Peek("Origin"))
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(struct {
		Hashrate float64 `json:"hashrate"`
	}{Hashrate: pow.Hashrate()})
}
