package core

import (
	"github.com/holiman/uint256"
	zktrie "github.com/kroma-network/zktrie/trie"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/triedb"
)

// hashAlloc computes the state root using ZKTrie according to the genesis specification.
func hashAllocZk(ga *types.GenesisAlloc, _ bool) (common.Hash, error) {
	memdb := zktrie.NewZkTrieMemoryDb()
	tr, err := trie.NewZkTrie(common.Hash{}, memdb)
	if err != nil {
		return common.Hash{}, nil
	}

	for addr, account := range *ga {
		acc := &types.StateAccount{
			Nonce:    account.Nonce,
			CodeHash: crypto.Keccak256Hash(account.Code).Bytes(),
			Root:     common.Hash{},
		}
		if account.Balance != nil {
			acc.Balance = uint256.MustFromBig(account.Balance)
		}
		storageTr, err := trie.NewZkTrie(common.Hash{}, memdb)
		if err != nil {
			return common.Hash{}, err
		}
		for key, value := range account.Storage {
			if value != (common.Hash{}) {
				err := storageTr.UpdateStorage(addr, key[:], common.TrimLeftZeroes(value[:]))
				if err != nil {
					return common.Hash{}, err
				}
			}
		}
		acc.Root = storageTr.Hash()
		err = tr.UpdateAccount(addr, acc)
		if err != nil {
			return common.Hash{}, err
		}
	}
	return tr.Hash(), nil
}

// flushAllocZk is very similar with hash, but the main difference is all the
// generated states will be persisted into the given database.
func flushAllocZk(ga *types.GenesisAlloc, triedb *triedb.Database) (common.Hash, error) {
	root, err := hashAllocZk(ga, false)
	if err != nil {
		return common.Hash{}, err
	}
	// Commit newly generated states into disk if it's not empty.
	if root != (common.Hash{}) {
		if err := triedb.Commit(root, true); err != nil {
			return common.Hash{}, err
		}
	}
	return root, nil
}
