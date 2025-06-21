package graphql

import (
	"github.com/Tsisar/starknet-indexer/generated/ent"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *ent.Client
	DB     *gorm.DB
}
