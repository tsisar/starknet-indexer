package model

import (
	"context"
	"github.com/Tsisar/starknet-indexer/internal/storage/generic"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	Hash        string    `gorm:"primaryKey;column:hash"`
	Index       int       `gorm:"column:index"`
	BlockNumber uint64    `gorm:"column:block_number"`
	Block       Block     `gorm:"foreignKey:BlockNumber"`
	Timestamp   uint64    `gorm:"column:timestamp"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Transaction) TableName() string {
	return "starknet.transactions"
}

func (t *Transaction) Init() {
}

func (t *Transaction) GetID() string {
	return t.Hash
}

func (t *Transaction) Load(ctx context.Context, db *gorm.DB) (bool, error) {
	return generic.Load(ctx, db, "hash", t.Hash, t, "Block")
}

func (t *Transaction) Save(ctx context.Context, db *gorm.DB) error {
	return generic.Save(ctx, db, "hash", t)
}
