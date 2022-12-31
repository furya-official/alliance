package kaiju

import (
	"github.com/furya-official/kaiju/x/kaiju/keeper"
	"github.com/furya-official/kaiju/x/kaiju/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func NewKaijuProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.MsgCreateKaijuProposal:
			return k.CreateKaiju(ctx, c)
		case *types.MsgUpdateKaijuProposal:
			return k.UpdateKaiju(ctx, c)
		case *types.MsgDeleteKaijuProposal:
			return k.DeleteKaiju(ctx, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized kaiju proposal content type: %T", c)
		}
	}
}
