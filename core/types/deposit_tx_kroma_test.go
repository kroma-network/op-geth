package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
)

func TestIsKromaLegacyDepositTx(t *testing.T) {
	depTx := NewTx(&DepositTx{
		SourceHash:          common.Hash{31: 1},
		From:                common.Address{0: 1},
		To:                  &common.Address{0: 2},
		Mint:                nil,
		Value:               common.Big0,
		Gas:                 21000,
		IsSystemTransaction: false,
		Data:                nil,
	})
	kromaDepTx := NewTx(&KromaLegacyDepositTx{
		SourceHash: common.Hash{31: 1},
		From:       common.Address{0: 1},
		To:         &common.Address{0: 2},
		Mint:       nil,
		Value:      common.Big0,
		Gas:        21000,
		Data:       nil,
	})

	depTxBytes, err := depTx.MarshalBinary()
	require.NoError(t, err)
	isKromaDepTx, err := IsKromaLegacyDepositTx(depTxBytes[1:])
	require.NoError(t, err)
	require.False(t, isKromaDepTx)

	kromaDepTxBytes, err := kromaDepTx.MarshalBinary()
	require.NoError(t, err)
	isKromaDepTx, err = IsKromaLegacyDepositTx(kromaDepTxBytes[1:])
	require.NoError(t, err)
	require.True(t, isKromaDepTx)
}

func TestKromaLegacyDepositTxHash(t *testing.T) {
	txHash := common.HexToHash("83dfcc10fc10f667b6e36009611aa4ff570cbd2e86159cb8f131dc4699dc52f5")
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	data := common.Hex2Bytes("efc674eb000000000000000000000000000000000000000000000000000000000113af370000000000000000000000000000000000000000000000000000000064f69043000000000000000000000000000000000000000000000000000000026717271fe459c500b760ed52a1ad799bf578b257af2c76f6ebe061a4c62627e9c605bced000000000000000000000000000000000000000000000000000000000000000100000000000000000000000041b8cd6791de4d8f9e0eaf7861ac506822adce1200000000000000000000000000000000000000000000000000000000000000bc00000000000000000000000000000000000000000000000000000000000a6fe00000000000000000000000000000000000000000000000000000000000002710")

	kromaDepTx := NewTx(&KromaLegacyDepositTx{
		SourceHash: common.HexToHash("be9873667a9e4d2f4bc74b0e771b674dcb8848e715bdab376be6c6c9c92db8ba"),
		From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
		To:         &toAddr,
		Mint:       nil,
		Value:      common.Big0,
		Gas:        1000000,
		Data:       data,
	})
	require.Equal(t, txHash, kromaDepTx.Hash())

	kromaDepTxBytes, err := kromaDepTx.MarshalBinary()
	require.NoError(t, err)

	var tx Transaction
	err = tx.UnmarshalBinary(kromaDepTxBytes)
	require.NoError(t, err)

	require.Equal(t, txHash, tx.Hash())
}
