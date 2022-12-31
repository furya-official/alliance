package client

import (
	"github.com/furya-official/kaiju/x/kaiju/client/cli"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	CreateKaijuProposalHandler = govclient.NewProposalHandler(cli.CreateKaiju)
	UpdateKaijuProposalHandler = govclient.NewProposalHandler(cli.UpdateKaiju)
	DeleteKaijuProposalHandler = govclient.NewProposalHandler(cli.DeleteKaiju)
)
