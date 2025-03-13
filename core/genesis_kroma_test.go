package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
)

func TestKromaChainsGenesis(t *testing.T) {
	tests := []struct {
		chainID      uint64
		expectedHash common.Hash
	}{
		{
			chainID:      params.KromaMainnetChainID,
			expectedHash: common.HexToHash("0xeab1dbcbd854942126643609f6b457e391b169c819b7e5d5042389ccf6012cbf"),
		},
		{
			chainID:      params.KromaSepoliaChainID,
			expectedHash: common.HexToHash("0x52ef8f66bb31c16326eb2072dd9b2fa734068728b845d5428f3a256a50bf252e"),
		},
		{
			chainID:      params.KromaDevnetChainID,
			expectedHash: common.HexToHash("0x1acfe78cf3b3278ca47f9d51d59d3c47612c8899085815ae08443125adba735f"),
		},
	}
	for _, tt := range tests {
		db := rawdb.NewMemoryDatabase()
		tdb := triedb.NewDatabase(db, defaultCacheConfig.triedbConfig(false))

		genesis, err := LoadOPStackGenesis(tt.chainID)
		require.NoError(t, err)
		require.NotNil(t, genesis.Config.Optimism)
		require.True(t, genesis.Config.Zktrie)
		overrides := &ChainOverrides{ApplySuperchainUpgrades: true}

		_, hash, err := SetupGenesisBlockWithOverride(db, tdb, genesis, overrides)
		require.NoError(t, err)
		require.Equal(t, tt.expectedHash, hash)
	}
}
