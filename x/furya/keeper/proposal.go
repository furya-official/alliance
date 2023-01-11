package keeper

import (
	"context"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/furya/x/furya/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CreateFurya(ctx context.Context, req *types.MsgCreateFuryaProposal) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	_, found := k.GetAssetByDenom(sdkCtx, req.Denom)

	if found {
		return status.Errorf(codes.AlreadyExists, "Asset with denom: %s already exists", req.Denom)
	}

	rewardStartTime := sdkCtx.BlockTime().Add(k.RewardDelayTime(sdkCtx))
	asset := types.FuryaAsset{
		Denom:                req.Denom,
		RewardWeight:         req.RewardWeight,
		TakeRate:             req.TakeRate,
		TotalTokens:          sdk.ZeroInt(),
		TotalValidatorShares: sdk.ZeroDec(),
		RewardStartTime:      rewardStartTime,
		RewardChangeRate:     req.RewardChangeRate,
		RewardChangeInterval: req.RewardChangeInterval,
		LastRewardChangeTime: rewardStartTime,
	}
	k.SetAsset(sdkCtx, asset)
	return nil
}

func (k Keeper) UpdateFurya(ctx context.Context, req *types.MsgUpdateFuryaProposal) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	asset, found := k.GetAssetByDenom(sdkCtx, req.Denom)

	if !found {
		return status.Errorf(codes.NotFound, "Asset with denom: %s does not exist", req.Denom)
	}

	asset.RewardWeight = req.RewardWeight
	asset.TakeRate = req.TakeRate
	asset.RewardChangeRate = req.RewardChangeRate
	asset.RewardChangeInterval = req.RewardChangeInterval

	err := k.UpdateFuryaAsset(sdkCtx, asset)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) DeleteFurya(ctx context.Context, req *types.MsgDeleteFuryaProposal) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	asset, found := k.GetAssetByDenom(sdkCtx, req.Denom)

	if !found {
		return status.Errorf(codes.NotFound, "Asset with denom: %s does not exist", req.Denom)
	}

	if asset.TotalTokens.GT(math.ZeroInt()) {
		return status.Errorf(codes.Internal, "Asset cannot be deleted because there are still %s delegations associated with it", asset.TotalTokens)
	}

	k.DeleteAsset(sdkCtx, req.Denom)

	return nil
}
