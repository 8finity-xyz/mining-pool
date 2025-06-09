package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
)

func QuoteS28(amountIn *big.Int) (amountOut *big.Int, err error) {
	jsonData, _ := json.Marshal(map[string]any{
		"chainId": 146,
		"inputTokens": []struct {
			TokenAddress string `json:"tokenAddress"`
			Amount       string `json:"amount"`
		}{{TokenAddress: "0x0000000000000000000000000000000000000000", Amount: amountIn.String()}},
		"outputTokens": []struct {
			TokenAddress string `json:"tokenAddress"`
			Proportion   int    `json:"proportion"`
		}{{TokenAddress: "0x888852d1c63c7b333efEb1c4C5C79E36ce918888", Proportion: 1}},
	})
	resp, err := http.Post("https://api.odos.xyz/sor/quote/v2", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response error: %s", string(body))
	}

	var quote struct {
		OutAmounts []string `json:"outAmounts"`
	}
	if err := json.Unmarshal(body, &quote); err != nil {
		return nil, err
	}

	if len(quote.OutAmounts) != 1 {
		return nil, errors.New("quote output broken - %s")
	}

	amountOut, ok := new(big.Int).SetString(quote.OutAmounts[0], 10)
	if !ok {
		return nil, fmt.Errorf("outAmounts broken - %s", quote.OutAmounts[0])
	}
	return amountOut, nil
}
