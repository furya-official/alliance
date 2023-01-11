package keeper_test

import (
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	test_helpers "github.com/furya-official/furya/app"
	"github.com/furya-official/furya/x/furya/types"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	app, ctx := createTestContext(t)
	app.FuryaKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.Params{
			RewardDelayTime:       time.Duration(1000000),
			TakeRateClaimInterval: time.Duration(1000000),
			LastTakeRateClaimTime: time.Unix(0, 0).UTC(),
		},
		Assets: []types.FuryaAsset{
			types.NewFuryaAsset("stake", sdk.NewDec(1), sdk.ZeroDec(), ctx.BlockTime()),
		},
	})

	delay := app.FuryaKeeper.RewardDelayTime(ctx)
	require.Equal(t, time.Duration(1000000), delay)

	interval := app.FuryaKeeper.RewardClaimInterval(ctx)
	require.Equal(t, time.Duration(1000000), interval)

	lastClaimTime := app.FuryaKeeper.LastRewardClaimTime(ctx)
	require.Equal(t, time.Unix(0, 0).UTC(), lastClaimTime)

	assets := app.FuryaKeeper.GetAllAssets(ctx)
	require.Equal(t, 1, len(assets))
	require.Equal(t, &types.FuryaAsset{
		Denom:                "stake",
		RewardWeight:         sdk.NewDec(1.0),
		TakeRate:             sdk.NewDec(0.0),
		TotalTokens:          sdk.ZeroInt(),
		TotalValidatorShares: sdk.ZeroDec(),
		RewardStartTime:      ctx.BlockTime(),
		RewardChangeRate:     sdk.OneDec(),
		RewardChangeInterval: 0,
	}, assets[0])
}

func TestExportAndImportGenesis(t *testing.T) {
	app, ctx := createTestContext(t)
	ctx = ctx.WithBlockTime(time.Now()).WithBlockHeight(1)
	app.FuryaKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.Params{
			RewardDelayTime:       time.Duration(1000000),
			TakeRateClaimInterval: time.Duration(1000000),
			LastTakeRateClaimTime: time.Unix(0, 0).UTC(),
		},
		Assets: []types.FuryaAsset{},
	})

	// All the addresses needed
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	require.Len(t, delegations, 1)
	delAddr, err := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	require.NoError(t, err)
	valAddr, err := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	require.NoError(t, err)
	val1, err := app.FuryaKeeper.GetFuryaValidator(ctx, valAddr)
	require.NoError(t, err)
	addrs := test_helpers.AddTestAddrsIncremental(app, ctx, 3, sdk.NewCoins(
		sdk.NewCoin(FURYA_TOKEN_DENOM, sdk.NewInt(1000_000)),
		sdk.NewCoin(FURYA_2_TOKEN_DENOM, sdk.NewInt(1000_000)),
	))
	valAddr2 := sdk.ValAddress(addrs[0])
	_val2 := teststaking.NewValidator(t, valAddr2, test_helpers.CreateTestPubKeys(1)[0])
	test_helpers.RegisterNewValidator(t, app, ctx, _val2)
	val2, err := app.FuryaKeeper.GetFuryaValidator(ctx, valAddr2)
	require.NoError(t, err)

	// Add furya asset
	err = app.FuryaKeeper.CreateFurya(ctx, &types.MsgCreateFuryaProposal{
		Title:                "",
		Description:          "",
		Denom:                FURYA_TOKEN_DENOM,
		RewardWeight:         sdk.NewDec(1),
		TakeRate:             sdk.NewDec(0),
		RewardChangeRate:     sdk.MustNewDecFromStr("0.5"),
		RewardChangeInterval: time.Hour * 24,
	})
	require.NoError(t, err)

	// Delegate
	delegationCoin := sdk.NewCoin(FURYA_TOKEN_DENOM, sdk.NewInt(1000_000_000))
	err = app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(delegationCoin))
	require.NoError(t, err)
	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(delegationCoin))
	require.NoError(t, err)
	_, err = app.FuryaKeeper.Delegate(ctx, delAddr, val1, delegationCoin)
	require.NoError(t, err)

	// Redelegate
	_, err = app.FuryaKeeper.Redelegate(ctx, delAddr, val1, val2, sdk.NewCoin(FURYA_TOKEN_DENOM, sdk.NewInt(500_000_000)))
	require.NoError(t, err)

	// Undelegate
	_, err = app.FuryaKeeper.Undelegate(ctx, delAddr, val1, sdk.NewCoin(FURYA_TOKEN_DENOM, sdk.NewInt(500_000_000)))
	require.NoError(t, err)

	// Trigger update asset
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Hour * 25)).WithBlockHeight(ctx.BlockHeight() + 1)
	err = app.FuryaKeeper.UpdateFuryaAsset(ctx, types.NewFuryaAsset(FURYA_TOKEN_DENOM, sdk.MustNewDecFromStr("0.5"), sdk.ZeroDec(), ctx.BlockTime()))
	require.NoError(t, err)

	genesisState := app.FuryaKeeper.ExportGenesis(ctx)
	require.NotNil(t, genesisState.Params)
	require.Greater(t, len(genesisState.Assets), 0)
	require.Greater(t, len(genesisState.ValidatorInfos), 0)
	require.Greater(t, len(genesisState.Delegations), 0)
	require.Greater(t, len(genesisState.Undelegations), 0)
	require.Greater(t, len(genesisState.Redelegations), 0)
	require.Greater(t, len(genesisState.RewardWeightChangeSnaphots), 0)

	store := ctx.KVStore(app.FuryaKeeper.StoreKey())
	iter := store.Iterator(nil, nil)

	// Init a new app
	app, ctx = createTestContext(t)
	ctx = ctx.WithBlockTime(time.Now()).WithBlockHeight(1)

	app.FuryaKeeper.InitGenesis(ctx, genesisState)

	// Check all items in the furya store match
	iter2 := store.Iterator(nil, nil)
	for ; iter.Valid(); iter.Next() {
		require.Equal(t, iter.Key(), iter2.Key())
		require.Equal(t, iter.Value(), iter2.Value())
		iter2.Next()
	}
}
