package keeper

import (
	"context"
	"fmt"
	"github.com/furya-official/kaiju/x/kaiju/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QueryServer struct {
	Keeper
}

func (k QueryServer) AllKaijusDelegations(c context.Context, req *types.QueryAllKaijusDelegationsRequest) (*types.QueryKaijusDelegationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	res := &types.QueryKaijusDelegationsResponse{
		Delegations: nil,
		Pagination:  nil,
	}

	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	delegationStore := prefix.NewStore(store, types.DelegationKey)

	pageRes, err := query.Paginate(delegationStore, req.Pagination, func(key []byte, value []byte) error {
		var delegation types.Delegation
		k.cdc.MustUnmarshal(value, &delegation)

		asset, found := k.GetAssetByDenom(ctx, delegation.Denom)
		if !found {
			return types.ErrUnknownAsset
		}

		valAddr, _ := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		validator, err := k.GetKaijuValidator(ctx, valAddr)
		if err != nil {
			return err
		}
		balance := types.GetDelegationTokens(delegation, validator, asset)

		delegationRes := types.DelegationResponse{
			Delegation: delegation,
			Balance:    balance,
		}
		res.Delegations = append(res.Delegations, delegationRes)
		return nil
	})
	if err != nil {
		return nil, err
	}
	res.Pagination = pageRes
	return res, nil
}

func (k QueryServer) KaijuValidator(c context.Context, req *types.QueryKaijuValidatorRequest) (*types.QueryKaijuValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	res := types.QueryKaijuValidatorResponse{
		ValidatorAddr:         req.ValidatorAddr,
		TotalDelegationShares: nil,
		ValidatorShares:       nil,
		TotalStaked:           nil,
	}
	valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("validator address %s invalid", req.ValidatorAddr))
	}
	val, err := k.GetKaijuValidator(ctx, valAddr)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("validator with address %s not found", req.ValidatorAddr))
	}
	res.ValidatorShares = val.ValidatorShares
	res.TotalDelegationShares = val.TotalDelegatorShares

	for _, share := range val.ValidatorShares {
		asset, found := k.GetAssetByDenom(ctx, share.Denom)
		if !found {
			continue
		}
		res.TotalStaked = append(res.TotalStaked, sdk.NewDecCoinFromDec(share.Denom, val.TotalTokensWithAsset(asset)))
	}
	return &res, nil
}

