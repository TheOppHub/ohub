package core

import (
	hexPkg "encoding/hex"
	"math/big"

	"github.com/ethereumproject/go-ethereum/logger/glog"
)

var (
	DefaultConfigMainnet *SufficientChainConfig
	DefaultConfigMorden  *SufficientChainConfig
)

func init() {

	var err error
	var nullHashArray [32]byte
	nullHashSlice, err := hexPkg.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		glog.Fatal("Unable to convert slice to array!")
	}
	copy(nullHashArray[:], nullHashSlice[:32])
	DefaultConfigMainnet = &SufficientChainConfig{
		Name:      "OHUB Mainnet",
		Identity:  "omain",
		Network:   1867,
		Consensus: "ethash",
		Genesis: &GenesisDump{
			Nonce:      "0x0000000000000033",
			Timestamp:  "0x00",
			ExtraData:  "0x54686520426C61636B20446F6C6C6172206973204F6E65206F6620746865204D6F737420496E666C75656E7469616C20696E20416D6572696361",
			ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
			GasLimit:   "0x2fefd8",
			Difficulty: "0x5000",
			Mixhash:    "0x0000000000000000000000000000000000000000000000000000000000000000",
			Coinbase:   "0x0000000000000000000000000000000000000000",
			Alloc: map[hex]*GenesisDumpAlloc{
				"915121963753C2bc63f480198aCa9613A7f42D46": &GenesisDumpAlloc{
					Balance: "100000",
				},
			},
		},
		ChainConfig: &ChainConfig{
			Forks: Forks{
				&Fork{
					Name:         "Defuse Ethereum Difficulty Bomb",
					Block:        big.NewInt(0),
					RequiredHash: nullHashArray,
					Features: []*ForkFeature{
						&ForkFeature{
							ID: "difficulty",
							Options: ChainFeatureConfigOptions{
								"type": "defused",
							},
						},
					},
				},
			},
			BadHashes: []*BadHash{
				&BadHash{
					Block: big.NewInt(0),
					Hash:  nullHashArray,
				},
			},
		},
		Bootstrap: nil, // TODO: Add Bootnodes here []string{}*/
	}
	DefaultConfigMorden = &SufficientChainConfig{
		Name:      "OHUB Testnet",
		Identity:  "otest",
		Network:   1112,
		Consensus: "ethash",
		Genesis: &GenesisDump{
			Nonce:      "0x0000000000000034",
			Timestamp:  "0x00",
			ExtraData:  "0x00",
			ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
			GasLimit:   "0x2fefd8",
			Difficulty: "0x1000",
			Mixhash:    "0x0000000000000000000000000000000000000000000000000000000000000000",
			Coinbase:   "0x0000000000000000000000000000000000000000",
			Alloc: map[hex]*GenesisDumpAlloc{
				"915121963753C2bc63f480198aCa9613A7f42D46": &GenesisDumpAlloc{
					Balance: "100000",
				},
			},
		},
		ChainConfig: &ChainConfig{
			Forks: Forks{
				&Fork{
					Name:         "Defuse Ethereum Difficulty Bomb",
					Block:        big.NewInt(0),
					RequiredHash: nullHashArray,
					Features: []*ForkFeature{
						&ForkFeature{
							ID: "difficulty",
							Options: ChainFeatureConfigOptions{
								"type": "defused",
							},
						},
					},
				},
			},
			BadHashes: []*BadHash{
				&BadHash{
					Block: big.NewInt(0),
					Hash:  nullHashArray,
				},
			},
		},
		Bootstrap: nil, // TODO: Add Bootnodes here []string{}*/
	}
}
