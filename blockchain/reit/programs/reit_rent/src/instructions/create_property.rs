use anchor_lang::prelude::*;
use anchor_spl::{
    token::{Mint, Token, TokenAccount},
    associated_token::AssociatedToken,
};

use crate::{state::*, events::PropertyCreated};

pub fn handler(ctx: Context<CreateProperty>, total_supply_nano: u64) -> Result<()> {
    let prop = &mut ctx.accounts.property;

    prop.reit_mint = ctx.accounts.reit_mint.key();
    prop.usdc_mint = ctx.accounts.usdc_mint.key();
    prop.admin = ctx.accounts.admin.key();
    prop.total_supply_nano = total_supply_nano;
    prop.accrual_index = 0;
    prop.bump = ctx.bumps.property;

    emit!(PropertyCreated {
        property: prop.key(),
        reit_mint: prop.reit_mint,
        usdc_mint: prop.usdc_mint,
        total_supply_nano,
    });
    Ok(())
}

#[derive(Accounts)]
#[instruction(total_supply_nano: u64)]
pub struct CreateProperty<'info> {
    /// Админ (payer)
    #[account(mut)]
    pub admin: Signer<'info>,

    /// Уже существующий mint токена долей (создан офчейн)
    pub reit_mint: Account<'info, Mint>,

    /// Payout mint (USDC)
    pub usdc_mint: Account<'info, Mint>,

    /// PDA стейт «паспорт объекта»
    #[account(
        init,
        payer = admin,
        space = PropertyState::SIZE,
        seeds = [b"property", reit_mint.key().as_ref()],
        bump
    )]
    pub property: Account<'info, PropertyState>,

    /// Казна USDC: ATA(owner = property PDA, mint = usdc_mint)
    #[account(
        init,
        payer = admin,
        associated_token::mint = usdc_mint,
        associated_token::authority = property
    )]
    pub treasury: Account<'info, TokenAccount>,

    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    pub associated_token_program: Program<'info, AssociatedToken>,
    pub rent: Sysvar<'info, Rent>,
}
