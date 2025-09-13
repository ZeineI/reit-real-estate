use anchor_lang::prelude::*;
use anchor_spl::token::{self, Token, TokenAccount, Transfer, Mint};

use crate::{state::*, math::accrue_index, events::RentDeposited};

pub fn handler(ctx: Context<DepositRent>, amount_micro: u64) -> Result<()> {
    // Перевод USDC от админа в казну
    let cpi_accounts = Transfer {
        from: ctx.accounts.admin_usdc_ata.to_account_info(),
        to: ctx.accounts.treasury.to_account_info(),
        authority: ctx.accounts.admin.to_account_info(),
    };
    let cpi_ctx = CpiContext::new(ctx.accounts.token_program.to_account_info(), cpi_accounts);
    token::transfer(cpi_ctx, amount_micro)?;

    // индекс дохода
    let prop = &mut ctx.accounts.property;
    prop.accrual_index = accrue_index(prop.accrual_index, amount_micro, prop.total_supply_nano)?;

    emit!(RentDeposited {
        property: prop.key(),
        admin: ctx.accounts.admin.key(),
        amount_micro,
        new_index: prop.accrual_index,
    });

    Ok(())
}

#[derive(Accounts)]
pub struct DepositRent<'info> {
    #[account(mut, address = property.admin)]
    pub admin: Signer<'info>,

    pub usdc_mint: Account<'info, Mint>,

    #[account(
        mut,
        seeds = [b"property", property.reit_mint.as_ref()],
        bump = property.bump,
    )]
    pub property: Account<'info, PropertyState>,

    /// USDC ATA админа, mint = usdc_mint
    #[account(mut)]
    pub admin_usdc_ata: Account<'info, TokenAccount>,

    /// Казна, ATA(owner = property, mint = usdc_mint)
    #[account(mut)]
    pub treasury: Account<'info, TokenAccount>,

    pub token_program: Program<'info, Token>,
}
