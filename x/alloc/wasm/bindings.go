package wasm

import (
	"encoding/json"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	allocKeeper "github.com/public-awesome/stargaze/v8/x/alloc/keeper"
	allocTypes "github.com/public-awesome/stargaze/v8/x/alloc/types"
)

func CustomMessageDecorators(allocKeeper *allocKeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:     old,
			allocKeeper: allocKeeper,
		}
	}
}

type CustomMessenger struct {
	wrapped     wasmkeeper.Messenger
	allocKeeper *allocKeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		var allocModMsg AllocModuleMsg
		if err := json.Unmarshal(msg.Custom, &allocModMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "alloc module msg")
		}
		if allocModMsg.Alloc == nil {
			return nil, nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "nil alloc field")
		}

		if allocModMsg.Alloc.FundFairburnPool != nil {
			m.fundFairburnPool(ctx, contractAddr, allocModMsg.Alloc.FundFairburnPool)
		}
	}

	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

func (m *CustomMessenger) fundFairburnPool(ctx sdk.Context, contractAddr sdk.AccAddress, fundFairburnPool *FundFairburnPool) ([]sdk.Event, [][]byte, error) {
	msgServer := allocKeeper.NewMsgServerImpl(*m.allocKeeper)
	amount, err := wasmkeeper.ConvertWasmCoinsToSdkCoins(fundFairburnPool.Amount)
	if err != nil {
		return nil, nil, err
	}

	msgFundFairburnPool := allocTypes.NewMsgFundFairburnPool(sdk.AccAddress(contractAddr.String()), amount)
	if err := msgFundFairburnPool.ValidateBasic(); err != nil {
		return nil, nil, sdkerrors.Wrap(err, "failed validating NewMsgFundFairburnPool")
	}

	resp, err := msgServer.FundFairburnPool(sdk.WrapSDKContext(ctx), msgFundFairburnPool)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "funding fairburn pool")
	}

	res, err := resp.Marshal()
	if err != nil {
		return nil, nil, err
	}

	return nil, [][]byte{res}, nil
}
