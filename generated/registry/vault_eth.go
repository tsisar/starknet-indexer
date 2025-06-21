package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/vault_eth"
)

var VaultEthRegistry = map[string]EventMapper{
	"stablecoin::core::vault::Vault::Deposit":  makeMapper[vault_eth.Deposit](nil),
	"stablecoin::core::vault::Vault::Withdraw": makeMapper[vault_eth.Withdraw](nil),
}
