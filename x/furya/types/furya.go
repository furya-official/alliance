package types

func NewRewardWeightChangeSnapshot(asset FuryaAsset, val FuryaValidator) RewardWeightChangeSnapshot {
	return RewardWeightChangeSnapshot{
		PrevRewardWeight: asset.RewardWeight,
		RewardHistories:  val.GlobalRewardHistory,
	}
}
