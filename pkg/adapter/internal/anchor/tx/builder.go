package tx

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	util "reit-real-estate/pkg/adapter/internal/anchor/utils"

	gsol "github.com/gagliardetto/solana-go"
)

type CreatePropertyInput struct {
	FeePayer        string // pubkey, кто платит комиссию и подписывает (админ)
	ScAddress       string // адрес твоей Anchor-программы (REIT_PROGRAM_ID)
	AdminPubkey     string // обычно = FeePayer
	USDCMint        string // payout mint (USDC)
	ReitMint        string // mint токена долей (если создаёте как PDA в контракте — можно вычислять)
	TotalSupplyNano uint64 // фиксированный выпуск в nano (9 dec)
	Blockhash       string // recent blockhash
}

var (
	tokenProgramID      = gsol.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	assocTokenProgramID = gsol.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25G6V5D8Nf")
	systemProgramID     = gsol.SystemProgramID
	rentSysVarID        = gsol.SysVarRentPubkey
)

// BuildCreatePropertyMessage — ОДНА ix: create_property(total_supply_nano)
func BuildCreatePropertyMessage(in CreatePropertyInput) (gsol.Message, error) {
	// 0) валидация
	if in.FeePayer == "" || in.ScAddress == "" || in.AdminPubkey == "" || in.USDCMint == "" || in.ReitMint == "" || in.Blockhash == "" {
		return gsol.Message{}, fmt.Errorf("tx: invalid CreatePropertyInput")
	}

	// 1) ключи
	progPK, err := gsol.PublicKeyFromBase58(in.ScAddress)
	if err != nil {
		return gsol.Message{}, fmt.Errorf("program id: %w", err)
	}
	feePayerPK, err := gsol.PublicKeyFromBase58(in.FeePayer)
	if err != nil {
		return gsol.Message{}, fmt.Errorf("fee payer: %w", err)
	}
	adminPK, err := gsol.PublicKeyFromBase58(in.AdminPubkey)
	if err != nil {
		return gsol.Message{}, fmt.Errorf("admin: %w", err)
	}
	usdcMintPK, err := gsol.PublicKeyFromBase58(in.USDCMint)
	if err != nil {
		return gsol.Message{}, fmt.Errorf("usdc mint: %w", err)
	}
	reitMintPK, err := gsol.PublicKeyFromBase58(in.ReitMint)
	if err != nil {
		return gsol.Message{}, fmt.Errorf("reit mint: %w", err)
	}
	recent, err := gsol.HashFromBase58(in.Blockhash)
	if err != nil {
		return gsol.Message{}, fmt.Errorf("blockhash: %w", err)
	}

	// 2) PDA/ATA только самое необходимое
	// state PDA (если твой #[derive(Accounts)] требует его явно):
	statePDA, _, err := util.DerivePDA(in.ScAddress, []byte("property_state"), reitMintPK.Bytes())
	if err != nil {
		return gsol.Message{}, err
	}
	statePDAPK := gsol.MustPublicKeyFromBase58(statePDA)

	// admin REIT ATA — куда упадёт стартовый supply
	adminReitATA, err := util.DeriveATA(in.AdminPubkey, in.ReitMint)
	if err != nil {
		return gsol.Message{}, err
	}
	adminReitATAPK := gsol.MustPublicKeyFromBase58(adminReitATA)

	// treasury USDC ATA — может создать программа сама; но адрес пригодится:
	// Если твой #[derive(Accounts)] просит передать treasury, раскомментируй:
	// treasuryATA, err := util.DeriveATA(in.AdminPubkey /*или PDA, если у тебя owner=программа*/, in.USDCMint)
	// if err != nil { return gsol.Message{}, err }
	// treasuryATAPK := gsol.MustPublicKeyFromBase58(treasuryATA)

	// 3) data: discriminator(8) + total_supply_nano(u64 LE)
	disc := sha256.Sum256([]byte("global:create_property"))
	data := make([]byte, 8+8)
	copy(data[0:8], disc[0:8])
	binary.LittleEndian.PutUint64(data[8:16], uint64(in.TotalSupplyNano))

	// 4) аккаунты (порядок = порядок в твоём #[derive(Accounts)] в Rust!)
	metas := gsol.AccountMetaSlice{
		{PublicKey: adminPK, IsSigner: true, IsWritable: true}, // admin (payer)
		// если в хендлере есть "reit_mint: init", мы передаём его адрес:
		{PublicKey: reitMintPK, IsSigner: false, IsWritable: true},     // reit_mint (init)
		{PublicKey: usdcMintPK, IsSigner: false, IsWritable: false},    // usdc_mint
		{PublicKey: adminReitATAPK, IsSigner: false, IsWritable: true}, // admin_reit_ata (init)
		{PublicKey: statePDAPK, IsSigner: false, IsWritable: true},     // property_state (init PDA)
		{PublicKey: tokenProgramID, IsSigner: false, IsWritable: false},
		{PublicKey: assocTokenProgramID, IsSigner: false, IsWritable: false},
		{PublicKey: systemProgramID, IsSigner: false, IsWritable: false},
		{PublicKey: rentSysVarID, IsSigner: false, IsWritable: false},
	}
	// Если твой Rust требует treasury_ata в аккаунтах — вставь его:
	// metas = append(metas, gsol.AccountMeta{PublicKey: treasuryATAPK, IsSigner: false, IsWritable: true})

	ix := gsol.NewInstruction(progPK, metas, data)

	tx, err := gsol.NewTransaction(
		[]gsol.Instruction{ix},
		recent,
		gsol.TransactionPayer(feePayerPK),
	)
	if err != nil {
		return gsol.Message{}, fmt.Errorf("new transaction: %w", err)
	}

	// Если тебе нужен именно base64 прямо тут — можно сразу вернуть строку.
	// Но если сигнатура функции должна вернуть Message — сделаем так:
	bin, err := tx.MarshalBinary()
	if err != nil {
		return gsol.Message{}, fmt.Errorf("marshal tx: %w", err)
	}
	_ = bin // ← можно вернуть как []byte или base64 выше по стеку

	// Если всё-таки нужен gsol.Message (для совместимости со старым кодом),
	// можно достать его из транзакции:
	msg := tx.Message
	return msg, nil
}
