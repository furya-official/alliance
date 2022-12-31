package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	"testing"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	test_helpers "github.com/furya-official/kaiju/app"
	"github.com/furya-official/kaiju/x/kaiju/keeper"
	"github.com/furya-official/kaiju/x/kaiju/types"
)

var UKAIJU_KAIJU = "ukaiju"

func TestQueryKaijus(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH KAIJUS ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:        KAIJU_TOKEN_DENOM,
				RewardWeight: sdk.NewDec(2),
				TakeRate:     sdk.NewDec(0),
				TotalTokens:  sdk.ZeroInt(),
			},
			{
				Denom:        KAIJU_2_TOKEN_DENOM,
				RewardWeight: sdk.NewDec(10),
				TakeRate:     sdk.MustNewDecFromStr("0.14159265359"),
				TotalTokens:  sdk.ZeroInt(),
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN: QUERYING THE KAIJUS LIST
	kaijus, err := queryServer.Kaijus(ctx, &types.QueryKaijusRequest{})

	// THEN: VALIDATE THAT BOTH KAIJUS HAVE THE CORRECT MODEL WHEN QUERYING
	require.Nil(t, err)
	require.Equal(t, &types.QueryKaijusResponse{
		Kaijus: []types.KaijuAsset{
			{
				Denom:                "kaiju",
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.NewDec(0),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
			{
				Denom:                "kaiju2",
				RewardWeight:         sdk.NewDec(10),
				TakeRate:             sdk.MustNewDecFromStr("0.14159265359"),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
		Pagination: &query.PageResponse{
			NextKey: nil,
			Total:   2,
		},
	}, kaijus)
}

func TestQueryAnUniqueKaiju(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH KAIJUS ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                KAIJU_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.NewDec(0),
				TotalTokens:          sdk.ZeroInt(),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
			{
				Denom:                KAIJU_2_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(10),
				TakeRate:             sdk.MustNewDecFromStr("0.14159265359"),
				TotalTokens:          sdk.ZeroInt(),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN: QUERYING THE KAIJUS LIST
	kaijus, err := queryServer.Kaiju(ctx, &types.QueryKaijuRequest{
		Denom: "kaiju2",
	})

	// THEN: VALIDATE THAT BOTH KAIJUS HAVE THE CORRECT MODEL WHEN QUERYING
	require.Nil(t, err)
	require.Equal(t, &types.QueryKaijuResponse{
		Kaiju: &types.KaijuAsset{
			Denom:                "kaiju2",
			RewardWeight:         sdk.NewDec(10),
			TakeRate:             sdk.MustNewDecFromStr("0.14159265359"),
			TotalTokens:          sdk.ZeroInt(),
			TotalValidatorShares: sdk.NewDec(0),
			RewardChangeRate:     sdk.NewDec(0),
			RewardChangeInterval: 0,
		},
	}, kaijus)
}

func TestQueryAnUniqueIBCKaiju(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH KAIJUS ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                "ibc/" + KAIJU_2_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(10),
				TakeRate:             sdk.MustNewDecFromStr("0.14159265359"),
				TotalTokens:          sdk.ZeroInt(),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN: QUERYING THE KAIJUS LIST
	kaijus, err := queryServer.IBCKaiju(ctx, &types.QueryIBCKaijuRequest{
		Hash: "kaiju2",
	})

	// THEN: VALIDATE THAT BOTH KAIJUS HAVE THE CORRECT MODEL WHEN QUERYING
	require.Nil(t, err)
	require.Equal(t, &types.QueryKaijuResponse{
		Kaiju: &types.KaijuAsset{
			Denom:                "ibc/kaiju2",
			RewardWeight:         sdk.NewDec(10),
			TakeRate:             sdk.MustNewDecFromStr("0.14159265359"),
			TotalTokens:          sdk.ZeroInt(),
			TotalValidatorShares: sdk.NewDec(0),
			RewardChangeRate:     sdk.NewDec(0),
			RewardChangeInterval: 0,
		},
	}, kaijus)
}

func TestQueryKaijuNotFound(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN: QUERYING THE KAIJU
	_, err := queryServer.Kaiju(ctx, &types.QueryKaijuRequest{
		Denom: "kaiju2",
	})

	// THEN: VALIDATE THE ERROR
	require.Equal(t, err.Error(), "kaiju asset is not whitelisted")
}

func TestQueryAllKaijus(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN: QUERYING THE KAIJU
	res, err := queryServer.Kaijus(ctx, &types.QueryKaijusRequest{})

	// THEN: VALIDATE THE ERROR
	require.Nil(t, err)
	require.Equal(t, len(res.Kaijus), 0)
	require.Equal(t, res.Pagination, &query.PageResponse{
		NextKey: nil,
		Total:   0,
	})
}

func TestQueryParams(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH AN KAIJU ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                KAIJU_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.NewDec(0),
				TotalTokens:          sdk.ZeroInt(),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN: QUERYING THE PARAMS...
	queyParams, err := queryServer.Params(ctx, &types.QueryParamsRequest{})

	// THEN: VALIDATE THAT NO ERRORS HAVE BEEN PRODUCED AND OUTPUT IS AS WE EXPECT
	require.Nil(t, err)

	require.Equal(t, queyParams.Params.RewardDelayTime, time.Hour)
	require.Equal(t, queyParams.Params.TakeRateClaimInterval, time.Minute*5)

	// there is no way to match the exact time when the module is being instantiated
	// but we know that this time should be older than actually the time when this
	// following two lines are executed
	require.NotNil(t, queyParams.Params.LastTakeRateClaimTime)
	require.LessOrEqual(t, queyParams.Params.LastTakeRateClaimTime, time.Now())
}

func TestClaimQueryReward(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH ACCOUNTS
	app, ctx := createTestContext(t)
	startTime := time.Now().UTC()
	ctx = ctx.WithBlockTime(startTime)
	ctx = ctx.WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.Params{
			RewardDelayTime:       time.Minute * 60,
			TakeRateClaimInterval: time.Minute * 5,
			LastTakeRateClaimTime: startTime,
		},
		Assets: []types.KaijuAsset{
			{
				Denom:                UKAIJU_KAIJU,
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.MustNewDecFromStr("0.00005"),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)
	feeCollectorAddr := app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	valAddr, _ := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	val1, _ := app.KaijuKeeper.GetKaijuValidator(ctx, valAddr)
	delAddr := test_helpers.AddTestAddrsIncremental(app, ctx, 1, sdk.NewCoins(sdk.NewCoin(UKAIJU_KAIJU, sdk.NewInt(1000_000_000))))[0]

	// WHEN: DELEGATING ...
	delRes, delErr := app.KaijuKeeper.Delegate(ctx, delAddr, val1, sdk.NewCoin(UKAIJU_KAIJU, sdk.NewInt(1000_000_000)))
	require.Nil(t, delErr)
	require.Equal(t, sdk.NewDec(1000000000), *delRes)
	assets := app.KaijuKeeper.GetAllAssets(ctx)
	err := app.KaijuKeeper.RebalanceBondTokenWeights(ctx, assets)
	require.NoError(t, err)

	// ...and advance block...
	timePassed := time.Minute*5 + time.Second
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(timePassed))
	ctx = ctx.WithBlockHeight(2)
	app.KaijuKeeper.DeductAssetsHook(ctx, assets)
	app.BankKeeper.GetAllBalances(ctx, feeCollectorAddr)
	require.Equal(t, startTime.Add(time.Minute*5), app.KaijuKeeper.LastRewardClaimTime(ctx))
	app.KaijuKeeper.GetAssetByDenom(ctx, UKAIJU_KAIJU)

	// ... at the next begin block, tokens will be distributed from the fee pool...
	cons, _ := val1.GetConsAddr()
	app.DistrKeeper.AllocateTokens(ctx, 1, 1, cons, []abcitypes.VoteInfo{
		{
			Validator: abcitypes.Validator{
				Address: cons,
				Power:   1,
			},
			SignedLastBlock: true,
		},
	})

	// THEN: Query the delegation rewards ...
	queryDelegation, queryErr := queryServer.KaijuDelegationRewards(ctx, &types.QueryKaijuDelegationRewardsRequest{
		DelegatorAddr: delAddr.String(),
		ValidatorAddr: valAddr.String(),
		Denom:         UKAIJU_KAIJU,
	})

	// ... validate that no error has been produced.
	require.Nil(t, queryErr)
	require.Equal(t, &types.QueryKaijuDelegationRewardsResponse{
		Rewards: []sdk.Coin{
			{
				Denom:  UKAIJU_KAIJU,
				Amount: math.NewInt(32666),
			},
		},
	}, queryDelegation)
}

func TestQueryKaijuDelegation(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH KAIJUS ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                KAIJU_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.NewDec(0),
				TotalTokens:          sdk.ZeroInt(),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	delAddr, _ := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	valAddr, _ := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	val, _ := app.KaijuKeeper.GetKaijuValidator(ctx, valAddr)
	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))

	// WHEN: DELEGATING AND QUERYING ...
	delegationTxRes, txErr := app.KaijuKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(1000_000)))
	queryDelegation, queryErr := queryServer.KaijuDelegation(ctx, &types.QueryKaijuDelegationRequest{
		DelegatorAddr: delAddr.String(),
		ValidatorAddr: val.OperatorAddress,
		Denom:         KAIJU_TOKEN_DENOM,
	})

	// THEN: VALIDATE THAT NO ERRORS HAVE BEEN PRODUCED AND BOTH OUTPUTS ARE AS WE EXPECT
	require.Nil(t, txErr)
	require.Nil(t, queryErr)
	require.Equal(t, &types.QueryKaijuDelegationResponse{
		Delegation: types.DelegationResponse{
			Delegation: types.Delegation{
				DelegatorAddress:      delAddr.String(),
				ValidatorAddress:      val.OperatorAddress,
				Denom:                 KAIJU_TOKEN_DENOM,
				Shares:                sdk.NewDec(1000_000),
				RewardHistory:         nil,
				LastRewardClaimHeight: uint64(ctx.BlockHeight()),
			},
			Balance: sdk.Coin{
				Denom:  KAIJU_TOKEN_DENOM,
				Amount: sdk.NewInt(1000_000),
			},
		},
	}, queryDelegation)
	require.Equal(t, sdk.NewDec(1000000), *delegationTxRes)
}

