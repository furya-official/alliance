package keeper_test

import (
	"github.com/furya-official/kaiju/x/kaiju/keeper"
	"github.com/furya-official/kaiju/x/kaiju/types"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestCreateKaiju(t *testing.T) {
	// GIVEN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx.WithBlockTime(startTime).WithBlockHeight(1)
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)
	rewardDuration := app.KaijuKeeper.RewardDelayTime(ctx)

	// WHEN
	createErr := app.KaijuKeeper.CreateKaiju(ctx, &types.MsgCreateKaijuProposal{
		Title:        "",
		Description:  "",
		Denom:        "ukaiju",
		RewardWeight: sdk.OneDec(),
		TakeRate:     sdk.OneDec(),
	})
	kaijusRes, kaijusErr := queryServer.Kaijus(ctx, &types.QueryKaijusRequest{})

	// THEN
	require.Nil(t, createErr)
	require.Nil(t, kaijusErr)
	require.Equal(t, kaijusRes, &types.QueryKaijusResponse{
		Kaijus: []types.KaijuAsset{
			{
				Denom:                "ukaiju",
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

func TestCreateKaijuFailWithDuplicatedDenom(t *testing.T) {
	// GIVEN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			types.NewKaijuAsset("ukaiju", sdk.NewDec(1), sdk.NewDec(0), startTime),
		},
	})

	// WHEN
	createErr := app.KaijuKeeper.CreateKaiju(ctx, &types.MsgCreateKaijuProposal{
		Title:        "",
		Description:  "",
		Denom:        "ukaiju",
		RewardWeight: sdk.OneDec(),
		TakeRate:     sdk.OneDec(),
	})

	// THEN
	require.Error(t, createErr)
}

func TestUpdateKaiju(t *testing.T) {
	// GIVEN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                "ukaiju",
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.OneDec(),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN
	updateErr := app.KaijuKeeper.UpdateKaiju(ctx, &types.MsgUpdateKaijuProposal{
		Title:                "",
		Description:          "",
		Denom:                "ukaiju",
		RewardWeight:         sdk.NewDec(6),
		TakeRate:             sdk.NewDec(7),
		RewardChangeInterval: 0,
		RewardChangeRate:     sdk.ZeroDec(),
	})
	kaijusRes, kaijusErr := queryServer.Kaijus(ctx, &types.QueryKaijusRequest{})

	// THEN
	require.Nil(t, updateErr)
	require.Nil(t, kaijusErr)
	require.Equal(t, kaijusRes, &types.QueryKaijusResponse{
		Kaijus: []types.KaijuAsset{
			{
				Denom:                "ukaiju",
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

func TestDeleteKaiju(t *testing.T) {
	// GIVEN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:        "ukaiju",
				RewardWeight: sdk.NewDec(2),
				TakeRate:     sdk.OneDec(),
				TotalTokens:  sdk.ZeroInt(),
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN
	deleteErr := app.KaijuKeeper.DeleteKaiju(ctx, &types.MsgDeleteKaijuProposal{
		Denom: "ukaiju",
	})
	kaijusRes, kaijusErr := queryServer.Kaijus(ctx, &types.QueryKaijusRequest{})

	// THEN
	require.Nil(t, deleteErr)
	require.Nil(t, kaijusErr)
	require.Equal(t, kaijusRes, &types.QueryKaijusResponse{
		Kaijus: nil,
		Pagination: &query.PageResponse{
			NextKey: nil,
			Total:   0,
		},
	})
}
