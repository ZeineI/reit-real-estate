use anchor_lang::prelude::*;

pub mod state;
pub mod errors;
pub mod events;
pub mod math;
pub mod instructions;

use instructions::*;

declare_id!("8GKoDJSNEDbYWy1iF6Lewr8L3G9uuYSqqutB18QFC4Rw"); //TODO

#[program]
pub mod reit_rent {
    use super::*;

 pub fn create_property(
     ctx: Context<create_property::CreateProperty>,
     total_supply_nano: u64,
 ) -> Result<()> {
     create_property::handler(ctx, total_supply_nano)
 }


   pub fn deposit_rent(
       ctx: Context<deposit_rent::DepositRent>,
       amount_micro: u64,
   ) -> Result<()> {
       deposit_rent::handler(ctx, amount_micro)
   }

 pub fn sync_user(ctx: Context<sync_user::SyncUser>) -> Result<()> {
     sync_user::handler(ctx)
 }


   pub fn claim(ctx: Context<claim::Claim>) -> Result<()> {
       claim::handler(ctx)
   }

}