use anchor_lang::prelude::*;
use anchor_spl::token::{self, Token, TokenAccount, Transfer};

use crate::state::*;

pub fn handler(ctx: Context<Claim>) -> Result<()> {
    let prop = &ctx.accounts.property;

    // вычисление доли.
    let amount_micro: u64 = 0; // TODO: рассчёт

    // PDA signer seeds
    let seeds: &[&[u8]] = &[
        b"property",
        prop.reit_mint.as_ref(),
        &[prop.bump],
    ];
    let seeds_binding = &[seeds];

    let cpi_accounts = Transfer {
        from: ctx.accounts.treasury.to_account_info(),
        to: ctx.accounts.investor_ata.to_account_info(),
        authority: ctx.accounts.property.to_account_info(),
    };
    let cpi_ctx = CpiContext::new_with_signer(
        ctx.accounts.token_program.to_account_info(),
        cpi_accounts,
        seeds_binding,
    );
    token::transfer(cpi_ctx, amount_micro)?;
    Ok(())
}

#[derive(Accounts)]
pub struct Claim<'info> {
    #[account(mut)]
    pub investor: Signer<'info>,

    #[account(mut)]
    pub property: Account<'info, PropertyState>,

    #[account(mut)]
    pub treasury: Account<'info, TokenAccount>,

    #[account(mut)]
    pub investor_ata: Account<'info, TokenAccount>,

    pub token_program: Program<'info, Token>,
}
