use crate::{execute, Hydentity};
use client_sdk::{
    helpers::{risc0::Risc0Prover, ClientSdkExecutor},
    transaction_builder::{ProvableBlobTx, StateUpdater, TxExecutorBuilder},
};
use sdk::{identity_provider::IdentityAction, utils::as_hyle_output, ContractName, HyleOutput};
use std::any::Any;

pub mod metadata {
    pub const HYDENTITY_ELF: &[u8] = include_bytes!("../hydentity.img");
    pub const PROGRAM_ID: [u8; 32] = sdk::str_to_u8(include_str!("../hydentity.txt"));
}
use metadata::*;

struct HydentityPseudoExecutor {}
impl ClientSdkExecutor for HydentityPseudoExecutor {
    fn execute(
        &self,
        contract_input: &sdk::ContractInput,
    ) -> anyhow::Result<(Box<dyn Any>, HyleOutput)> {
        let initial_state: Hydentity = borsh::from_slice(contract_input.state.as_slice())?;
        let mut res = execute(initial_state.clone(), contract_input.clone());

        let output = as_hyle_output(initial_state, contract_input.clone(), &mut res);
        match res {
            Ok(res) => Ok((Box::new(res.1.clone()), output)),
            Err(e) => Err(anyhow::anyhow!(e)),
        }
    }
}

impl Hydentity {
    pub fn setup_builder<S: StateUpdater>(
        &self,
        contract_name: ContractName,
        builder: &mut TxExecutorBuilder<S>,
    ) {
        builder.init_with(
            contract_name,
            HydentityPseudoExecutor {},
            Risc0Prover::new(HYDENTITY_ELF),
        );
    }
}

pub fn verify_identity(
    builder: &mut ProvableBlobTx,
    contract_name: ContractName,
    state: &Hydentity,
    password: String,
) -> anyhow::Result<()> {
    let nonce = state
        .get_nonce(builder.identity.0.as_str())
        .map_err(|e| anyhow::anyhow!(e))?;

    let password = password.into_bytes().to_vec();

    builder.add_action(
        contract_name,
        IdentityAction::VerifyIdentity {
            account: builder.identity.0.clone(),
            nonce,
        },
        Some(password),
        None,
        None,
    )?;
    Ok(())
}

pub fn register_identity(
    builder: &mut ProvableBlobTx,
    contract_name: ContractName,
    password: String,
) -> anyhow::Result<()> {
    let password = password.into_bytes().to_vec();

    builder.add_action(
        contract_name,
        IdentityAction::RegisterIdentity {
            account: builder.identity.0.clone(),
        },
        Some(password),
        None,
        None,
    )?;
    Ok(())
}
