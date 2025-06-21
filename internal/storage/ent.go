package storage

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/Tsisar/starknet-indexer/generated/ent"
)

func InitEnt(ctx context.Context) (*ent.Client, error) {
	driver, err := sql.Open(dialect.Postgres, DSNFromConfig())
	if err != nil {
		return nil, fmt.Errorf("sql open failed: %w", err)
	}

	client := ent.NewClient(ent.Driver(driver))

	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return client, nil
}

func TruncateAllTables(ctx context.Context, client *ent.Client) error {
	if _, err := client.PositionActivity.Delete().Exec(ctx); err != nil {
		return err
	}
	if _, err := client.Position.Delete().Exec(ctx); err != nil {
		return err
	}
	if _, err := client.Pool.Delete().Exec(ctx); err != nil {
		return err
	}
	if _, err := client.ProtocolStat.Delete().Exec(ctx); err != nil {
		return err
	}
	if _, err := client.User.Delete().Exec(ctx); err != nil {
		return err
	}
	return nil
}