func (k QueryServer) AllKaijuValidators(c context.Context, req *types.QueryAllKaijuValidatorsRequest) (*types.QueryKaijuValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	res := &types.QueryKaijuValidatorsResponse{
		Validators: nil,
		Pagination: nil,
	}

	store := ctx.KVStore(k.storeKey)
	valStore := prefix.NewStore(store, types.ValidatorInfoKey)

	pageRes, err := query.Paginate(valStore, req.Pagination, func(key []byte, value []byte) error {
		valAddr := sdk.ValAddress(key[1:]) // Due to length prefix when encoding the key
		val, err := k.GetKaijuValidator(ctx, valAddr)
		if err != nil {
			return err
		}

		totalStaked := sdk.NewDecCoins()
		for _, share := range val.ValidatorShares {
			asset, found := k.GetAssetByDenom(ctx, share.Denom)
			if !found {
				continue
			}
			totalStaked = append(totalStaked, sdk.NewDecCoinFromDec(share.Denom, val.TotalTokensWithAsset(asset)))
		}

		res.Validators = append(res.Validators, types.QueryKaijuValidatorResponse{
			ValidatorAddr:         valAddr.String(),
			TotalDelegationShares: val.TotalDelegatorShares,
			ValidatorShares:       val.ValidatorShares,
			TotalStaked:           totalStaked,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	res.Pagination = pageRes
	return res, nil
}

var _ types.QueryServer = QueryServer{}

func (k QueryServer) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	// Define a variable that will store the params
	var params types.Params

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(c)

	k.paramstore.GetParamSet(ctx, &params)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}

func (k QueryServer) Kaijus(c context.Context, req *types.QueryKaijusRequest) (*types.QueryKaijusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of assets
	var kaijus []types.KaijuAsset

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(c)

	// Get the key-value module store using the store key
	store := ctx.KVStore(k.storeKey)

	// Get the part of the store that keeps assets
	assetsStore := prefix.NewStore(store, types.AssetKey)

	// Paginate the assets store based on PageRequest
	pageRes, err := query.Paginate(assetsStore, req.Pagination, func(key []byte, value []byte) error {
		var asset types.KaijuAsset
		if err := k.cdc.Unmarshal(value, &asset); err != nil {
			return err
		}

		kaijus = append(kaijus, asset)

		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return a struct containing a list of assets and pagination info
	return &types.QueryKaijusResponse{
		Kaijus:  kaijus,
		Pagination: pageRes,
	}, nil
}

func (k QueryServer) Kaiju(c context.Context, req *types.QueryKaijuRequest) (*types.QueryKaijuResponse, error) {
	var asset types.KaijuAsset

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(c)

	// Get the part of the store that keeps assets
	asset, found := k.GetAssetByDenom(ctx, req.Denom)

	if !found {
		return nil, types.ErrUnknownAsset
	}

	// Return parsed asset, true since the asset exists
	return &types.QueryKaijuResponse{
		Kaiju: &asset,
	}, nil
}

func (k QueryServer) IBCKaiju(c context.Context, request *types.QueryIBCKaijuRequest) (*types.QueryKaijuResponse, error) {
	req := types.QueryKaijuRequest{
		Denom: "ibc/" + request.Hash,
	}
	return k.Kaiju(c, &req)
}

func (k QueryServer) KaijuDelegationRewards(context context.Context, request *types.QueryKaijuDelegationRewardsRequest) (*types.QueryKaijuDelegationRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	delAddr, err := sdk.AccAddressFromBech32(request.DelegatorAddr)
	if err != nil {
		return nil, err
	}
	valAddr, err := sdk.ValAddressFromBech32(request.ValidatorAddr)
	if err != nil {
		return nil, err
	}
	_, found := k.GetAssetByDenom(ctx, request.Denom)
	if !found {
		return nil, types.ErrUnknownAsset
	}

	val, err := k.GetKaijuValidator(ctx, valAddr)
	if err != nil {
		return nil, err
	}

	_, found = k.GetDelegation(ctx, delAddr, val, request.Denom)
	if !found {
		return nil, stakingtypes.ErrNoDelegation
	}

	rewards, err := k.ClaimDelegationRewards(ctx, delAddr, val, request.Denom)
	if err != nil {
		return nil, err
	}
	return &types.QueryKaijuDelegationRewardsResponse{
		Rewards: rewards,
	}, nil
}

func (k QueryServer) IBCKaijuDelegationRewards(context context.Context, request *types.QueryIBCKaijuDelegationRewardsRequest) (*types.QueryKaijuDelegationRewardsResponse, error) {
	req := types.QueryKaijuDelegationRewardsRequest{
		DelegatorAddr: request.DelegatorAddr,
		ValidatorAddr: request.ValidatorAddr,
		Denom:         "ibc/" + request.Hash,
		Pagination:    request.Pagination,
	}

	return k.KaijuDelegationRewards(context, &req)
}

func (k QueryServer) KaijusDelegation(c context.Context, req *types.QueryKaijusDelegationsRequest) (*types.QueryKaijusDelegationsResponse, error) {
	var delegationsRes []types.DelegationResponse

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(c)

	delAddr, err := sdk.AccAddressFromBech32(req.DelegatorAddr)
	if err != nil {
		return nil, err
	}

	// Get the key-value module store using the store key
	store := ctx.KVStore(k.storeKey)

	// Get the specific delegations key
	key := types.GetDelegationsKey(delAddr)

	// Get the part of the store that keeps assets
	delegationsStore := prefix.NewStore(store, key)

	// Paginate the assets store based on PageRequest
	pageRes, err := query.Paginate(delegationsStore, req.Pagination, func(key []byte, value []byte) error {
		var delegation types.Delegation
		if err := k.cdc.Unmarshal(value, &delegation); err != nil {
			return err
		}

		asset, found := k.GetAssetByDenom(ctx, delegation.Denom)
		if !found {
			return types.ErrUnknownAsset
		}

		valAddr, _ := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		validator, err := k.GetKaijuValidator(ctx, valAddr)
		if err != nil {
			return err
		}
		balance := types.GetDelegationTokens(delegation, validator, asset)

		delegationRes := types.DelegationResponse{
			Delegation: delegation,
			Balance:    balance,
		}

		delegationsRes = append(delegationsRes, delegationRes)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &types.QueryKaijusDelegationsResponse{
		Delegations: delegationsRes,
		Pagination:  pageRes,
	}, nil
}

func (k QueryServer) KaijusDelegationByValidator(c context.Context, req *types.QueryKaijusDelegationByValidatorRequest) (*types.QueryKaijusDelegationsResponse, error) {
	var delegationsRes []types.DelegationResponse
	ctx := sdk.UnwrapSDKContext(c)

	delAddr, err := sdk.AccAddressFromBech32(req.DelegatorAddr)
	if err != nil {
		return nil, err
	}

	valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddr)
	if err != nil {
		return nil, err
	}

	_, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "Validator not found by address %s", req.ValidatorAddr)
	}

	// Get the key-value module store using the store key
	store := ctx.KVStore(k.storeKey)

	// Get the specific delegations key
	key := types.GetDelegationsKeyForAllDenoms(delAddr, valAddr)

	// Get the part of the store that keeps assets
	delegationStore := prefix.NewStore(store, key)

	// Paginate the assets store based on PageRequest
	pageRes, err := query.Paginate(delegationStore, req.Pagination, func(key []byte, value []byte) error {
		var delegation types.Delegation
		if err := k.cdc.Unmarshal(value, &delegation); err != nil {
			return err
		}

		asset, found := k.GetAssetByDenom(ctx, delegation.Denom)
		if !found {
			return types.ErrUnknownAsset
		}

		valAddr, _ := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		validator, err := k.GetKaijuValidator(ctx, valAddr)
		if err != nil {
			return err
		}
		balance := types.GetDelegationTokens(delegation, validator, asset)

		delegationRes := types.DelegationResponse{
			Delegation: delegation,
			Balance:    balance,
		}

		delegationsRes = append(delegationsRes, delegationRes)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &types.QueryKaijusDelegationsResponse{
		Delegations: delegationsRes,
		Pagination:  pageRes,
	}, nil
}

func (k QueryServer) KaijuDelegation(c context.Context, req *types.QueryKaijuDelegationRequest) (*types.QueryKaijuDelegationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	delAddr, err := sdk.AccAddressFromBech32(req.DelegatorAddr)
	if err != nil {
		return nil, err
	}

	valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddr)
	if err != nil {
		return nil, err
	}

	validator, err := k.GetKaijuValidator(ctx, valAddr)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Validator not found by address %s", req.ValidatorAddr)
	}

	asset, found := k.GetAssetByDenom(ctx, req.Denom)

	if !found {
		return nil, status.Errorf(codes.NotFound, "KaijuAsset not found by denom %s", req.Denom)
	}

	delegation, found := k.GetDelegation(ctx, delAddr, validator, req.Denom)
	if !found {
		return &types.QueryKaijuDelegationResponse{
			Delegation: types.DelegationResponse{
				Delegation: types.NewDelegation(ctx, delAddr, valAddr, req.Denom, sdk.ZeroDec(), []types.RewardHistory{}),
				Balance:    sdk.NewCoin(req.Denom, sdk.ZeroInt()),
			}}, nil
	}

	balance := types.GetDelegationTokens(delegation, validator, asset)
	return &types.QueryKaijuDelegationResponse{
		Delegation: types.DelegationResponse{
			Delegation: delegation,
			Balance:    balance,
		},
	}, nil
}

func (k QueryServer) IBCKaijuDelegation(c context.Context, request *types.QueryIBCKaijuDelegationRequest) (*types.QueryKaijuDelegationResponse, error) {
	req := types.QueryKaijuDelegationRequest{
		DelegatorAddr: request.DelegatorAddr,
		ValidatorAddr: request.ValidatorAddr,
		Denom:         "ibc/" + request.Hash,
		Pagination:    request.Pagination,
	}
	return k.KaijuDelegation(c, &req)
}
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &QueryServer{Keeper: keeper}
}
