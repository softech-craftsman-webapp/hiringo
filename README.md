# Hiringo API
Main API endpoints of the Hiringo API. (The core features, it is cannot changeable)

## Metrics (Prometheus)
```
${API_URL}/metrics
```

## OpenAPI
```
${API_URL}/openapi
```

## How to start?

### Step 1. Install all dependencies
```
go mod vendor
```

### Step 2. Create an environment variable file
```
cp .env.example .env
```

Note: Update an environment variable file with your own values.

## Step 3. Get Public key from Auth Server
```
bash ./public_key.sh
```

Note: You have to update constant variables inside ```crypto/const.go``` file. By default it is set to use ECDSA (ES256).

### Step 4. Run the server

#### Virtual Environment (Production)
```
docker-compose up --build -d
```

Note: Base path in docker environment is ```/app``` path.

#### Basic (Development)
```
go run main.go
```