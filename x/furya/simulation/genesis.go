package simulation

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/tendermint/tendermint/libs/json"
	"github.com/furya-official/furya/x/furya/types"
	"math/rand"
	"time"
)

func genRewardDelayTime(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 60, 60*60*24*3*2)) * time.Second
}

func genTakeRateClaimInterval(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 1, 60*60)) * time.Second
}

func genNumOfFuryaAssets(r *rand.Rand) int {
	return simulation.RandIntBetween(r, 0, 50)
}

func RandomizedGenesisState(simState *module.SimulationState) {
	var (
		rewardDelayTime     time.Duration
		rewardClaimInterval time.Duration
		numOfFuryaAssets int
	)

	r := simState.Rand
	rewardDelayTime = genRewardDelayTime(r)
	rewardClaimInterval = genTakeRateClaimInterval(r)
	numOfFuryaAssets = genNumOfFuryaAssets(r)

	var furyaAssets []types.FuryaAsset
	for i := 0; i < numOfFuryaAssets; i += 1 {
		rewardRate := simulation.RandomDecAmount(r, sdk.NewDec(5))
		takeRate := simulation.RandomDecAmount(r, sdk.MustNewDecFromStr("0.5"))
		startTime := time.Now().Add(time.Duration(simulation.RandIntBetween(r, 60, 60*60*24*3*2)) * time.Second)
		furyaAssets = append(furyaAssets, types.NewFuryaAsset(fmt.Sprintf("ASSET%d", i), rewardRate, takeRate, startTime))
	}

	furyaGenesis := types.GenesisState{
		Params: types.Params{
			RewardDelayTime:       rewardDelayTime,
			TakeRateClaimInterval: rewardClaimInterval,
			LastTakeRateClaimTime: simState.GenTimestamp,
		},
		Assets: furyaAssets,
	}

	bz, err := json.MarshalIndent(&furyaGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated furya parameters:\n%s\n", bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&furyaGenesis)
}