func TestQueryKaijuDelegationNotFound(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	delAddr, _ := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	valAddr, _ := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	val, _ := app.StakingKeeper.GetValidator(ctx, valAddr)
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN: QUERYING ...
	_, err := queryServer.KaijuDelegation(ctx, &types.QueryKaijuDelegationRequest{
		DelegatorAddr: delAddr.String(),
		ValidatorAddr: val.OperatorAddress,
		Denom:         KAIJU_TOKEN_DENOM,
	})

	// THEN: VALIDATE THAT NO ERRORS HAVE BEEN PRODUCED AND BOTH OUTPUTS ARE AS WE EXPECT
	require.Equal(t, err, status.Error(codes.NotFound, "KaijuAsset not found by denom kaiju"))
}

func TestQueryKaijuValidatorNotFound(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	delAddr, _ := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN: QUERYING ...
	_, err := queryServer.KaijuDelegation(ctx, &types.QueryKaijuDelegationRequest{
		DelegatorAddr: delAddr.String(),
		ValidatorAddr: "cosmosvaloper19lss6zgdh5vvcpjhfftdghrpsw7a4434elpwpu",
		Denom:         KAIJU_TOKEN_DENOM,
	})

	// THEN: VALIDATE THAT NO ERRORS HAVE BEEN PRODUCED AND BOTH OUTPUTS ARE AS WE EXPECT
	require.Equal(t, err, status.Error(codes.NotFound, "Validator not found by address cosmosvaloper19lss6zgdh5vvcpjhfftdghrpsw7a4434elpwpu"))
}

