package core

import (
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/superchain-registry/superchain"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

func LoadKromaGenesis(chainID uint64) (*superchain.Genesis, error) {
	chConfig, ok := params.KromaChains[chainID]
	if !ok {
		return nil, fmt.Errorf("unknown chain ID: %d", chainID)
	}

	return &superchain.Genesis{
		Nonce:         0,
		Timestamp:     chConfig.Genesis.L2Time,
		ExtraData:     *chConfig.Genesis.ExtraData,
		GasLimit:      30000000,
		Difficulty:    (*superchain.HexBig)(new(big.Int)),
		Mixhash:       superchain.Hash{},
		Coinbase:      superchain.Address{},
		Number:        0,
		GasUsed:       0,
		ParentHash:    superchain.Hash{},
		BaseFee:       (*superchain.HexBig)(new(big.Int).SetUint64(1000000000)),
		ExcessBlobGas: nil,
		BlobGasUsed:   nil,
		Alloc:         make(map[superchain.Address]superchain.GenesisAccount),
		StateHash:     nil,
	}, nil

}

func LoadKromaGenesisAlloc(chainID uint64) types.GenesisAlloc {
	if chainID == params.KromaMainnetChainID {
		return decodePrealloc(kromaMainnetAllocData)
	} else if chainID == params.KromaSepoliaChainID {
		return decodePrealloc(kromaSepoliaAllocData)
	} else if chainID == params.KromaDevnetChainID {
		return decodePrealloc(kromaDevnetAllocData)
	}
	return nil
}
