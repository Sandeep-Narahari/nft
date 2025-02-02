package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Sandeep-Narahari/nft/types"
)

// GetOwner gets all the ID Collections owned by an address and denomID
func (k Keeper) GetOwner(ctx sdk.Context, address sdk.AccAddress, denom string) types.Owner {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(address, denom, ""))
	defer iterator.Close()

	owner := types.Owner{
		Address:       address.String(),
		IDCollections: types.IDCollections{},
	}
	idsMap := make(map[string][]string)

	for ; iterator.Valid(); iterator.Next() {
		_, denom, tokenID, _ := types.SplitKeyOwner(iterator.Key())
		if ids, ok := idsMap[denom]; ok {
			idsMap[denom] = append(ids, tokenID)
		} else {
			idsMap[denom] = []string{tokenID}
			owner.IDCollections = append(owner.IDCollections, types.IDCollection{
				DenomId: denom,
			})
		}
	}

	for i := 0; i < len(owner.IDCollections); i++ {
		owner.IDCollections[i].NftIds = idsMap[owner.IDCollections[i].DenomId]
	}
	return owner
}

// GetOwner gets all the ID Collections
func (k Keeper) GetOwners(ctx sdk.Context) (owners types.Owners) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.KeyOwner(nil, "", ""))
	defer iterator.Close()

	idcsMap := make(map[string]types.IDCollections)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		address, denom, id, _ := types.SplitKeyOwner(key)
		if _, ok := idcsMap[address.String()]; !ok {
			idcsMap[address.String()] = types.IDCollections{}
			owners = append(owners, types.Owner{
				Address: address.String(),
			})
		}
		idcs := idcsMap[address.String()]
		idcs = idcs.Add(denom, id)
		idcsMap[address.String()] = idcs
	}
	for i, owner := range owners {
		owners[i].IDCollections = idcsMap[owner.Address]
	}
	return owners
}

func (k Keeper) deleteOwner(ctx sdk.Context,
	denomID, tokenID string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyOwner(owner, denomID, tokenID))
}

func (k Keeper) setOwner(ctx sdk.Context,
	denomID, tokenID string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	bz := types.MustMarshalTokenID(k.cdc, tokenID)
	store.Set(types.KeyOwner(owner, denomID, tokenID), bz)
}

func (k Keeper) swapOwner(ctx sdk.Context,
	denomID, tokenID string,
	srcOwner, dstOwner sdk.AccAddress) {

	// delete old owner key
	k.deleteOwner(ctx, denomID, tokenID, srcOwner)

	// set new owner key
	k.setOwner(ctx, denomID, tokenID, dstOwner)
}

func (k Keeper) GetOwnerNFTs(ctx sdk.Context, owner sdk.AccAddress) []types.OwnerNFTCollection {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(owner, "", ""))
	owner1 := types.Owner{
		Address:       owner.String(),
		IDCollections: types.IDCollections{},
	}
	idsMap := make(map[string][]string)
	var ownerCollections []types.OwnerNFTCollection

	for ; iterator.Valid(); iterator.Next() {
		_, denomID, tokenID, _ := types.SplitKeyOwner(iterator.Key())
		if ids, ok := idsMap[denomID]; ok {
			idsMap[denomID] = append(ids, tokenID)
		} else {
			idsMap[denomID] = []string{tokenID}
			owner1.IDCollections = append(
				owner1.IDCollections,
				types.IDCollection{
					DenomId: denomID,
				},
			)
		}
	}

	for i := 0; i < len(owner1.IDCollections); i++ {
		owner1.IDCollections[i].NftIds = idsMap[owner1.IDCollections[i].DenomId]
		denom, _ := k.GetDenom(ctx, owner1.IDCollections[i].DenomId)
		var nfts []types.NFT
		for _, nftId := range owner1.IDCollections[i].NftIds {
			nft, _ := k.GetNFT(ctx, denom.Id, nftId)
			nfts = append(nfts, nft.(types.NFT))
		}

		ownerCollection := types.OwnerNFTCollection{
			Denom: denom,
			Nfts:  nfts,
		}

		ownerCollections = append(ownerCollections, ownerCollection)
	}

	return ownerCollections
}
