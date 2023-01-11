package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/furya/x/furya/types"
)

func (k Keeper) GetFuryaValidator(ctx sdk.Context, valAddr sdk.ValAddress) (types.FuryaValidator, error) {
	val, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return types.FuryaValidator{}, fmt.Errorf("validator with address %s does not exist", valAddr.String())
	}
	valInfo, found := k.GetFuryaValidatorInfo(ctx, valAddr)
	if !found {
		valInfo = k.createFuryaValidatorInfo(ctx, valAddr)
	}
	return types.FuryaValidator{
		Validator:             &val,
		FuryaValidatorInfo: &valInfo,
	}, nil
}

func (k Keeper) GetFuryaValidatorInfo(ctx sdk.Context, valAddr sdk.ValAddress) (types.FuryaValidatorInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFuryaValidatorInfoKey(valAddr)
	vb := store.Get(key)
	var info types.FuryaValidatorInfo
	if vb == nil {
		return info, false
	} else {
		k.cdc.MustUnmarshal(vb, &info)
		return info, true
	}
}

func (k Keeper) createFuryaValidatorInfo(ctx sdk.Context, valAddr sdk.ValAddress) (val types.FuryaValidatorInfo) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFuryaValidatorInfoKey(valAddr)
	val = types.NewFuryaValidatorInfo()
	vb := k.cdc.MustMarshal(&val)
	store.Set(key, vb)
	return val
}

func (k Keeper) IterateFuryaValidatorInfo(ctx sdk.Context, cb func(valAddr sdk.ValAddress, info types.FuryaValidatorInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorInfoKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.FuryaValidatorInfo
		b := iter.Value()
		k.cdc.MustUnmarshal(b, &info)
		valAddr := types.ParseFuryaValidatorKey(iter.Key())
		if cb(valAddr, info) {
			return
		}
	}
}

func (k Keeper) GetAllFuryaValidatorInfo(ctx sdk.Context) []types.FuryaValidatorInfo {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorInfoKey)
	defer iter.Close()
	var infos []types.FuryaValidatorInfo
	for ; iter.Valid(); iter.Next() {
		b := iter.Value()
		var info types.FuryaValidatorInfo
		k.cdc.UnmarshalInterface(b, &info)
		infos = append(infos, info)
	}
	return infos
}
