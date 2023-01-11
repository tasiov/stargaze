package wasm

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	distributionKeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
)

func CustomMessageDecorators(distrKeeper *distributionKeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:     old,
			distrKeeper: distrKeeper,
		}
	}
}

type CustomMessenger struct {
	wrapped     wasmkeeper.Messenger
	distrKeeper *distributionKeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	return nil, nil, nil
}
