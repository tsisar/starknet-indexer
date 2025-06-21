package model

import (
	"context"
	"github.com/Tsisar/starknet-indexer/internal/storage/generic"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID              string         `gorm:"primaryKey;column:id"`
	EventIndex      int            `gorm:"column:event_index"`
	EventName       string         `gorm:"column:event_name"`
	TxIndex         int            `gorm:"column:tx_index"`
	TxHash          string         `gorm:"column:tx_hash"`
	Transaction     Transaction    `gorm:"foreignKey:TxHash"`
	BlockNumber     uint64         `gorm:"column:block_number"`
	Block           Block          `gorm:"foreignKey:BlockNumber"`
	ContractAddress string         `gorm:"column:contract_address"`
	Contract        Contract       `gorm:"foreignKey:ContractAddress"`
	RawData         datatypes.JSON `gorm:"column:raw_data;type:jsonb"`
	Timestamp       uint64         `gorm:"column:timestamp"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;autoUpdateTime"`
}

func (Event) TableName() string {
	return "starknet.events"
}

func (e *Event) Init() {
}

func (e *Event) GetID() string {
	return e.ID
}

func (e *Event) Load(ctx context.Context, db *gorm.DB) (bool, error) {
	return generic.Load(ctx, db, "id", e.ID, e, "Transaction", "Block", "Contract")
}

func (e *Event) Save(ctx context.Context, db *gorm.DB) error {
	return generic.Save(ctx, db, "id", e)
}

func GetEventsByBlock(ctx context.Context, db *gorm.DB, blockNumber uint64) ([]Event, error) {
	var events []Event
	err := db.WithContext(ctx).
		Where("block_number = ?", blockNumber).
		Preload("Transaction").
		Preload("Block").
		Preload("Contract").
		Order("tx_index ASC, event_index ASC").
		Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}
