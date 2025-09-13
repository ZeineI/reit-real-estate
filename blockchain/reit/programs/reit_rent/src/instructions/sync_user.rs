use anchor_lang::prelude::*;
use crate::state::*;

pub fn handler(ctx: Context<SyncUser>) -> Result<()> {
    let u = &mut ctx.accounts.user_state;
    let prop = &ctx.accounts.property;

    // синхронизируем доходность
    let current_idx = prop.accrual_index;
    u.last_index = current_idx;

    Ok(())
}

#[derive(Accounts)]
pub struct SyncUser<'info> {
    #[account(mut)]
    pub user: Signer<'info>,

    pub property: Account<'info, PropertyState>,

    #[account(
        init,
        payer = user,
        space = UserState::SIZE,
        seeds = [b"user_state", property.key().as_ref(), user.key().as_ref()],
        bump
    )]
    pub user_state: Account<'info, UserState>,

    pub system_program: Program<'info, System>,
    pub rent: Sysvar<'info, Rent>,
}
