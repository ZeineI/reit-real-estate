use crate::errors::ReitError;

/// new_index = old_index + amount_micro * 1e12 / total_supply_nano
pub fn accrue_index(old_index: u128, amount_micro: u64, total_supply_nano: u64) -> Result<u128, ReitError> {
    if total_supply_nano == 0 { return Err(ReitError::ZeroSupply); }
    let num = (amount_micro as u128)
        .checked_mul(1_000_000_000_000u128)
        .ok_or(ReitError::MathOverflow)?;
    let inc = num.checked_div(total_supply_nano as u128).ok_or(ReitError::MathOverflow)?;
    old_index.checked_add(inc).ok_or(ReitError::MathOverflow)
}

/// claim_micro = balance_nano * (global_index - last_index) / 1e12
pub fn calc_claim(balance_nano: u64, idx_now: u128, idx_last: u128) -> Result<u64, ReitError> {
    let delta = idx_now.checked_sub(idx_last).ok_or(ReitError::MathOverflow)?;
    let num = (balance_nano as u128)
        .checked_mul(delta).ok_or(ReitError::MathOverflow)?;
    Ok((num / 1_000_000_000_000u128) as u64)
}
