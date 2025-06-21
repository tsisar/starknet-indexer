package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/vault_wbtc"
)

var VaultWbtcRegistry = map[string]EventMapper{
	"stablecoin::core::vault::Vault::Deposit":  makeMapper[vault_wbtc.Deposit](nil),
	"stablecoin::core::vault::Vault::Withdraw": makeMapper[vault_wbtc.Withdraw](nil),
}
