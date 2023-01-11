package wasm

import (
	"encoding/json"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	claimKeeper "github.com/public-awesome/stargaze/v8/x/claim/keeper"
	claimtypes "github.com/public-awesome/stargaze/v8/x/claim/types"
)

func CustomMessageDecorators(allocKeeper *claimKeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:     old,
			claimKeeper: allocKeeper,
		}
	}
}

type CustomMessenger struct {
	wrapped     wasmkeeper.Messenger
	claimKeeper *claimKeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		var claimModMsg ClaimModuleMsg
		if err := json.Unmarshal(msg.Custom, &claimModMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "claim module msg")
		}
		if claimModMsg.Claim == nil {
			return nil, nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "nil claim field")
		}

		if claimModMsg.Claim.ClaimFor != nil {
			m.fundFairburnPool(ctx, contractAddr, claimModMsg.Claim.ClaimFor)
		}
	}

	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

func (m *CustomMessenger) fundFairburnPool(ctx sdk.Context, contractAddr sdk.AccAddress, claimFor *ClaimFor) ([]sdk.Event, [][]byte, error) {
	msgServer := claimKeeper.NewMsgServerImpl(*m.claimKeeper)

	action, err := claimFor.Action.ToAction()
	if err != nil {
		return nil, nil, err
	}
	msgClaimFor := claimtypes.NewMsgClaimFor(contractAddr.String(), claimFor.Address, action)
	if err := msgClaimFor.ValidateBasic(); err != nil {
		return nil, nil, sdkerrors.Wrap(err, "failed validating MsgClaimFor")
	}

	resp, err := msgServer.ClaimFor(sdk.WrapSDKContext(ctx), msgClaimFor)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "claim for action")
	}

	res, err := resp.Marshal()
	if err != nil {
		return nil, nil, err
	}

	return nil, [][]byte{res}, nil
}
