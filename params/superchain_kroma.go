package params

import "github.com/ethereum-optimism/superchain-registry/superchain"

const (
	KromaMainnetChainID = 255
	KromaSepoliaChainID = 2358
	KromaDevnetChainID  = 7791
)

const (
	KromaMainnetBedrockBlock = 22684807
	KromaSepoliaBedrockBlock = 24190434
	KromaDevnetBedrockBlock  = 12347964
)

var KromaChains = map[uint64]*superchain.ChainConfig{
	KromaMainnetChainID: {
		Chain:      "kroma-mainnet",
		Name:       "Kroma Mainnet",
		Superchain: "kroma",
		ChainID:    KromaMainnetChainID,
		HardForkConfiguration: superchain.HardForkConfiguration{
			CanyonTime:  uint64ptr(1708502400),
			EcotoneTime: uint64ptr(1714032001),
		},
		Genesis: superchain.ChainGenesis{
			L1: superchain.BlockID{
				Hash:   mustHexToHash("0xe459c500b760ed52a1ad799bf578b257af2c76f6ebe061a4c62627e9c605bced"),
				Number: 18067255,
			},
			L2: superchain.BlockID{
				Hash:   mustHexToHash("0xeab1dbcbd854942126643609f6b457e391b169c819b7e5d5042389ccf6012cbf"),
				Number: 0,
			},
			L2Time:    1693880387,
			ExtraData: mustHexToHexBytes("0x4c696d69746c657373205765623320556e6976657273653a204b726f6d61"),
			SystemConfig: superchain.SystemConfig{
				BatcherAddr:       superchain.MustHexToAddress("0x41b8cd6791de4d8f9e0eaf7861ac506822adce12"),
				Overhead:          mustHexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc"),
				Scalar:            mustHexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0"),
				GasLimit:          30000000,
				BaseFeeScalar:     nil,
				BlobBaseFeeScalar: nil,
			},
		},
		Optimism: &superchain.OptimismConfig{
			EIP1559Elasticity:  6,
			EIP1559Denominator: 50,
		},
	},
	KromaSepoliaChainID: {
		Chain:      "kroma-sepolia",
		Name:       "Kroma Sepolia",
		Superchain: "kroma",
		ChainID:    KromaSepoliaChainID,
		HardForkConfiguration: superchain.HardForkConfiguration{
			CanyonTime:  uint64ptr(1707897600),
			EcotoneTime: uint64ptr(1713340800),
		},
		Genesis: superchain.ChainGenesis{
			L1: superchain.BlockID{
				Hash:   mustHexToHash("0x936e490e33e6e136ecd9095090e30ed7def3903ef2bae3e05966b376e493ad76"),
				Number: 3841490,
			},
			L2: superchain.BlockID{
				Hash:   mustHexToHash("0x52ef8f66bb31c16326eb2072dd9b2fa734068728b845d5428f3a256a50bf252e"),
				Number: 0,
			},
			L2Time:    1688709132,
			ExtraData: &superchain.HexBytes{},
			SystemConfig: superchain.SystemConfig{
				BatcherAddr:       superchain.MustHexToAddress("0xf15dc770221b99c98d4aaed568f2ab04b9d16e42"),
				Overhead:          mustHexToHash("0x0000000000000000000000000000000000000000000000000000000000000834"),
				Scalar:            mustHexToHash("0x000000000000000000000000000000000000000000000000000000000016e360"),
				GasLimit:          30000000,
				BaseFeeScalar:     nil,
				BlobBaseFeeScalar: nil,
			},
		},
		Optimism: &superchain.OptimismConfig{
			EIP1559Elasticity:  6,
			EIP1559Denominator: 50,
		},
	},
	KromaDevnetChainID: {
		Chain:      "kroma-devnet",
		Name:       "Kroma Devnet",
		Superchain: "kroma",
		ChainID:    KromaDevnetChainID,
		HardForkConfiguration: superchain.HardForkConfiguration{
			CanyonTime:  uint64ptr(1707292800),
			EcotoneTime: uint64ptr(1712908800),
		},
		Genesis: superchain.ChainGenesis{
			L1: superchain.BlockID{
				Hash:   mustHexToHash("0x160a43453346b65e861074053112f4c002b263e3260ba25840b37369d574e118"),
				Number: 1192134,
			},
			L2: superchain.BlockID{
				Hash:   mustHexToHash("0x1acfe78cf3b3278ca47f9d51d59d3c47612c8899085815ae08443125adba735f"),
				Number: 0,
			},
			L2Time:    1711098072,
			ExtraData: mustHexToHexBytes("0x4c696d69746c657373205765623320556e6976657273653a204b726f6d61"),
			SystemConfig: superchain.SystemConfig{
				BatcherAddr:       superchain.MustHexToAddress("0xde59a71027b7605deb314a39a8455f1d69b7a9a3"),
				Overhead:          mustHexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc"),
				Scalar:            mustHexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0"),
				GasLimit:          30000000,
				BaseFeeScalar:     nil,
				BlobBaseFeeScalar: nil,
			},
		},
		Optimism: &superchain.OptimismConfig{
			EIP1559Elasticity:        6,
			EIP1559Denominator:       50,
			EIP1559DenominatorCanyon: uint64ptr(250),
		},
	},
}

func IsKromaChain(chainID uint64) bool {
	return chainID == KromaMainnetChainID || chainID == KromaSepoliaChainID || chainID == KromaDevnetChainID
}

func mustHexToHexBytes(hex string) *superchain.HexBytes {
	h := &superchain.HexBytes{}
	_ = h.UnmarshalText([]byte(hex))
	return h
}

func mustHexToHash(hex string) superchain.Hash {
	h := superchain.Hash{}
	_ = h.UnmarshalText([]byte(hex))
	return h
}
