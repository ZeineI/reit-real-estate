import { Connection, Keypair, PublicKey, clusterApiUrl, LAMPORTS_PER_SOL } from "@solana/web3.js";
import { createMint, getOrCreateAssociatedTokenAccount, mintTo, setAuthority, AuthorityType } from "@solana/spl-token";
import * as fs from "fs";

// === Настройки ===
// DEVNET по умолчанию, можно заменить через переменную окружения SOLANA_CLUSTER_URL
const RPC = process.env.SOLANA_CLUSTER_URL || clusterApiUrl("devnet");
// Путь к keypair админа (Phantom экспорт или solana-keygen id.json в JSON)
const KEYPAIR_PATH = process.env.KEYPAIR_PATH || `${process.env.HOME}/.config/solana/id.json`;
// Сколько токенов выпустить (человеческих, без учёта decimals)
const TOTAL_SUPPLY = Number(process.env.TOTAL_SUPPLY || "1000");
// Decimals для REIT-токена (9 = nano)
const DECIMALS = 9;

function loadKeypair(path: string): Keypair {
    const raw = JSON.parse(fs.readFileSync(path, "utf-8"));
    const secretKey = Uint8Array.from(raw);
    return Keypair.fromSecretKey(secretKey);
}

(async () => {
    console.log("RPC:", RPC);
    console.log("KEYPAIR_PATH:", KEYPAIR_PATH);

    const admin = loadKeypair(KEYPAIR_PATH);
    const connection = new Connection(RPC, "confirmed");

    // Airdrop на devnet (если нужно)
    try {
        const bal = await connection.getBalance(admin.publicKey);
        if (bal < 0.5 * LAMPORTS_PER_SOL) {
            console.log("Airdrop 1 SOL (devnet)...");
            const sig = await connection.requestAirdrop(admin.publicKey, 1 * LAMPORTS_PER_SOL);
            await connection.confirmTransaction(sig, "confirmed");
        }
    } catch { /* ignore on mainnet/local */ }

    // 1) Создаём mint с decimals=9, mint authority = admin
    console.log("Creating REIT mint...");
    const reitMint = await createMint(
        connection,
        admin,                 // payer
        admin.publicKey,       // mintAuthority
        null,                  // freezeAuthority (null = нет freeze)
        DECIMALS               // decimals
    );
    console.log("REIT_MINT:", reitMint.toBase58());

    // 2) Создаём/получаем ATA для админа под этот mint
    const adminAta = await getOrCreateAssociatedTokenAccount(
        connection,
        admin,
        reitMint,
        admin.publicKey // owner
    );
    console.log("ADMIN_REIT_ATA:", adminAta.address.toBase58());

    // 3) Mint фиксированный выпуск на ATA админа
    const amountRaw = BigInt(TOTAL_SUPPLY) * BigInt(10 ** DECIMALS);
    console.log(`Minting ${TOTAL_SUPPLY} tokens (raw=${amountRaw.toString()})...`);
    await mintTo(
        connection,
        admin,
        reitMint,
        adminAta.address,
        admin, // authority = admin
        Number(amountRaw) // spl-token lib принимает number; для больших сумм — делайте чанками
    );

    // 4) Отключить дальнейший mint (фиксируем выпуск)
    console.log("Disabling further minting (set mintAuthority to null)...");
    await setAuthority(
        connection,
        admin,
        reitMint,
        admin.publicKey,
        AuthorityType.MintTokens,
        null // newAuthority = null
    );

    console.log("\n✅ DONE");
    console.log("ReitMint pubkey:", reitMint.toBase58());
    console.log("Admin ATA pubkey:", adminAta.address.toBase58());
    console.log("Decimals:", DECIMALS);
    console.log("TotalSupply (human):", TOTAL_SUPPLY);
})();
