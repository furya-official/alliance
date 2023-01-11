package furya

import (
	"github.com/furya-official/furya/x/furya/keeper"
	"github.com/furya-official/furya/x/furya/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func NewFuryaProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.MsgCreateFuryaProposal:
			return k.CreateFurya(ctx, c)
		case *types.MsgUpdateFuryaProposal:
			return k.UpdateFurya(ctx, c)
		case *types.MsgDeleteFuryaProposal:
			return k.DeleteFurya(ctx, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized furya proposal content type: %T", c)
		}
	}
}
