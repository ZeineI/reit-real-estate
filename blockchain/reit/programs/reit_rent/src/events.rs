use anchor_lang::prelude::*;

#[event]
pub struct PropertyCreated {
    pub property: Pubkey,
    pub reit_mint: Pubkey,
    pub usdc_mint: Pubkey,
    pub total_supply_nano: u64,
}

#[event]
pub struct RentDeposited {
    pub property: Pubkey,
    pub admin: Pubkey,
    pub amount_micro: u64,
    pub new_index: u128,
}
