package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const (
	ProposalTypeCreateKaiju = "msg_create_kaiju_proposal"
	ProposalTypeUpdateKaiju = "msg_update_kaiju_proposal"
	ProposalTypeDeleteKaiju = "msg_delete_kaiju_proposal"
)

var (
	_ govtypes.Content = &MsgCreateKaijuProposal{}
	_ govtypes.Content = &MsgUpdateKaijuProposal{}
	_ govtypes.Content = &MsgDeleteKaijuProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateKaiju)
	govtypes.RegisterProposalType(ProposalTypeUpdateKaiju)
	govtypes.RegisterProposalType(ProposalTypeDeleteKaiju)
}
func NewMsgCreateKaijuProposal(title, description, denom string, rewardWeight, takeRate sdk.Dec, rewardChangeRate sdk.Dec, rewardChangeInterval time.Duration) govtypes.Content {
	return &MsgCreateKaijuProposal{
		Title:                title,
		Description:          description,
		Denom:                denom,
		RewardWeight:         rewardWeight,
		TakeRate:             takeRate,
		RewardChangeRate:     rewardChangeRate,
		RewardChangeInterval: rewardChangeInterval,
	}
}
func (m *MsgCreateKaijuProposal) GetTitle() string       { return m.Title }
func (m *MsgCreateKaijuProposal) GetDescription() string { return m.Description }
func (m *MsgCreateKaijuProposal) ProposalRoute() string  { return RouterKey }
func (m *MsgCreateKaijuProposal) ProposalType() string   { return ProposalTypeCreateKaiju }

func (m *MsgCreateKaijuProposal) ValidateBasic() error {

	if m.Denom == "" {
		return status.Errorf(codes.InvalidArgument, "Kaiju denom must have a value")
	}

	if m.RewardWeight.IsNil() || m.RewardWeight.LTE(sdk.ZeroDec()) {
		return status.Errorf(codes.InvalidArgument, "Kaiju rewardWeight must be a positive number")
	}

	if m.TakeRate.IsNil() || m.TakeRate.IsNegative() || m.TakeRate.GTE(sdk.OneDec()) {
		return status.Errorf(codes.InvalidArgument, "Kaiju takeRate must be more or equals to 0 but strictly less than 1")
	}

	if m.RewardChangeRate.IsZero() || m.RewardChangeRate.IsNegative() {
		return status.Errorf(codes.InvalidArgument, "Kaiju rewardChangeRate must be strictly a positive number")
	}

	return nil
}

func NewMsgUpdateKaijuProposal(title, description, denom string, rewardWeight, takeRate sdk.Dec, rewardChangeRate sdk.Dec, rewardChangeInterval time.Duration) govtypes.Content {
	return &MsgUpdateKaijuProposal{
		Title:                title,
		Description:          description,
		Denom:                denom,
		RewardWeight:         rewardWeight,
		TakeRate:             takeRate,
		RewardChangeRate:     rewardChangeRate,
		RewardChangeInterval: rewardChangeInterval,
	}
}
func (m *MsgUpdateKaijuProposal) GetTitle() string       { return m.Title }
func (m *MsgUpdateKaijuProposal) GetDescription() string { return m.Description }
func (m *MsgUpdateKaijuProposal) ProposalRoute() string  { return RouterKey }
func (m *MsgUpdateKaijuProposal) ProposalType() string   { return ProposalTypeUpdateKaiju }

func (m *MsgUpdateKaijuProposal) ValidateBasic() error {
	if m.Denom == "" {
		return status.Errorf(codes.InvalidArgument, "Kaiju denom must have a value")
	}

	if m.RewardWeight.IsNil() || m.RewardWeight.LTE(sdk.ZeroDec()) {
		return status.Errorf(codes.InvalidArgument, "Kaiju rewardWeight must be a positive number")
	}

	if m.TakeRate.IsNil() || m.TakeRate.IsNegative() || m.TakeRate.GTE(sdk.OneDec()) {
		return status.Errorf(codes.InvalidArgument, "Kaiju takeRate must be more or equals to 0 but strictly less than 1")
	}

	if m.RewardChangeRate.IsZero() || m.RewardChangeRate.IsNegative() {
		return status.Errorf(codes.InvalidArgument, "Kaiju rewardChangeRate must be strictly a positive number")
	}

	return nil
}

func NewMsgDeleteKaijuProposal(title, description, denom string) govtypes.Content {
	return &MsgDeleteKaijuProposal{
		Title:       title,
		Description: description,
		Denom:       denom,
	}
}
func (m *MsgDeleteKaijuProposal) GetTitle() string       { return m.Title }
func (m *MsgDeleteKaijuProposal) GetDescription() string { return m.Description }
func (m *MsgDeleteKaijuProposal) ProposalRoute() string  { return RouterKey }
func (m *MsgDeleteKaijuProposal) ProposalType() string   { return ProposalTypeDeleteKaiju }

func (m *MsgDeleteKaijuProposal) ValidateBasic() error {
	if m.Denom == "" {
		return status.Errorf(codes.InvalidArgument, "Kaiju denom must have a value")
	}
	return nil
}