func TestQueryKaijusDelegationByValidator(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH KAIJUS ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                KAIJU_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.NewDec(0),
				TotalTokens:          sdk.ZeroInt(),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	delAddr, _ := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	valAddr, _ := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	val, _ := app.KaijuKeeper.GetKaijuValidator(ctx, valAddr)
	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))

	// WHEN: DELEGATING AND QUERYING ...
	delegationTxRes, txErr := app.KaijuKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(1000_000)))
	queryDelegation, queryErr := queryServer.KaijusDelegationByValidator(ctx, &types.QueryKaijusDelegationByValidatorRequest{
		DelegatorAddr: delAddr.String(),
		ValidatorAddr: val.OperatorAddress,
	})

	// THEN: VALIDATE THAT NO ERRORS HAVE BEEN PRODUCED AND BOTH OUTPUTS ARE AS WE EXPECT
	require.Nil(t, txErr)
	require.Nil(t, queryErr)
	require.Equal(t, &types.QueryKaijusDelegationsResponse{
		Delegations: []types.DelegationResponse{
			{
				Delegation: types.Delegation{
					DelegatorAddress:      delAddr.String(),
					ValidatorAddress:      val.OperatorAddress,
					Denom:                 KAIJU_TOKEN_DENOM,
					Shares:                sdk.NewDec(1000_000),
					RewardHistory:         nil,
					LastRewardClaimHeight: uint64(ctx.BlockHeight()),
				},
				Balance: sdk.Coin{
					Denom:  KAIJU_TOKEN_DENOM,
					Amount: sdk.NewInt(1000_000),
				},
			},
		},
		Pagination: &query.PageResponse{
			NextKey: nil,
			Total:   1,
		},
	}, queryDelegation)
	require.Equal(t, sdk.NewDec(1000_000), *delegationTxRes)
}

