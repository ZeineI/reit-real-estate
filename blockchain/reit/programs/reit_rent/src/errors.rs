use anchor_lang::prelude::*;

#[error_code]
pub enum ReitError {
    #[msg("Math overflow")]
    MathOverflow,
    #[msg("Zero supply")]
    ZeroSupply,
}
