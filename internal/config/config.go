package config

import (
	"github.com/joho/godotenv"
	"github.com/tsisar/extended-log-go/log"
	"os"
)

var App *config

type config struct {
	IndexerName             string
	ResumeFromLastSignature bool
	RPCEndpoint             string
	RPCWSEndpoint           string
	StartBlock              uint64
	Version                 string
	Contracts               map[string]string
	Postgres                postgres
	Metrics                 metrics
}

type postgres struct {
	User     string
	Password string
	DB       string
	Host     string
	Port     string
}

type metrics struct {
	Enabled bool
	Port    string
}

func init() {
	if os.Getenv("RUNNING_IN_CONTAINER") != "true" {
		if err := godotenv.Load(); err == nil {
			log.Info(".env file successfully loaded")
		}
	}

	var err error
	App, err = loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
}

func loadConfig() (*config, error) {
	return &config{
		IndexerName:             getString("INDEXER_NAME", ""),
		ResumeFromLastSignature: getBool("RESUME_FROM_LAST_SIGNATURE", false),
		RPCEndpoint:             getString("RPC_ENDPOINT", "https://api.mainnet-beta.solana.com"),
		RPCWSEndpoint:           getString("RPC_WS_ENDPOINT", "wss://api.mainnet-beta.solana.com"),
		StartBlock:              getUint64("START_BLOCK", 0),
		Version:                 getString("VERSION", "v.unknown"),
		Contracts:               loadContractAddressesFromEnv(),
		Postgres: postgres{
			User:     getString("POSTGRES_USER", "postgres"),
			Password: getString("POSTGRES_PASSWORD", "postgres"),
			DB:       getString("POSTGRES_DB", "indexer"),
			Host:     getString("POSTGRES_HOST", "localhost"),
			Port:     getString("POSTGRES_PORT", "5432"),
		},
		Metrics: metrics{
			Enabled: getBool("METRICS_ENABLED", true),
			Port:    getString("METRICS_PORT", "8040"),
		},
	}, nil
}
