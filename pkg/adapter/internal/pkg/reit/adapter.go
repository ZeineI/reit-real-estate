package reit

import (
	"context"
	"errors"

	"reit-real-estate/pkg/adapter/internal/anchor/tx"
	"reit-real-estate/pkg/adapter/internal/anchor/utils"
	"reit-real-estate/pkg/adapter/internal/solana"
)

type Adapter interface {
	// network
	Ping(ctx context.Context) error

	// tx builders
	CreatePropertyTx(ctx context.Context, p CreatePropertyParams) (TxBase64, error)
}

type Options struct {
	RPCURL    string
	ProgramID string
}

func New(opts Options) (Adapter, error) {
	if opts.RPCURL == "" || opts.ProgramID == "" {
		return nil, errors.New("reit: RPCURL and ProgramID are required")
	}
	return &adapter{
		rpc:     solana.NewClient(opts.RPCURL),
		program: opts.ProgramID,
	}, nil
}

type adapter struct {
	rpc     *solana.Client
	program string
}

func (a *adapter) Ping(ctx context.Context) error {
	_, err := a.rpc.RecentBlockhash(ctx)
	return err
}

func (a *adapter) CreatePropertyTx(ctx context.Context, p CreatePropertyParams) (TxBase64, error) {
	if p.WalletAddress == "" {
		return TxBase64{}, errors.New("reit: empty WalletAddress")
	}
	if p.TokenAddress == "" {
		return TxBase64{}, errors.New("reit: empty TokenAddress (payout mint)")
	}
	if p.TotalSupplyTokens == "" {
		return TxBase64{}, errors.New("reit: empty TotalSupplyTokens")
	}

	totalSupplyNano, err := util.TokensToNano(p.TotalSupplyTokens)
	if err != nil {
		return TxBase64{}, err
	}

	blockhash, err := a.rpc.RecentBlockhash(ctx)
	if err != nil {
		return TxBase64{}, err
	}

	msg, err := tx.BuildCreatePropertyMessage(tx.CreatePropertyInput{
		FeePayer:        p.WalletAddress,
		ScAddress:       a.program,
		AdminPubkey:     p.WalletAddress,
		USDCMint:        p.TokenAddress,
		ReitMint:        p.ReitMint,
		TotalSupplyNano: uint64(totalSupplyNano),
		Blockhash:       blockhash,
	})
	if err != nil {
		return TxBase64{}, err
	}

	b64, err := tx.EncodeMessageBase64(msg)
	if err != nil {
		return TxBase64{}, err
	}
	return TxBase64{Tx: b64}, nil
}
