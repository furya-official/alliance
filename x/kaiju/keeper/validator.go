package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/kaiju/x/kaiju/types"
)

func (k Keeper) GetKaijuValidator(ctx sdk.Context, valAddr sdk.ValAddress) (types.KaijuValidator, error) {
	val, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return types.KaijuValidator{}, fmt.Errorf("validator with address %s does not exist", valAddr.String())
	}
	valInfo, found := k.GetKaijuValidatorInfo(ctx, valAddr)
	if !found {
		valInfo = k.createKaijuValidatorInfo(ctx, valAddr)
	}
	return types.KaijuValidator{
		Validator:             &val,
		KaijuValidatorInfo: &valInfo,
	}, nil
}

func (k Keeper) GetKaijuValidatorInfo(ctx sdk.Context, valAddr sdk.ValAddress) (types.KaijuValidatorInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKaijuValidatorInfoKey(valAddr)
	vb := store.Get(key)
	var info types.KaijuValidatorInfo
	if vb == nil {
		return info, false
	} else {
		k.cdc.MustUnmarshal(vb, &info)
		return info, true
	}
}

func (k Keeper) createKaijuValidatorInfo(ctx sdk.Context, valAddr sdk.ValAddress) (val types.KaijuValidatorInfo) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKaijuValidatorInfoKey(valAddr)
	val = types.NewKaijuValidatorInfo()
	vb := k.cdc.MustMarshal(&val)
	store.Set(key, vb)
	return val
}

func (k Keeper) IterateKaijuValidatorInfo(ctx sdk.Context, cb func(valAddr sdk.ValAddress, info types.KaijuValidatorInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorInfoKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.KaijuValidatorInfo
		b := iter.Value()
		k.cdc.MustUnmarshal(b, &info)
		valAddr := types.ParseKaijuValidatorKey(iter.Key())
		if cb(valAddr, info) {
			return
		}
	}
}

func (k Keeper) GetAllKaijuValidatorInfo(ctx sdk.Context) []types.KaijuValidatorInfo {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorInfoKey)
	defer iter.Close()
	var infos []types.KaijuValidatorInfo
	for ; iter.Valid(); iter.Next() {
		b := iter.Value()
		var info types.KaijuValidatorInfo
		k.cdc.UnmarshalInterface(b, &info)
		infos = append(infos, info)
	}
	return infos
}
