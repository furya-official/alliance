package cli

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/furya/x/furya/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())

	cmd.AddCommand(CmdQueryFuryas())
	cmd.AddCommand(CmdQueryFurya())

	cmd.AddCommand(CmdQueryValidator())
	cmd.AddCommand(CmdQueryValidators())

	cmd.AddCommand(CmdQueryAllFuryasDelegations())
	cmd.AddCommand(CmdQueryFuryasDelegation())
	cmd.AddCommand(CmdQueryFuryasDelegationByValidator())
	cmd.AddCommand(CmdQueryFuryaDelegation())
	cmd.AddCommand(CmdQueryRewards())

	return cmd
}

func CmdQueryFuryas() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "furyas",
		Short: "Query paginated furyas",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryFuryasRequest{}

			res, err := queryClient.Furyas(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryFurya() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "furya denom",
		Short: "Query a specific furya by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			denom := args[0]

			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			query := types.NewQueryClient(ctx)

			params := &types.QueryFuryaRequest{Denom: denom}

			res, err := query.Furya(cmd.Context(), params)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator validator-addr",
		Short: "Query a specific furya validator by addr",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			query := types.NewQueryClient(ctx)

			req := &types.QueryFuryaValidatorRequest{ValidatorAddr: valAddr.String()}

			res, err := query.FuryaValidator(cmd.Context(), req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryValidators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validators",
		Short: "Query all furya validators",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			query := types.NewQueryClient(ctx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryAllFuryaValidatorsRequest{
				Pagination: pageReq,
			}

			res, err := query.AllFuryaValidators(cmd.Context(), req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryAllFuryasDelegations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegations",
		Short: "Query all paginated furyas delegations",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := client.GetClientContextFromCmd(cmd)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			query := types.NewQueryClient(ctx)

			params := &types.QueryAllFuryasDelegationsRequest{
				Pagination: pageReq,
			}

			res, err := query.AllFuryasDelegations(context.Background(), params)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryFuryasDelegation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegations-by-delegator delegator_addr",
		Short: "Query all paginated furyas delegations for a delegator_addr",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			delegator_addr := args[0]
			ctx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			query := types.NewQueryClient(ctx)

			params := &types.QueryFuryasDelegationsRequest{
				DelegatorAddr: delegator_addr,
				Pagination:    pageReq,
			}

			res, err := query.FuryasDelegation(context.Background(), params)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryFuryasDelegationByValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegations-by-delegator-and-validator delegator_addr validator_addr",
		Short: "Query all paginated furya delegations for a delegator_addr and validator_addr",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			delegator_addr := args[0]
			validator_addr := args[1]
			ctx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			query := types.NewQueryClient(ctx)

			params := &types.QueryFuryasDelegationByValidatorRequest{
				Pagination:    pageReq,
				DelegatorAddr: delegator_addr,
				ValidatorAddr: validator_addr,
			}

			res, err := query.FuryasDelegationByValidator(context.Background(), params)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryFuryaDelegation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegation delegator_addr validator_addr denom",
		Short: "Query a delegation to an furya by delegator_addr, validator_addr and denom",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			delegator_addr := args[0]
			validator_addr := args[1]
			denom := args[2]
			ctx := client.GetClientContextFromCmd(cmd)

			if err != nil {
				return err
			}
			query := types.NewQueryClient(ctx)

			params := &types.QueryFuryaDelegationRequest{
				DelegatorAddr: delegator_addr,
				ValidatorAddr: validator_addr,
				Denom:         denom,
			}

			res, err := query.FuryaDelegation(context.Background(), params)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query module parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rewards delegator_addr validator_addr denom",
		Short: "Query module parameters",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			delegatorAddr := args[0]
			validatorAddr := args[1]
			denom := args[2]
			ctx := client.GetClientContextFromCmd(cmd)
			query := types.NewQueryClient(ctx)
			params := &types.QueryFuryaDelegationRewardsRequest{
				DelegatorAddr: delegatorAddr,
				ValidatorAddr: validatorAddr,
				Denom:         denom,
			}

			res, err := query.FuryaDelegationRewards(context.Background(), params)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
