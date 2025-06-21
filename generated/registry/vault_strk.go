package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/vault_strk"
)

var VaultStrkRegistry = map[string]EventMapper{
	"stablecoin::core::vault::Vault::Deposit":  makeMapper[vault_strk.Deposit](nil),
	"stablecoin::core::vault::Vault::Withdraw": makeMapper[vault_strk.Withdraw](nil),
}
