# Starknet Stablecoin Indexer

A high-performance indexer for Starknet stablecoin protocol, built with Go. This project provides a GraphQL API for querying protocol data, including positions, pools, users, and swap events.

## Features

- Real-time indexing of Starknet stablecoin protocol events
- GraphQL API for querying protocol data
- Support for filtering and sorting of all entities
- PostgreSQL database for data storage
- Docker support for easy deployment
- Comprehensive test coverage

## Prerequisites

- Go 1.24.2 or higher
- PostgreSQL 12 or higher
- Docker (optional, for containerized deployment)
- Make (for using Makefile commands)

## Project Structure

```
.
├── abi/              # Smart contract ABIs
├── cmd/              # Command-line applications
├── generated/        # Generated code (ent, GraphQL)
├── graphql/          # GraphQL schema and resolvers
├── internal/         # Internal packages
│   ├── config/      # Configuration
│   ├── indexer/     # Indexer implementation
│   ├── mapper/      # Data mapping
│   └── storage/     # Database storage
├── Makefile         # Build and development commands
├── Dockerfile       # Docker configuration
└── go.mod           # Go module definition
```

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/tsisar/starknet-indexer.git
cd starknet-indexer
```

2. Install dependencies:
```bash
make get
```

3. Generate code:
```bash
make generate
```

4. Build the application:
```bash
make build
```

5. Run the application:
```bash
./indexer
```

## Development

### Available Make Commands

- `make help` - Show available commands
- `make format` - Format Go code
- `make lint` - Run linter
- `make test` - Run tests
- `make get` - Get dependencies
- `make build` - Build the application
- `make image` - Build Docker image
- `make push` - Push Docker image to registry
- `make clean` - Clean build artifacts
- `make dev` - Development build
- `make release` - Production release build and push

### Code Generation

The project uses several code generators:

- `make generate-indexer` - Generate core indexer code
- `make generate-filters` - Generate filter code
- `make generate-ent` - Generate ent code
- `make generate-gqlgen` - Generate GraphQL schema
- `make generate-mapper` - Generate mapper code
- `make generate` - Run all generators

## Docker Support

Build and run with Docker:

```bash
# Build the image
make image

# Run the container
docker run -p 8080:8080 intothefathom/starknet-indexer.stablecoin.indexer:latest
```

## GraphQL API

The project provides a GraphQL API for querying protocol data. The API supports:

- Querying positions, pools, users, and swap events
- Filtering by various fields
- Sorting by any field
- Pagination

Example query:

```graphql
query {
  positions(
    where: {
      positionStatus: SAFE
    }
    orderBy: {
      field: TVL
      direction: DESC
    }
    first: 10
  ) {
    id
    positionAddress
    lockedCollateral
    debtValue
    tvl
  }
}
```
