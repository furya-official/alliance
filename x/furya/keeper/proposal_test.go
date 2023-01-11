package keeper_test

import (
	"github.com/furya-official/furya/x/furya/keeper"
	"github.com/furya-official/furya/x/furya/types"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestCreateFurya(t *testing.T) {
	// GIVEN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx.WithBlockTime(startTime).WithBlockHeight(1)
	queryServer := keeper.NewQueryServerImpl(app.FuryaKeeper)
	rewardDuration := app.FuryaKeeper.RewardDelayTime(ctx)

	// WHEN
	createErr := app.FuryaKeeper.CreateFurya(ctx, &types.MsgCreateFuryaProposal{
		Title:        "",
		Description:  "",
		Denom:        "ufury",
		RewardWeight: sdk.OneDec(),
		TakeRate:     sdk.OneDec(),
	})
	furyasRes, furyasErr := queryServer.Furyas(ctx, &types.QueryFuryasRequest{})

	// THEN
	require.Nil(t, createErr)
	require.Nil(t, furyasErr)
	require.Equal(t, furyasRes, &types.QueryFuryasResponse{
		Furyas: []types.FuryaAsset{
			{
				Denom:                "ufury",
				RewardWeight:         sdk.NewDec(1),
				TakeRate:             sdk.NewDec(1),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardStartTime:      ctx.BlockTime().Add(rewardDuration),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
				LastRewardChangeTime: ctx.BlockTime().Add(rewardDuration),
			},
		},
		Pagination: &query.PageResponse{
			NextKey: nil,
			Total:   1,
		},
	})
}

func TestCreateFuryaFailWithDuplicatedDenom(t *testing.T) {
	// GIVEN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.FuryaKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.FuryaAsset{
			types.NewFuryaAsset("ufury", sdk.NewDec(1), sdk.NewDec(0), startTime),
		},
	})

	// WHEN
	createErr := app.FuryaKeeper.CreateFurya(ctx, &types.MsgCreateFuryaProposal{
		Title:        "",
		Description:  "",
		Denom:        "ufury",
		RewardWeight: sdk.OneDec(),
		TakeRate:     sdk.OneDec(),
	})

	// THEN
	require.Error(t, createErr)
}

func TestUpdateFurya(t *testing.T) {
	// GIVEN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.FuryaKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.FuryaAsset{
			{
				Denom:                "ufury",
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.OneDec(),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.FuryaKeeper)

	// WHEN
	updateErr := app.FuryaKeeper.UpdateFurya(ctx, &types.MsgUpdateFuryaProposal{
		Title:                "",
		Description:          "",
		Denom:                "ufury",
		RewardWeight:         sdk.NewDec(6),
		TakeRate:             sdk.NewDec(7),
		RewardChangeInterval: 0,
		RewardChangeRate:     sdk.ZeroDec(),
	})
	furyasRes, furyasErr := queryServer.Furyas(ctx, &types.QueryFuryasRequest{})

	// THEN
	require.Nil(t, updateErr)
	require.Nil(t, furyasErr)
	require.Equal(t, furyasRes, &types.QueryFuryasResponse{
		Furyas: []types.FuryaAsset{
			{
				Denom:                "ufury",
				RewardWeight:         sdk.NewDec(6),
				TakeRate:             sdk.NewDec(7),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
		Pagination: &query.PageResponse{
			NextKey: nil,
			Total:   1,
		},
	})
}

func TestDeleteFurya(t *testing.T) {
	// GIVEN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.FuryaKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.FuryaAsset{
			{
				Denom:        "ufury",
				RewardWeight: sdk.NewDec(2),
				TakeRate:     sdk.OneDec(),
				TotalTokens:  sdk.ZeroInt(),
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.FuryaKeeper)

	// WHEN
	deleteErr := app.FuryaKeeper.DeleteFurya(ctx, &types.MsgDeleteFuryaProposal{
		Denom: "ufury",
	})
	furyasRes, furyasErr := queryServer.Furyas(ctx, &types.QueryFuryasRequest{})

	// THEN
	require.Nil(t, deleteErr)
	require.Nil(t, furyasErr)
	require.Equal(t, furyasRes, &types.QueryFuryasResponse{
		Furyas: nil,
		Pagination: &query.PageResponse{
			NextKey: nil,
			Total:   0,
		},
	})
}
