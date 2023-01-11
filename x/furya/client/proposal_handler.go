package client

import (
	"github.com/furya-official/furya/x/furya/client/cli"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	CreateFuryaProposalHandler = govclient.NewProposalHandler(cli.CreateFurya)
	UpdateFuryaProposalHandler = govclient.NewProposalHandler(cli.UpdateFurya)
	DeleteFuryaProposalHandler = govclient.NewProposalHandler(cli.DeleteFurya)
)
