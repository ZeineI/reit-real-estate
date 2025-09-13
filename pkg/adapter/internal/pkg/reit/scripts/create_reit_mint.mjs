import { Connection, Keypair, clusterApiUrl, LAMPORTS_PER_SOL } from "@solana/web3.js";
import { createMint, getOrCreateAssociatedTokenAccount, mintTo, setAuthority, AuthorityType } from "@solana/spl-token";
import fs from "fs";

// === настройки ===
const RPC = process.env.SOLANA_CLUSTER_URL || clusterApiUrl("devnet");
const DECIMALS = 9;                          // 9 знаков для REIT
const TOTAL_SUPPLY = Number(process.env.TOTAL_SUPPLY || "1000"); // человеческих токенов

// сохранение/загрузка ключа, чтобы последующие запуски были тем же аккаунтом
const KEYFILE = "./admin-keypair.json";
function loadOrCreateKeypair() {
  if (fs.existsSync(KEYFILE)) {
    const raw = JSON.parse(fs.readFileSync(KEYFILE, "utf-8"));
    return Keypair.fromSecretKey(Uint8Array.from(raw));
  }
  const kp = Keypair.generate();
  fs.writeFileSync(KEYFILE, JSON.stringify(Array.from(kp.secretKey)));
  console.log(`🗝  Новый ключ сохранён в ${KEYFILE}`);
  return kp;
}

(async () => {
  console.log("RPC:", RPC);
  const admin = loadOrCreateKeypair();
  console.log("ADMIN:", admin.publicKey.toBase58());

  const conn = new Connection(RPC, "confirmed");

  // airdrop (devnet)
  try {
    const bal = await conn.getBalance(admin.publicKey);
    if (bal < 0.5 * LAMPORTS_PER_SOL) {
      console.log("💧 Airdrop 1 SOL…");
      const sig = await conn.requestAirdrop(admin.publicKey, 1 * LAMPORTS_PER_SOL);
      await conn.confirmTransaction(sig, "confirmed");
    }
  } catch { /* если mainnet/local — пропустим */ }

  // 1) mint (decimals=9, authority=admin)
  console.log("🪙 Создаю mint…");
  const reitMint = await createMint(conn, admin, admin.publicKey, null, DECIMALS);
  console.log("REIT_MINT:", reitMint.toBase58());

  // 2) ATA для админа
  const adminAta = await getOrCreateAssociatedTokenAccount(conn, admin, reitMint, admin.publicKey);
  console.log("ADMIN_REIT_ATA:", adminAta.address.toBase58());

  // 3) mint фикс-сапплай
  const rawAmount = BigInt(TOTAL_SUPPLY) * BigInt(10 ** DECIMALS);
  console.log(`⛏  Mint ${TOTAL_SUPPLY} (${rawAmount.toString()} raw)…`);
  // В библиотеке mintTo amount — number; если не влезает, минтьте несколькими вызовами
    const amountNumber = Number(rawAmount); // 1_000_000_000_000 для 1000 * 1e9
    console.log(`⛏  Mint one-shot: ${amountNumber} raw...`);
    const sig = await mintTo(
        conn,
        admin,
        reitMint,
        adminAta.address,
        admin,
        amountNumber
    );

// Явно дожидаемся подтверждения:
    await conn.confirmTransaction(sig, "confirmed");
    console.log("   Mint confirmed:", sig);

  // 4) выключить возможность минтинга (фиксируем выпуск)
  console.log("🔒 Отключаю mint authority…");
  await setAuthority(conn, admin, reitMint, admin.publicKey, AuthorityType.MintTokens, null);

  console.log("\n✅ DONE");
  console.log("ReitMint pubkey:", reitMint.toBase58());
  console.log("Admin ATA pubkey:", adminAta.address.toBase58());
  console.log("Decimals:", DECIMALS);
  console.log("TotalSupply (human):", TOTAL_SUPPLY);
})();
