package util

import (
	"fmt"

	gsol "github.com/gagliardetto/solana-go"
)

func DeriveATA(owner, mint string) (string, error) {
	ownerPk, err := gsol.PublicKeyFromBase58(owner)
	if err != nil {
		return "", fmt.Errorf("DeriveATA: invalid owner: %w", err)
	}
	mintPk, err := gsol.PublicKeyFromBase58(mint)
	if err != nil {
		return "", fmt.Errorf("DeriveATA: invalid mint: %w", err)
	}

	ata, _, err := gsol.FindAssociatedTokenAddress(ownerPk, mintPk)
	if err != nil {
		return "", fmt.Errorf("DeriveATA: %w", err)
	}
	return ata.String(), nil
}

func DerivePDA(programID string, seeds ...[]byte) (pubkey string, bump uint8, err error) {
	prog, err := gsol.PublicKeyFromBase58(programID)
	if err != nil {
		return "", 0, fmt.Errorf("DerivePDA: invalid program id: %w", err)
	}
	pda, bump, err := gsol.FindProgramAddress(seeds, prog)
	if err != nil {
		return "", 0, fmt.Errorf("DerivePDA: %w", err)
	}
	return pda.String(), bump, nil
}

//ix := solana.NewInstruction(
//programID,
//solana.AccountMetaSlice{
//{PubKey: admin,        IsSigner: true,  IsWritable: true},  // админ
//{PubKey: statePDA,     IsSigner: false, IsWritable: true},  // state
//{PubKey: reitMint,     IsSigner: false, IsWritable: true},  // mint
//{PubKey: treasuryATA,  IsSigner: false, IsWritable: true},  // казна
//{PubKey: usdcMint,     IsSigner: false, IsWritable: false}, // USDC mint
//// ... + token_program, system_program, rent ...
//},
//dataBytes, // данные (discriminator + args)
//)