func TestQueryKaijusDelegationByValidatorNotFound(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	delAddr, _ := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)

	// WHEN: QUERYING ...
	_, err := queryServer.KaijusDelegationByValidator(ctx, &types.QueryKaijusDelegationByValidatorRequest{
		DelegatorAddr: delAddr.String(),
		ValidatorAddr: "cosmosvaloper19lss6zgdh5vvcpjhfftdghrpsw7a4434elpwpu",
	})

	// THEN: VALIDATE THAT NO ERRORS HAVE BEEN PRODUCED AND BOTH OUTPUTS ARE AS WE EXPECT
	require.Equal(t, err, status.Error(codes.NotFound, "Validator not found by address cosmosvaloper19lss6zgdh5vvcpjhfftdghrpsw7a4434elpwpu"))
}

func TestQueryKaijusKaijusDelegation(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH KAIJUS ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                KAIJU_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.NewDec(0),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
			{
				Denom:                KAIJU_2_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(10),
				TakeRate:             sdk.MustNewDecFromStr("0.14159265359"),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	delAddr, _ := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	valAddr, _ := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	val, _ := app.KaijuKeeper.GetKaijuValidator(ctx, valAddr)
	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(2000_000))))

	// WHEN: DELEGATING AND QUERYING ...
	delegationTxRes, txErr := app.KaijuKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(1000_000)))
	delegation2TxRes, tx2Err := app.KaijuKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(1000_000)))
	queryDelegation, queryErr := queryServer.KaijusDelegation(ctx, &types.QueryKaijusDelegationsRequest{
		DelegatorAddr: delAddr.String(),
	})

	// THEN: VALIDATE THAT NO ERRORS HAVE BEEN PRODUCED AND BOTH OUTPUTS ARE AS WE EXPECT
	require.Nil(t, txErr)
	require.Nil(t, tx2Err)
	require.Nil(t, queryErr)
	require.Equal(t, &types.QueryKaijusDelegationsResponse{
		Delegations: []types.DelegationResponse{
			{
				Delegation: types.Delegation{
					DelegatorAddress:      delAddr.String(),
					ValidatorAddress:      val.OperatorAddress,
					Denom:                 KAIJU_TOKEN_DENOM,
					Shares:                sdk.NewDec(1000_000),
					RewardHistory:         nil,
					LastRewardClaimHeight: uint64(ctx.BlockHeight()),
				},
				Balance: sdk.Coin{
					Denom:  KAIJU_TOKEN_DENOM,
					Amount: sdk.NewInt(1000_000),
				},
			},
			{
				Delegation: types.Delegation{
					DelegatorAddress:      delAddr.String(),
					ValidatorAddress:      val.OperatorAddress,
					Denom:                 KAIJU_2_TOKEN_DENOM,
					Shares:                sdk.NewDec(1000_000),
					RewardHistory:         nil,
					LastRewardClaimHeight: uint64(ctx.BlockHeight()),
				},
				Balance: sdk.Coin{
					Denom:  KAIJU_2_TOKEN_DENOM,
					Amount: sdk.NewInt(1000_000),
				},
			},
		},
		Pagination: &query.PageResponse{
			NextKey: nil,
			Total:   2,
		},
	}, queryDelegation)
	require.Equal(t, sdk.NewDec(1000_000), *delegationTxRes)
	require.Equal(t, sdk.NewDec(1000_000), *delegation2TxRes)
}

