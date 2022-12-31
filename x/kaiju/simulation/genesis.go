package simulation

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/tendermint/tendermint/libs/json"
	"github.com/furya-official/kaiju/x/kaiju/types"
	"math/rand"
	"time"
)

func genRewardDelayTime(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 60, 60*60*24*3*2)) * time.Second
}

func genTakeRateClaimInterval(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 1, 60*60)) * time.Second
}

func genNumOfKaijuAssets(r *rand.Rand) int {
	return simulation.RandIntBetween(r, 0, 50)
}

func RandomizedGenesisState(simState *module.SimulationState) {
	var (
		rewardDelayTime     time.Duration
		rewardClaimInterval time.Duration
		numOfKaijuAssets int
	)

	r := simState.Rand
	rewardDelayTime = genRewardDelayTime(r)
	rewardClaimInterval = genTakeRateClaimInterval(r)
	numOfKaijuAssets = genNumOfKaijuAssets(r)

	var kaijuAssets []types.KaijuAsset
	for i := 0; i < numOfKaijuAssets; i += 1 {
		rewardRate := simulation.RandomDecAmount(r, sdk.NewDec(5))
		takeRate := simulation.RandomDecAmount(r, sdk.MustNewDecFromStr("0.5"))
		startTime := time.Now().Add(time.Duration(simulation.RandIntBetween(r, 60, 60*60*24*3*2)) * time.Second)
		kaijuAssets = append(kaijuAssets, types.NewKaijuAsset(fmt.Sprintf("ASSET%d", i), rewardRate, takeRate, startTime))
	}

	kaijuGenesis := types.GenesisState{
		Params: types.Params{
			RewardDelayTime:       rewardDelayTime,
			TakeRateClaimInterval: rewardClaimInterval,
			LastTakeRateClaimTime: simState.GenTimestamp,
		},
		Assets: kaijuAssets,
	}

	bz, err := json.MarshalIndent(&kaijuGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated kaiju parameters:\n%s\n", bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&kaijuGenesis)
}
