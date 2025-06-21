package model

import (
	"context"
	"github.com/tsisar/starknet-indexer/internal/storage/generic"
	"gorm.io/gorm"
	"time"
)

type Contract struct {
	Address   string    `gorm:"primaryKey;column:address"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Contract) TableName() string {
	return "starknet.contracts"
}

func (c *Contract) Init() {
}

func (c *Contract) GetID() string {
	return c.Address
}

func (c *Contract) Load(ctx context.Context, db *gorm.DB) (bool, error) {
	return generic.Load(ctx, db, "address", c.Address, c)
}

func (c *Contract) Save(ctx context.Context, db *gorm.DB) error {
	return generic.Save(ctx, db, "address", c)
}
