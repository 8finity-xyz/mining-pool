# Mining Pool

INFINITY mining pool implementation. You need to have Docker installed to use this software.

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/8finity-xyz/mining-pool
cd mining-pool
```

2. Create a `.env` file in the project root with your configuration:
```bash
# Network endpoints
INFINITY_POOL_RPC=https://rpc.soniclabs.com
INFINITY_POOL_WS=wss://rpc.soniclabs.com

# Database configuration (modify if needed)
INFINITY_POOL_REDIS_URI=redis://redis
INFINITY_POOL_POSTGRES_URI=postgresql://postgres:mysecretpassword@postgres?sslmode=disable

# Your private key for pool operations
INFINITY_POOL_PRIVATE_KEY=your_private_key
```

3. Start the services:
```bash
docker compose up -d && docker compose logs -f pool 
```

The mining pool will be available at `http://your_ip:18888`

## Configuration

The pool can be configured through environment variables in the `.env` file. Here are the key settings:

- `INFINITY_POOL_RPC`: RPC endpoint for the INFINITY network
- `INFINITY_POOL_WS`: WebSocket endpoint for real-time updates
- `INFINITY_POOL_REDIS_URI`: Redis connection string for caching
- `INFINITY_POOL_POSTGRES_URI`: PostgreSQL connection string for persistent storage
- `INFINITY_POOL_PRIVATE_KEY`: Your private key for pool operations

## Stopping the Pool

To stop all services:
```bash
docker compose down
```

## Development

To rebuild the pool service after making changes:
```bash
docker compose up -d --build pool
```
