package trigger

import (
	"database/sql"
	"fmt"
)

func SetupCurrentBlockTrigger(db *sql.DB) error {
	const createFunction = `
	CREATE OR REPLACE FUNCTION notify_starknet_current_block_update()
	RETURNS trigger AS $$
	BEGIN
	  PERFORM pg_notify('starknet_current_block_update', NEW.current_block::text);
	  RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;
	`

	const dropTrigger = `
	DROP TRIGGER IF EXISTS current_block_update_trigger ON starknet.status;
	`

	const createTrigger = `
	CREATE TRIGGER current_block_update_trigger
	AFTER UPDATE OF current_block ON starknet.status
	FOR EACH ROW
	WHEN (OLD.current_block IS DISTINCT FROM NEW.current_block)
	EXECUTE FUNCTION notify_starknet_current_block_update();
	`

	if _, err := db.Exec(createFunction); err != nil {
		return fmt.Errorf("create function failed: %w", err)
	}
	if _, err := db.Exec(dropTrigger); err != nil {
		return fmt.Errorf("drop trigger failed: %w", err)
	}
	if _, err := db.Exec(createTrigger); err != nil {
		return fmt.Errorf("create trigger failed: %w", err)
	}

	return nil
}

func SetupSyncedTrigger(db *sql.DB) error {
	const createFunction = `
	CREATE OR REPLACE FUNCTION notify_starknet_synced_update()
	RETURNS trigger AS $$
	BEGIN
	  PERFORM pg_notify('starknet_synced_update', NEW.synced::text);
	  RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;
	`

	const dropTrigger = `
	DROP TRIGGER IF EXISTS synced_update_trigger ON starknet.status;
	`

	const createTrigger = `
	CREATE TRIGGER synced_update_trigger
	AFTER UPDATE OF synced ON starknet.status
	FOR EACH ROW
	WHEN (OLD.synced IS DISTINCT FROM NEW.synced)
	EXECUTE FUNCTION notify_starknet_synced_update();
	`

	if _, err := db.Exec(createFunction); err != nil {
		return fmt.Errorf("create function (synced) failed: %w", err)
	}
	if _, err := db.Exec(dropTrigger); err != nil {
		return fmt.Errorf("drop trigger (synced) failed: %w", err)
	}
	if _, err := db.Exec(createTrigger); err != nil {
		return fmt.Errorf("create trigger (synced) failed: %w", err)
	}

	return nil
}
