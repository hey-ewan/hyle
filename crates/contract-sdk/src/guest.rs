use alloc::string::{String, ToString};
use alloc::vec;
use borsh::{BorshDeserialize, BorshSerialize};

use crate::{
    flatten_blobs,
    utils::{as_hyle_output, check_caller_callees, parse_blob, parse_structured_blob},
    ContractInput, Digestable, HyleOutput, Identity, StructuredBlob,
};

pub trait GuestEnv {
    fn log(&self, message: &str);
    fn commit(&self, output: &HyleOutput);
    fn read<State: BorshDeserialize + 'static>(&self) -> (State, ContractInput);
}

pub struct Risc0Env;

#[cfg(feature = "risc0")]
impl GuestEnv for Risc0Env {
    fn log(&self, message: &str) {
        risc0_zkvm::guest::env::log(message);
    }

    fn commit(&self, output: &HyleOutput) {
        risc0_zkvm::guest::env::commit(output);
    }

    fn read<State: BorshDeserialize>(&self) -> (State, ContractInput) {
        let len: usize = risc0_zkvm::guest::env::read();
        let mut slice = vec![0u8; len];
        risc0_zkvm::guest::env::read_slice(&mut slice);
        let contract_input: ContractInput = borsh::from_slice(&slice).unwrap();
        let state: State = borsh::from_slice(&contract_input.state).unwrap();
        (state, contract_input)
    }
}

pub struct SP1Env;

// For coverage tests, assume risc0 if both are active
#[cfg(all(feature = "sp1", not(feature = "risc0")))]
impl GuestEnv for SP1Env {
    fn log(&self, message: &str) {
        // TODO: this does nothing actually
        sp1_zkvm::io::hint(&message);
    }

    fn commit(&self, output: &HyleOutput) {
        let vec = borsh::to_vec(&output).unwrap();
        sp1_zkvm::io::commit_slice(&vec);
    }

    fn read<State: BorshDeserialize>(&self) -> (State, ContractInput) {
        let vec = sp1_zkvm::io::read_vec();
        let contract_input: ContractInput = borsh::from_slice(&slice).unwrap();
        let state: State = borsh::from_slice(&contract_input.state).unwrap();
        (state, contract_input)
    }
}

pub fn fail<State: Digestable>(
    input: ContractInput,
    initial_state: State,
    message: &str,
) -> HyleOutput {
    let digest = initial_state.as_digest();
    HyleOutput {
        version: 1,
        initial_state: digest.clone(),
        next_state: digest,
        identity: input.identity,
        index: input.index,
        blobs: flatten_blobs(&input.blobs),
        success: false,
        tx_hash: input.tx_hash,
        tx_ctx: input.tx_ctx,
        registered_contracts: vec![],
        program_outputs: message.to_string().into_bytes(),
    }
}

pub fn panic(env: impl GuestEnv, message: &str) {
    env.log(message);
    // should we env::commit ?
    panic!("{}", message);
}

pub fn init_raw<Parameters>(input: ContractInput) -> (ContractInput, Option<Parameters>)
where
    Parameters: BorshDeserialize,
{
    let parsed_blob = parse_blob::<Parameters>(&input.blobs, &input.index);

    (input, parsed_blob)
}

pub fn init_with_caller<Parameters>(
    input: ContractInput,
) -> Result<(ContractInput, StructuredBlob<Parameters>, Identity), String>
where
    Parameters: BorshSerialize + BorshDeserialize,
{
    let parsed_blob = parse_structured_blob::<Parameters>(&input.blobs, &input.index);

    let parsed_blob = parsed_blob.ok_or("Failed to parse input blob".to_string())?;

    let caller = check_caller_callees::<Parameters>(&input, &parsed_blob)?;

    Ok((input, parsed_blob, caller))
}

pub fn commit<State>(
    env: impl GuestEnv,
    initial_state: State,
    contract_input: ContractInput,
    mut res: crate::RunResult<State>,
) where
    State: Digestable + BorshDeserialize,
{
    env.commit(&as_hyle_output(initial_state, contract_input, &mut res));
}
