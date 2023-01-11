package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const (
	ProposalTypeCreateFurya = "msg_create_furya_proposal"
	ProposalTypeUpdateFurya = "msg_update_furya_proposal"
	ProposalTypeDeleteFurya = "msg_delete_furya_proposal"
)

var (
	_ govtypes.Content = &MsgCreateFuryaProposal{}
	_ govtypes.Content = &MsgUpdateFuryaProposal{}
	_ govtypes.Content = &MsgDeleteFuryaProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateFurya)
	govtypes.RegisterProposalType(ProposalTypeUpdateFurya)
	govtypes.RegisterProposalType(ProposalTypeDeleteFurya)
}
func NewMsgCreateFuryaProposal(title, description, denom string, rewardWeight, takeRate sdk.Dec, rewardChangeRate sdk.Dec, rewardChangeInterval time.Duration) govtypes.Content {
	return &MsgCreateFuryaProposal{
		Title:                title,
		Description:          description,
		Denom:                denom,
		RewardWeight:         rewardWeight,
		TakeRate:             takeRate,
		RewardChangeRate:     rewardChangeRate,
		RewardChangeInterval: rewardChangeInterval,
	}
}
func (m *MsgCreateFuryaProposal) GetTitle() string       { return m.Title }
func (m *MsgCreateFuryaProposal) GetDescription() string { return m.Description }
func (m *MsgCreateFuryaProposal) ProposalRoute() string  { return RouterKey }
func (m *MsgCreateFuryaProposal) ProposalType() string   { return ProposalTypeCreateFurya }

func (m *MsgCreateFuryaProposal) ValidateBasic() error {

	if m.Denom == "" {
		return status.Errorf(codes.InvalidArgument, "Furya denom must have a value")
	}

	if m.RewardWeight.IsNil() || m.RewardWeight.LTE(sdk.ZeroDec()) {
		return status.Errorf(codes.InvalidArgument, "Furya rewardWeight must be a positive number")
	}

	if m.TakeRate.IsNil() || m.TakeRate.IsNegative() || m.TakeRate.GTE(sdk.OneDec()) {
		return status.Errorf(codes.InvalidArgument, "Furya takeRate must be more or equals to 0 but strictly less than 1")
	}

	if m.RewardChangeRate.IsZero() || m.RewardChangeRate.IsNegative() {
		return status.Errorf(codes.InvalidArgument, "Furya rewardChangeRate must be strictly a positive number")
	}

	return nil
}

func NewMsgUpdateFuryaProposal(title, description, denom string, rewardWeight, takeRate sdk.Dec, rewardChangeRate sdk.Dec, rewardChangeInterval time.Duration) govtypes.Content {
	return &MsgUpdateFuryaProposal{
		Title:                title,
		Description:          description,
		Denom:                denom,
		RewardWeight:         rewardWeight,
		TakeRate:             takeRate,
		RewardChangeRate:     rewardChangeRate,
		RewardChangeInterval: rewardChangeInterval,
	}
}
func (m *MsgUpdateFuryaProposal) GetTitle() string       { return m.Title }
func (m *MsgUpdateFuryaProposal) GetDescription() string { return m.Description }
func (m *MsgUpdateFuryaProposal) ProposalRoute() string  { return RouterKey }
func (m *MsgUpdateFuryaProposal) ProposalType() string   { return ProposalTypeUpdateFurya }

func (m *MsgUpdateFuryaProposal) ValidateBasic() error {
	if m.Denom == "" {
		return status.Errorf(codes.InvalidArgument, "Furya denom must have a value")
	}

	if m.RewardWeight.IsNil() || m.RewardWeight.LTE(sdk.ZeroDec()) {
		return status.Errorf(codes.InvalidArgument, "Furya rewardWeight must be a positive number")
	}

	if m.TakeRate.IsNil() || m.TakeRate.IsNegative() || m.TakeRate.GTE(sdk.OneDec()) {
		return status.Errorf(codes.InvalidArgument, "Furya takeRate must be more or equals to 0 but strictly less than 1")
	}

	if m.RewardChangeRate.IsZero() || m.RewardChangeRate.IsNegative() {
		return status.Errorf(codes.InvalidArgument, "Furya rewardChangeRate must be strictly a positive number")
	}

	return nil
}

func NewMsgDeleteFuryaProposal(title, description, denom string) govtypes.Content {
	return &MsgDeleteFuryaProposal{
		Title:       title,
		Description: description,
		Denom:       denom,
	}
}
func (m *MsgDeleteFuryaProposal) GetTitle() string       { return m.Title }
func (m *MsgDeleteFuryaProposal) GetDescription() string { return m.Description }
func (m *MsgDeleteFuryaProposal) ProposalRoute() string  { return RouterKey }
func (m *MsgDeleteFuryaProposal) ProposalType() string   { return ProposalTypeDeleteFurya }

func (m *MsgDeleteFuryaProposal) ValidateBasic() error {
	if m.Denom == "" {
		return status.Errorf(codes.InvalidArgument, "Furya denom must have a value")
	}
	return nil
}
