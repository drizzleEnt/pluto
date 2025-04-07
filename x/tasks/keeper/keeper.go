package keeper

import (
	"cosmossdk.io/store/prefix"
	"cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group/errors"
	mytypes "github.com/drizzleent/pluto/x/tasks/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey types.KVStoreKey
}

func NewKeeper(cdc codec.BinaryCodec, store types.KVStoreKey) *Keeper {
	return &Keeper{
		storeKey: store,
		cdc:      cdc,
	}
}

func (k *Keeper) SetTask(ctx sdk.Context, creator string, description string) {
	store := prefix.NewStore(ctx.KVStore(&k.storeKey), []byte("tasks"))
	id := k.getNextID(ctx)
	task := mytypes.Task{
		Id:          id,
		Creator:     creator,
		Description: description,
		Completed:   false,
	}

	b := k.cdc.MustMarshal(&task)
	store.Set(sdk.Uint64ToBigEndian(id), b)
	k.setNextID(ctx, id+1)
}

func (k *Keeper) CompleteTask(ctx sdk.Context, creator string, id uint64) error {
	store := prefix.NewStore(ctx.KVStore(&k.storeKey), []byte("tasks"))
	bz := store.Get(sdk.Uint64ToBigEndian(id))
	if bz == nil {
		return errors.ErrEmpty.Wrapf("task %d not found", id)
	}

	var task mytypes.Task

	err := k.cdc.Unmarshal(bz, &task)
	if err != nil {
		return errors.ErrEmpty.Wrap(err.Error())
	}

	if task.Creator != creator {
		return errors.ErrUnauthorized.Wrapf("not task creator")
	}

	task.Completed = true
	bz = k.cdc.MustMarshal(&task)
	store.Set(sdk.Uint64ToBigEndian(id), bz)

	return nil
}

func (k *Keeper) GetTask() {}

func (k Keeper) getNextID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(&k.storeKey)
	bz := store.Get([]byte("nextTaskID"))
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) setNextID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(&k.storeKey)
	store.Set([]byte("nextTaskID"), sdk.Uint64ToBigEndian(id))
}