func TestQueryAllDelegations(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH KAIJUS ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                KAIJU_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.NewDec(0),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
			{
				Denom:                KAIJU_2_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(10),
				TakeRate:             sdk.MustNewDecFromStr("0.14159265359"),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	delAddr, _ := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	valAddr, _ := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	val, _ := app.KaijuKeeper.GetKaijuValidator(ctx, valAddr)
	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(2000_000))))

	// WHEN: DELEGATING AND QUERYING ...
	_, txErr := app.KaijuKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.NoError(t, txErr)
	_, tx2Err := app.KaijuKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.NoError(t, tx2Err)
	queryDelegations, queryErr := queryServer.AllKaijusDelegations(ctx, &types.QueryAllKaijusDelegationsRequest{
		Pagination: &query.PageRequest{
			Key:        nil,
			Offset:     0,
			Limit:      1,
			CountTotal: false,
			Reverse:    false,
		},
	})
	require.NoError(t, queryErr)
	require.Equal(t, 1, len(queryDelegations.Delegations))

	require.Equal(t, types.DelegationResponse{
		Delegation: types.Delegation{
			DelegatorAddress:      delAddr.String(),
			ValidatorAddress:      val.OperatorAddress,
			Denom:                 KAIJU_TOKEN_DENOM,
			Shares:                sdk.NewDec(1000_000),
			RewardHistory:         nil,
			LastRewardClaimHeight: uint64(ctx.BlockHeight()),
		},
		Balance: sdk.Coin{
			Denom:  KAIJU_TOKEN_DENOM,
			Amount: sdk.NewInt(1000_000),
		},
	}, queryDelegations.Delegations[0])

	queryDelegations, queryErr = queryServer.AllKaijusDelegations(ctx, &types.QueryAllKaijusDelegationsRequest{
		Pagination: &query.PageRequest{
			Key:        queryDelegations.Pagination.NextKey,
			Offset:     0,
			Limit:      1,
			CountTotal: false,
			Reverse:    false,
		},
	})
	require.NoError(t, queryErr)
	require.Equal(t, types.DelegationResponse{
		Delegation: types.Delegation{
			DelegatorAddress:      delAddr.String(),
			ValidatorAddress:      val.OperatorAddress,
			Denom:                 KAIJU_2_TOKEN_DENOM,
			Shares:                sdk.NewDec(1000_000),
			RewardHistory:         nil,
			LastRewardClaimHeight: uint64(ctx.BlockHeight()),
		},
		Balance: sdk.Coin{
			Denom:  KAIJU_2_TOKEN_DENOM,
			Amount: sdk.NewInt(1000_000),
		},
	}, queryDelegations.Delegations[0])
}

func TestQueryValidator(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH KAIJUS ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                KAIJU_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.NewDec(0),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
			{
				Denom:                KAIJU_2_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(10),
				TakeRate:             sdk.MustNewDecFromStr("0.14159265359"),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	delAddr, _ := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	valAddr, _ := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	val, _ := app.KaijuKeeper.GetKaijuValidator(ctx, valAddr)
	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(2000_000))))

	// WHEN: DELEGATING AND QUERYING ...
	_, txErr := app.KaijuKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.NoError(t, txErr)
	_, tx2Err := app.KaijuKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.NoError(t, tx2Err)

	queryVal, queryErr := queryServer.KaijuValidator(ctx, &types.QueryKaijuValidatorRequest{
		ValidatorAddr: val.GetOperator().String(),
	})

	require.NoError(t, queryErr)
	require.Equal(t, &types.QueryKaijuValidatorResponse{
		ValidatorAddr: val.GetOperator().String(),
		TotalDelegationShares: sdk.NewDecCoins(
			sdk.NewDecCoinFromDec(KAIJU_TOKEN_DENOM, sdk.NewDec(1000000)),
			sdk.NewDecCoinFromDec(KAIJU_2_TOKEN_DENOM, sdk.NewDec(1000000)),
		),
		ValidatorShares: sdk.NewDecCoins(
			sdk.NewDecCoinFromDec(KAIJU_TOKEN_DENOM, sdk.NewDec(1000000)),
			sdk.NewDecCoinFromDec(KAIJU_2_TOKEN_DENOM, sdk.NewDec(1000000)),
		),
		TotalStaked: sdk.NewDecCoins(
			sdk.NewDecCoinFromDec(KAIJU_TOKEN_DENOM, sdk.NewDec(1000_000)),
			sdk.NewDecCoinFromDec(KAIJU_2_TOKEN_DENOM, sdk.NewDec(1000_000)),
		),
	}, queryVal)
}

