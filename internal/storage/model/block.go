package model

import (
	"context"
	"github.com/Tsisar/starknet-indexer/internal/storage/generic"
	"gorm.io/gorm"
	"time"
)

type Block struct {
	Number    uint64    `gorm:"primaryKey;column:number"`
	Timestamp uint64    `gorm:"column:timestamp"`
	Hash      string    `gorm:"column:hash;uniqueIndex"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Block) TableName() string {
	return "starknet.blocks"
}

func (b *Block) Init() {
}

func (b *Block) GetID() uint64 {
	return b.Number
}

func (b *Block) Load(ctx context.Context, db *gorm.DB) (bool, error) {
	return generic.Load(ctx, db, "number", b.Number, b)
}

func (b *Block) Save(ctx context.Context, db *gorm.DB) error {
	return generic.Save(ctx, db, "number", b)
}

func GetAllBlocks(ctx context.Context, db *gorm.DB) ([]Block, error) {
	var blocks []Block
	err := db.WithContext(ctx).
		Model(&Block{}).
		Order("number ASC").
		Find(&blocks).Error
	if err != nil {
		return nil, err
	}
	return blocks, nil
}
