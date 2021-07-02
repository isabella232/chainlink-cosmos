package keeper

import (
	"testing"

	"github.com/ChainSafe/chainlink-cosmos/x/chainlink/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

// nolint
func setupKeeper(t testing.TB) (*Keeper, sdk.Context) {
	feedStoreKey := sdk.NewKVStoreKey(types.FeedStoreKey)
	roundStoreKey := sdk.NewKVStoreKey(types.RoundStoreKey)
	moduleStoreKey := sdk.NewKVStoreKey(types.ModuleStoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(feedStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(roundStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(moduleStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	keeper := NewKeeper(codec.NewProtoCodec(registry), feedStoreKey, roundStoreKey, moduleStoreKey, memStoreKey)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return keeper, ctx
}

// func setupKeeper(t testing.TB) (*Keeper, sdk.Context) {
// 	storeKey := sdk.NewKVStoreKey(types.StoreKey)
// 	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

// 	db := tmdb.NewMemDB()
// 	stateStore := store.NewCommitMultiStore(db)
// 	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
// 	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
// 	require.NoError(t, stateStore.LoadLatestVersion())

// 	registry := codectypes.NewInterfaceRegistry()
// 	keeper := NewKeeper(codec.NewProtoCodec(registry), storeKey, memStoreKey)

// 	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
// 	return keeper, ctx
// }