func TestQueryValidators(t *testing.T) {
	// GIVEN: THE BLOCKCHAIN WITH KAIJUS ON GENESIS
	app, ctx := createTestContext(t)
	startTime := time.Now()
	ctx = ctx.WithBlockTime(startTime).WithBlockHeight(1)
	app.KaijuKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.DefaultParams(),
		Assets: []types.KaijuAsset{
			{
				Denom:                KAIJU_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(2),
				TakeRate:             sdk.NewDec(0),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
			{
				Denom:                KAIJU_2_TOKEN_DENOM,
				RewardWeight:         sdk.NewDec(10),
				TakeRate:             sdk.MustNewDecFromStr("0.14159265359"),
				TotalTokens:          sdk.ZeroInt(),
				TotalValidatorShares: sdk.NewDec(0),
				RewardChangeRate:     sdk.NewDec(0),
				RewardChangeInterval: 0,
			},
		},
	})
	queryServer := keeper.NewQueryServerImpl(app.KaijuKeeper)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	delAddr, _ := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	valAddr, _ := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	val, _ := app.KaijuKeeper.GetKaijuValidator(ctx, valAddr)

	addrs := test_helpers.AddTestAddrsIncremental(app, ctx, 3, sdk.NewCoins(
		sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(1000_000)),
		sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(1000_000)),
	))
	valAddr2 := sdk.ValAddress(addrs[0])
	_val2 := teststaking.NewValidator(t, valAddr2, test_helpers.CreateTestPubKeys(1)[0])
	test_helpers.RegisterNewValidator(t, app, ctx, _val2)
	val2, err := app.KaijuKeeper.GetKaijuValidator(ctx, valAddr2)
	require.NoError(t, err)

	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(2000_000))))
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(2000_000))))

	// WHEN: DELEGATING AND QUERYING ...
	_, txErr := app.KaijuKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(KAIJU_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.NoError(t, txErr)
	_, tx2Err := app.KaijuKeeper.Delegate(ctx, delAddr, val2, sdk.NewCoin(KAIJU_2_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.NoError(t, tx2Err)

	queryVal, queryErr := queryServer.AllKaijuValidators(ctx, &types.QueryAllKaijuValidatorsRequest{
		Pagination: &query.PageRequest{
			Key:        nil,
			Offset:     0,
			Limit:      1,
			CountTotal: false,
			Reverse:    false,
		},
	})

	require.NoError(t, queryErr)
	require.Equal(t, &types.QueryKaijuValidatorsResponse{
		Validators: []types.QueryKaijuValidatorResponse{
			{
				ValidatorAddr: val.GetOperator().String(),
				TotalDelegationShares: sdk.NewDecCoins(
					sdk.NewDecCoinFromDec(KAIJU_TOKEN_DENOM, sdk.NewDec(1000000)),
				),
				ValidatorShares: sdk.NewDecCoins(
					sdk.NewDecCoinFromDec(KAIJU_TOKEN_DENOM, sdk.NewDec(1000000)),
				),
				TotalStaked: sdk.NewDecCoins(
					sdk.NewDecCoinFromDec(KAIJU_TOKEN_DENOM, sdk.NewDec(1000_000)),
				),
			},
		},
		Pagination: queryVal.Pagination,
	}, queryVal)

	queryVal2, queryErr := queryServer.AllKaijuValidators(ctx, &types.QueryAllKaijuValidatorsRequest{
		Pagination: &query.PageRequest{
			Key:        queryVal.Pagination.NextKey,
			Offset:     0,
			Limit:      1,
			CountTotal: false,
			Reverse:    false,
		},
	})

	require.NoError(t, queryErr)
	require.Equal(t, &types.QueryKaijuValidatorsResponse{
		Validators: []types.QueryKaijuValidatorResponse{
			{
				ValidatorAddr: val2.GetOperator().String(),
				TotalDelegationShares: sdk.NewDecCoins(
					sdk.NewDecCoinFromDec(KAIJU_2_TOKEN_DENOM, sdk.NewDec(1000000)),
				),
				ValidatorShares: sdk.NewDecCoins(
					sdk.NewDecCoinFromDec(KAIJU_2_TOKEN_DENOM, sdk.NewDec(1000000)),
				),
				TotalStaked: sdk.NewDecCoins(
					sdk.NewDecCoinFromDec(KAIJU_2_TOKEN_DENOM, sdk.NewDec(1000_000)),
				),
			},
		},
		Pagination: queryVal2.Pagination,
	}, queryVal2)
}
