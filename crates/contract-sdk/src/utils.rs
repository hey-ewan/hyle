use crate::{
    alloc::string::{String, ToString},
    guest::fail,
    Identity, StructuredBlobData,
};
use borsh::{BorshDeserialize, BorshSerialize};
use core::result::Result;

use hyle_model::{
    flatten_blobs, Blob, BlobIndex, ContractInput, Digestable, DropEndOfReader, HyleOutput,
    StructuredBlob,
};

pub fn parse_blob<Parameters>(blobs: &[Blob], index: &BlobIndex) -> Option<Parameters>
where
    Parameters: BorshDeserialize,
{
    let blob = match blobs.get(index.0) {
        Some(v) => v,
        None => {
            return None;
        }
    };

    let Ok(parameters) = borsh::from_slice::<Parameters>(blob.data.0.as_slice()) else {
        return None;
    };

    Some(parameters)
}

pub fn parse_structured_blob<Parameters>(
    blobs: &[Blob],
    index: &BlobIndex,
) -> Option<StructuredBlob<Parameters>>
where
    Parameters: BorshDeserialize,
{
    let blob = match blobs.get(index.0) {
        Some(v) => v,
        None => {
            return None;
        }
    };

    let parsed_blob: StructuredBlob<Parameters> = StructuredBlob::try_from(blob.clone())
        .unwrap_or_else(|e| {
            panic!("Failed to decode blob: {:?}", e);
        });
    Some(parsed_blob)
}

pub fn as_hyle_output<State: Digestable + BorshDeserialize>(
    initial_state: State,
    contract_input: ContractInput,
    res: &mut crate::RunResult<State>,
) -> HyleOutput {
    match res {
        Ok(res) => HyleOutput {
            version: 1,
            initial_state: initial_state.as_digest(),
            next_state: res.1.as_digest(),
            identity: contract_input.identity,
            index: contract_input.index,
            blobs: flatten_blobs(&contract_input.blobs),
            success: true,
            tx_hash: contract_input.tx_hash,
            tx_ctx: contract_input.tx_ctx,
            registered_contracts: core::mem::take(&mut res.2),
            program_outputs: core::mem::take(&mut res.0).into_bytes(),
        },
        Err(message) => fail(contract_input, initial_state, message),
    }
}

pub fn check_caller_callees<Paramaters>(
    input: &ContractInput,
    parameters: &StructuredBlob<Paramaters>,
) -> Result<Identity, String>
where
    Paramaters: BorshSerialize + BorshDeserialize,
{
    // Check that callees has this blob as caller
    if let Some(callees) = parameters.data.callees.as_ref() {
        for callee_index in callees {
            let callee_blob = input.blobs[callee_index.0].clone();
            let callee_structured_blob: StructuredBlobData<DropEndOfReader> =
                callee_blob.data.try_into().expect("Failed to decode blob");
            if callee_structured_blob.caller != Some(input.index) {
                return Err("One Callee does not have this blob as caller".to_string());
            }
        }
    }
    // Extract the correct caller
    if let Some(caller_index) = parameters.data.caller.as_ref() {
        let caller_blob = input.blobs[caller_index.0].clone();
        let caller_structured_blob: StructuredBlobData<DropEndOfReader> =
            caller_blob.data.try_into().expect("Failed to decode blob");
        // Check that caller has this blob as callee
        if caller_structured_blob.callees.is_some()
            && !caller_structured_blob
                .callees
                .unwrap()
                .contains(&input.index)
        {
            return Err("Incorrect Caller for this blob".to_string());
        }
        return Ok(caller_blob.contract_name.0.clone().into());
    }

    // No callers detected, use the identity
    Ok(input.identity.clone())
}
