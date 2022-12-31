package types

func NewRewardWeightChangeSnapshot(asset KaijuAsset, val KaijuValidator) RewardWeightChangeSnapshot {
	return RewardWeightChangeSnapshot{
		PrevRewardWeight: asset.RewardWeight,
		RewardHistories:  val.GlobalRewardHistory,
	}
}
