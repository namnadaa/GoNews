# GoNews 

**GoNews** is a simple news server implemented in Go. It provides a REST API for managing publications and supports various data storage backends: PostgreSQL, MongoDB, or an in-memory store.

## Features

- Retrieve a list of posts
- Create, update, and delete posts
- Supports multiple storage backends (PostgreSQL, MongoDB, in-memory)
- Clean modular architecture
- Simple configuration via environment variables

## Project Structure

```
GoNews/
│
├── cmd/  
│   └── server/              # Application entry point (server.go)
│
├── pkg/
│   ├── api/                 # HTTP API handlers and routing logic
│   └── storage/             # Storage interface and implementations
│       ├── memdb/           # In-memory storage implementation
│       ├── mongo/           # MongoDB storage implementation
│       └── postgres/        # PostgreSQL storage implementation
│
├── docker-compose.yaml      # Docker Compose configuration for running PostgreSQL and MongoDB as services
├── go.mod                   # Go module definition file  
├── go.sum                   # Dependency checksums file
├── README.md                # Project documentation
└── schema.md                # Data schema description
```

## Configuration

Specify the storage type using environment variables:

```bash
# Storage type: memory, postgres, or mongo
# Run with in-memory storage
STORAGE=memory go run ./cmd/server

# Run with PostgreSQL
STORAGE=postgres POSTGRES_DSN=postgres://myuser:mypassword@localhost:5434/mydb go run ./cmd/server

# Run with MongoDB
STORAGE=mongo MONGO_DSN=mongodb://localhost:27017/ MONGO_DB=data MONGO_COLLECTION=languages go run ./cmd/server
```