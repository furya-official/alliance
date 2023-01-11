package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cobra"
	"github.com/furya-official/furya/x/furya/types"
	"time"
)

func CreateFurya() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-furya denom rewards-weight take-rate reward-change-rate reward-change-interval",
		Args:  cobra.ExactArgs(5),
		Short: "Create an furya with the specified parameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			title, err := cmd.Flags().GetString(govcli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(govcli.FlagDescription)
			if err != nil {
				return err
			}

			rewardWeight, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}

			takeRate, err := sdk.NewDecFromStr(args[2])
			if err != nil {
				return err
			}

			rewardChangeRate, err := sdk.NewDecFromStr(args[3])
			if err != nil {
				return err
			}

			rewardChangeInterval, err := time.ParseDuration(args[4])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewMsgCreateFuryaProposal(
				title,
				description,
				args[0],
				rewardWeight,
				takeRate,
				rewardChangeRate,
				rewardChangeInterval,
			)

			err = content.ValidateBasic()

			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)

			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")
	return cmd
}

func UpdateFurya() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-furya denom rewards-weight take-rate reward-change-rate reward-change-interval",
		Args:  cobra.ExactArgs(5),
		Short: "Update an furya with the specified parameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			title, err := cmd.Flags().GetString(govcli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(govcli.FlagDescription)
			if err != nil {
				return err
			}

			rewardWeight, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}

			takeRate, err := sdk.NewDecFromStr(args[2])
			if err != nil {
				return err
			}

			rewardChangeRate, err := sdk.NewDecFromStr(args[3])
			if err != nil {
				return err
			}

			rewardChangeInterval, err := time.ParseDuration(args[4])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewMsgCreateFuryaProposal(
				title,
				description,
				args[0],
				rewardWeight,
				takeRate,
				rewardChangeRate,
				rewardChangeInterval,
			)

			err = content.ValidateBasic()

			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)

			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")
	return cmd
}

func DeleteFurya() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-furya denom",
		Args:  cobra.ExactArgs(1),
		Short: "Delete an furya with the specified denom",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			title, err := cmd.Flags().GetString(govcli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(govcli.FlagDescription)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewMsgDeleteFuryaProposal(
				title,
				description,
				args[0],
			)

			err = content.ValidateBasic()

			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)

			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")
	return cmd
}
