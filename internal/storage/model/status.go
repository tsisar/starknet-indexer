package model

import (
	"context"
	"github.com/Tsisar/starknet-indexer/internal/storage/generic"
	"gorm.io/gorm"
	"time"
)

type Status struct {
	ID           uint64    `gorm:"primaryKey;column:id"`
	CurrentBlock uint64    `gorm:"column:current_block"`
	Timestamp    uint64    `gorm:"column:timestamp"`
	Hash         string    `gorm:"column:hash;uniqueIndex"`
	Synced       bool      `gorm:"column:synced"`
	HasError     bool      `gorm:"column:has_error"`
	ErrorMsg     string    `gorm:"column:error_msg"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Status) TableName() string {
	return "starknet.status"
}

func (s *Status) Init() {
	s.ID = 1 // Default ID for the status record
	s.CurrentBlock = 0
	s.Synced = false
	s.HasError = false
	s.ErrorMsg = ""
}

func (s *Status) GetID() uint64 {
	return s.ID
}

func (s *Status) Load(ctx context.Context, db *gorm.DB) (bool, error) {
	s.ID = 1 // Default ID for the status record
	return generic.Load(ctx, db, "id", s.ID, s)
}

func (s *Status) Save(ctx context.Context, db *gorm.DB) error {
	s.ID = 1 // Default ID for the status record
	return generic.Save(ctx, db, "id", s)
}

func (s *Status) UpdateCurrentBlock(ctx context.Context, db *gorm.DB, blockNumber uint64) error {
	s.CurrentBlock = blockNumber
	return s.Save(ctx, db)
}

func (s *Status) SetSynced(ctx context.Context, db *gorm.DB, synced bool) error {
	if synced {
		s.HasError = false
		s.ErrorMsg = ""
	}
	s.Synced = synced
	return s.Save(ctx, db)
}

func (s *Status) SetError(ctx context.Context, db *gorm.DB, hasError bool, errorMsg string) error {
	s.HasError = hasError
	s.ErrorMsg = errorMsg
	return s.Save(ctx, db)
}

func (s *Status) UpdateHash(ctx context.Context, db *gorm.DB, hash string, timestamp uint64) error {
	s.Hash = hash
	s.Timestamp = timestamp
	return s.Save(ctx, db)
}
