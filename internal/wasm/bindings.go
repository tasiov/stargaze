package wasm

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	allocKeeper "github.com/public-awesome/stargaze/v8/x/alloc/keeper"
	claimKeeper "github.com/public-awesome/stargaze/v8/x/claim/keeper"
)

func CustomMessageDecorators(allocKeeper *allocKeeper.Keeper, claimKeeper *claimKeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:     old,
			allocKeeper: allocKeeper,
			claimKeeper: claimKeeper,
		}
	}
}

type CustomMessenger struct {
	wrapped     wasmkeeper.Messenger
	allocKeeper *allocKeeper.Keeper
	claimKeeper *claimKeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	return nil, nil, nil
}
