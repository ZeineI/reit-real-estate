package solana

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	rpcURL     string
	httpClient *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		rpcURL: url,
		httpClient: &http.Client{
			Timeout: 6 * time.Second,
		},
	}
}

func (c *Client) doRPC(ctx context.Context, method string, params []interface{}, out any) error {
	if c.rpcURL == "" {
		return errors.New("solana: empty rpc url")
	}
	reqBody, _ := json.Marshal(rpcRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.rpcURL, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		// ok
	default:
		return fmt.Errorf("solana: rpc http %d", resp.StatusCode)
	}

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(out)
}

func (c *Client) RecentBlockhash(ctx context.Context) (string, error) {
	var r rpcResponse[getLatestBlockhashResult]
	// запрос без params равнозначен ["finalized"]
	if err := c.doRPC(ctx, "getLatestBlockhash", nil, &r); err != nil {
		return "", err
	}
	if r.Error != nil {
		return "", fmt.Errorf("solana: rpc error %d %s", r.Error.Code, r.Error.Message)
	}
	if r.Result.Value.Blockhash == "" {
		return "", errors.New("solana: empty blockhash")
	}
	return r.Result.Value.Blockhash, nil
}
