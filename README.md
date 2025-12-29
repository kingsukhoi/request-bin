# Request Bin

A lightweight HTTP request inspection tool for capturing, viewing, and debugging webhooks and HTTP requests. Request Bin
provides a simple interface to inspect incoming HTTP requests including headers, query parameters, and request bodies.

## Features

- Capture all HTTP requests sent to the server
- View request details including:
    - HTTP method and path
    - Headers
    - Query parameters
    - Request body (with gzip decompression support)
    - Source IP address
    - Response codes
    - Timestamps
- Web-based UI for browsing captured requests
- JWT-based authentication
- PostgreSQL storage with automated migrations
- Optional TLS/HTTPS support
- Docker support

## Architecture

- **Backend**: Go (Gin framework)
- **Frontend**: React + TypeScript (Vite)
- **Database**: PostgreSQL 17
- **Authentication**: JWT with Argon2id password hashing

## Prerequisites

- Go 1.24.0 or higher
- Node.js 22 or higher
- PostgreSQL 17
- pnpm (for frontend package management)

## Quick Start

### Using Docker Compose

1. Start the PostgreSQL database:

```bash
docker-compose up -d
```

2. Set the database URL:

```bash
export DB_URL="postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
```

3. Build and run the backend:

```bash
go run cmd/web/main.go
```

4. Build the frontend:

```bash
cd frontend
pnpm install
pnpm run build
```

The application will be available at `http://localhost:8080`

## Configuration

The application can be configured using environment variables or a YAML configuration file.

### Environment Variables

| Variable               | Description                       | Default           |
|------------------------|-----------------------------------|-------------------|
| `DB_URL`               | PostgreSQL connection string      | Required          |
| `FRONT_END_PATH`       | Path to frontend dist folder      | `./frontend/dist` |
| `TLS_CERT_PATH`        | Path to TLS certificate           | -                 |
| `TLS_PRIVATE_KEY_PATH` | Path to TLS private key           | -                 |
| `TLS_PORT`             | Port for HTTPS server             | `0.0.0.0:8443`    |
| `CUSTOM_ROUTES`        | Custom routes for request capture | -                 |

### Example YAML Configuration

```yaml
db_url: postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable
front_end_path: ./frontend/dist
tls:
  cert_path: /path/to/cert.pem
  key_path: /path/to/key.pem
  port: 0.0.0.0:8443
```

## Development

### Backend Development

```bash
# Run the server
go run cmd/web/main.go

# Run migrations manually
go run helpers/seedDb/main.go
```

### Frontend Development

```bash
cd frontend

# Install dependencies
pnpm install

# Run development server
pnpm run dev

# Build for production
pnpm run build

# Lint code
pnpm run lint
```

## Database Migrations

Migrations are automatically run on application startup using [dbmate](https://github.com/amacneil/dbmate). Migration
files are located in `db/migrations/`.

## API Endpoints

- `GET /health` - Health check endpoint
- `POST /api/auth/login` - User authentication
- `GET /api/requests` - Get captured requests (paginated)
- `GET /api/requests/:id` - Get specific request details
- `GET /api/requests/:id/headers` - Get request headers
- `GET /api/requests/:id/query-params` - Get query parameters

## License

See [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
