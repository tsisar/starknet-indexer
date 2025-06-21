package storage

import (
	"fmt"
	"github.com/Tsisar/starknet-indexer/internal/config"
)

func DSNFromConfig() string {
	cfg := config.App.Postgres
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=UTC",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB,
	)
}
