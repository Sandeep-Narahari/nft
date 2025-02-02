package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Sandeep-Narahari/nft/types"
)

// HasDenomID returns whether the specified denomID exists
func (k Keeper) HasDenomID(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyDenomID(id))
}

// HasDenomNm returns whether the specified denomNm exists
func (k Keeper) HasDenomNm(ctx sdk.Context, name string) bool {
	if len(name) == 0 {
		return false
	}
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyDenomName(name))
}

// SetDenom is responsible for saving the definition of denomID
func (k Keeper) SetDenom(ctx sdk.Context, denom types.Denom) error {
	if k.HasDenomID(ctx, denom.Id) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s has already exists", denom.Id)
	}

	if k.HasDenomNm(ctx, denom.Name) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomName %s has already exists", denom.Name)
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&denom)
	store.Set(types.KeyDenomID(denom.Id), bz)
	if len(denom.Name) > 0 {
		store.Set(types.KeyDenomName(denom.Name), []byte(denom.Id))
	}

	return nil
}

// SetDenom is responsible for saving the definition of denomID
func (k Keeper) GetDenom(ctx sdk.Context, id string) (denom types.Denom, err error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyDenomID(id))
	if bz == nil || len(bz) == 0 {
		return denom, sdkerrors.Wrapf(types.ErrInvalidDenom, "not found denomID: %s", id)
	}

	k.cdc.MustUnmarshal(bz, &denom)
	return denom, nil
}

// GetDenoms return all the denoms
func (k Keeper) GetDenoms(ctx sdk.Context) (denoms []types.Denom) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyDenomID(""))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var denom types.Denom
		k.cdc.MustUnmarshal(iterator.Value(), &denom)
		denoms = append(denoms, denom)
	}
	return denoms
}

func (k Keeper) GetDenomsByOwner(ctx sdk.Context, address string) ([]types.Denom, error) {

	denoms := k.GetDenoms(ctx)
	if len(denoms) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "unable to find collection")
	}

	var data []types.Denom
	for _, denom := range denoms {
		if strings.EqualFold(address, denom.Creator) {
			data = append(data, denom)
		}
	}

	return data, nil
}
