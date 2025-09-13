use anchor_lang::prelude::*;

#[account]
pub struct PropertyState {
    pub reit_mint: Pubkey,       // адрес реит-минта (SPL)
    pub usdc_mint: Pubkey,       // адрес payout mint
    pub admin: Pubkey,           // кто создал
    pub total_supply_nano: u64,  // фиксированный выпуск (9 dec)
    pub accrual_index: u128,     // индекс дохода (в микрух/на токен, напр.)
    pub bump: u8,                // bump PDA
}

impl PropertyState {
    // 8 (disc) + 32*3 + 8 + 16 + 1 = 8 + 96 + 8 + 16 + 1 = 129 => округлим до 136
    pub const SIZE: usize = 8 + 32 + 32 + 32 + 8 + 16 + 1;
}

#[account]
pub struct UserState {
    pub user: Pubkey,
    pub property: Pubkey,
    pub last_index: u128, // индекс на момент последней синхронизации
    pub bump: u8,
}

impl UserState {
    pub const SIZE: usize = 8 + 32 + 32 + 16 + 1;
}
