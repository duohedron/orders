# Orders Microservice

A minimal, Orders microservice built in Go. Implements REST endpoints, PostgreSQL persistence.

## Features

### API Endpoints

- **POST /orders** - Create an order  
- **GET /orders/{id}** - Retrieve an order by ID  
- **GET /healthz** - Health check  

## Getting Started

### 1. Start Postgres + the application

```bash
docker-compose up
```

The API becomes available at: <http://localhost:8080>

## Example requests

### Create an Order

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"item": "Laptop"}'
```

### Fetch an Order

```bash
curl http://localhost:8080/orders/<uuid>
```
