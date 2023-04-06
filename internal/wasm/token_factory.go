package wasm

import (
	"encoding/json"
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	"github.com/CosmWasm/token-factory/x/tokenfactory/bindings"
	bindingstypes "github.com/CosmWasm/token-factory/x/tokenfactory/bindings/types"
	tokenfactorykeeper "github.com/CosmWasm/token-factory/x/tokenfactory/keeper"
	tokenfactorytypes "github.com/CosmWasm/token-factory/x/tokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	TokenFactoryRoute = "token"
)

var _ Encoder = CustomTokenFactoryEncoder

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	tokenFactory *tokenfactorykeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := bindings.NewQueryPlugin(bank, tokenFactory)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: bindings.CustomQuerier(wasmQueryPlugin),
	})
	_ = wasmkeeper.WithMessageHandlerDecorator(bindings.CustomMessageDecorator(bank, tokenFactory))

	return []wasm.Option{
		queryPluginOpt,
		//messengerDecoratorOpt,
	}
}

func CustomTokenFactoryEncoder(contract sdk.AccAddress, data json.RawMessage, version string) ([]sdk.Msg, error) {
	var msg bindingstypes.TokenFactoryMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	tokenMsg := msg.Token
	if tokenMsg.CreateDenom != nil {
		return EncodeCreateToken(contract, msg.Token.CreateDenom)
	}
	if tokenMsg.MintTokens != nil {
		return EncodeMintToken(contract, msg.Token.MintTokens)
	}
	if tokenMsg.ChangeAdmin != nil {
		return EncodeChangeAdmin(contract, msg.Token.ChangeAdmin)
	}
	if tokenMsg.BurnTokens != nil {
		return EncodeBurnTokens(contract, msg.Token.BurnTokens)
	}
	if tokenMsg.SetMetadata != nil {
		return EncodeSetMetadata(contract, msg.Token.SetMetadata)
	}
	if tokenMsg.ForceTransfer != nil {
		return EncodeForceTransfer(contract, msg.Token.ForceTransfer)
	}

	return nil, fmt.Errorf("wasm: invalid custom token factory message")
}

func EncodeCreateToken(contract sdk.AccAddress, createDenom *bindingstypes.CreateDenom) ([]sdk.Msg, error) {
	createDenomMsg := tokenfactorytypes.NewMsgCreateDenom(contract.String(), createDenom.Subdenom)
	msgs := []sdk.Msg{createDenomMsg}

	if createDenom.Metadata != nil {
		setMetadataMsg := tokenfactorytypes.NewMsgSetDenomMetadata(contract.String(), transformMetadata(*createDenom.Metadata))
		msgs = append(msgs, setMetadataMsg)
	}
	return msgs, nil
}

func EncodeMintToken(contract sdk.AccAddress, mintTokens *bindingstypes.MintTokens) ([]sdk.Msg, error) {
	amount := sdk.NewCoin(mintTokens.Denom, mintTokens.Amount)
	rcpt, err := parseAddress(mintTokens.MintToAddress)
	if err != nil {
		return nil, err
	}

	mintTokensMsg := tokenfactorytypes.NewMsgMint(contract.String(), amount)
	bankSendMsg := banktypes.NewMsgSend(contract, rcpt, sdk.NewCoins(amount))
	return []sdk.Msg{mintTokensMsg, bankSendMsg}, nil
}

func EncodeChangeAdmin(contract sdk.AccAddress, changeAdmin *bindingstypes.ChangeAdmin) ([]sdk.Msg, error) {
	newAdminAddr, err := parseAddress(changeAdmin.NewAdminAddress)
	if err != nil {
		return nil, err
	}
	changeAdminMsg := tokenfactorytypes.NewMsgChangeAdmin(contract.String(), changeAdmin.Denom, newAdminAddr.String())
	return []sdk.Msg{changeAdminMsg}, nil
}

func EncodeSetMetadata(contract sdk.AccAddress, setMetadata *bindingstypes.SetMetadata) ([]sdk.Msg, error) {
	setMetadataMsg := tokenfactorytypes.NewMsgSetDenomMetadata(contract.String(), transformMetadata(setMetadata.Metadata))
	return []sdk.Msg{setMetadataMsg}, nil
}

func EncodeForceTransfer(contract sdk.AccAddress, forceTransfer *bindingstypes.ForceTransfer) ([]sdk.Msg, error) {
	amount := sdk.NewCoin(forceTransfer.Denom, forceTransfer.Amount)
	forceTransferMsg := tokenfactorytypes.NewMsgForceTransfer(contract.String(), amount, forceTransfer.FromAddress, forceTransfer.ToAddress)
	return []sdk.Msg{forceTransferMsg}, nil
}

func EncodeBurnTokens(contract sdk.AccAddress, burnTokens *bindingstypes.BurnTokens) ([]sdk.Msg, error) {
	amount := sdk.NewCoin(burnTokens.Denom, burnTokens.Amount)
	sdkMsg := tokenfactorytypes.NewMsgBurn(contract.String(), amount)
	return []sdk.Msg{sdkMsg}, nil
}

// parseAddress parses address from bech32 string and verifies its format.
func parseAddress(addr string) (sdk.AccAddress, error) {
	parsed, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "address from bech32")
	}
	err = sdk.VerifyAddressFormat(parsed)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "verify address format")
	}
	return parsed, nil
}

func transformMetadata(metadata bindingstypes.Metadata) banktypes.Metadata {
	denomUnits := make([]*banktypes.DenomUnit, len(metadata.DenomUnits))
	for _, denomUnit := range metadata.DenomUnits {
		denomUnits = append(denomUnits, &banktypes.DenomUnit{
			Denom:    denomUnit.Denom,
			Exponent: denomUnit.Exponent,
			Aliases:  denomUnit.Aliases,
		})
	}
	return banktypes.Metadata{
		Description: metadata.Description,
		DenomUnits:  denomUnits,
		Base:        metadata.Base,
		Display:     metadata.Display,
		Name:        metadata.Name,
		Symbol:      metadata.Symbol,
	}
}
