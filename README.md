# silver-bassoon

## Whats going on here?

- [Stack used 🥞](#stack-used-)
- [Setup for Local 🛠](#setup-)
  - [Database Initialization](DATABASE_INITIALIZATION.md)
- [Migration and ORM](#about-sqlc)
  - [SQLC](#about-sqlc)
- [How to Run 👟](#how-to-run-)
  - [Run Locally 🏃](#run-locally-)
  - [Dockering 🐳](#dockering-)
    - [Docker Compose](#using-docker-compose-Recommended)
- [Testing 🧪](#testing-)
  - [Golang](#backend-unit-testing)
  - [VueJS](#frontend-unit-testing)

## Stack used 🥞

- PostgreSQL
- Golang 1.21
- SQLC ([jump](#about-sqlc))
- VueJS
- Vite (bundler)

## Setup 🛠

After cloning the repo, you need to install Go dependencies. Run the following command:

```bash
go mod tidy
```

Then, install NodeJS dependencies. Run the following command:

```bash
yarn install
```

## Initialize the Database

For database initialization section you can found it here: [Click the link](DATABASE_INITIALIZATION.md)

## About SQLC

By using [sqlc](https://sqlc.dev/), this app does not require you to write the models mapping of your table into Golang native structs yourself.

Moreover, you can write a raw SQL query in a `.sql` file and convert it into a Golang native `func` by placing the `.sql` file under the `scripts/queries` directory.

```bash
├── cmd
│   ├── migrate
│   │   ├── main.go
│   │   └── migrations
│   │       └── 00_initial.up.sql
├── scripts
│   └── queries
│       ├── order_items.sql (EXAMPLE)
│       └── orders.sql (EXAMPLE)
```

### Generate `struct` as models and `func` as queries

Since you already have your migration and query files, the next step is to generate the `struct` and `func` for it by running this command:

```bash
go run cmd/sqlc/main.go
```

This will generate `.go` files with an expected output like this:

```bash
├── internal
│   └── repo
│       └── psql
│           ├── db.go
│           ├── models.go
│           ├── order_items.sql.go
│           ├── orders.sql.go
│           └── querier.go
```

You can read the file as well by:

```bash
cat internal/repo/psql/models.go
```

## How to Run 👟

### Run Locally 🏃

**_NOTE: If you've done `Database Initialization` steps from [here](DATABASE_INITIALIZATION.md), make sure to clean up the temporary files that generated in that step._**

To run the app locally, you need to run migration first (can skip if you are from `Database Initialization` steps)

```bash
DATABASE_URL=postgres://user:password@ip:5432/dbname go run cmd/migrate/main.go
```

Then you can start the backend:

```bash
DATABASE_URL=postgres://user:password@ip:5432/dbname go run cmd/server/*.go
# http start at :3980
```

Then, frontend:

```bash
VITE_SERVER_URL=http://localhost:3980 yarn start
#    ┌──────────────────────────────────────┐
#    │                                          │
#    │   Serving!                               │
#    │                                          │
#    │   - Local:    http://localhost:8080      │
#    │   - Network:  http://10.XX.XXX.XX:8080   │
#    │                                          │
#    │   Copied local address to clipboard!     │
#    │                                          │
#    └──────────────────────────────────────┘
```

Or:

```bash
VITE_SERVER_URL=http://localhost:3980 yarn start:dev
# Open: http://localhost:5173/
```

### Dockering 🐳

#### Using Dockerfile

**_NOTE: this steps below is assume that you are already have a running PostgreSQL instance._**

Run the migration (skip if already done):

```bash
DATABASE_URL=postgres://user:password@ip:5432/dbname go run cmd/migrate/main.go
```

Build docker image for backend:

```bash
docker build \
    -t silver-bassoon/backend \
    -f Dockerfile.backend \
    --no-cache .
```

Run backend container from created image:

```bash
docker run -d \
    --name packform-be \
    -e DATABASE_URL='postgres://user:password@ip:5432/silver_bassoon?sslmode=disable' \
    -p 3980:3980 \
    silver-bassoon/backend
```

Build docker image for frontend:

```bash
docker build \
    -t silver-bassoon/frontend \
    -f Dockerfile.frontend \
    --build-arg server_url=http://localhost:3980 \
    --build-arg use_browser_tz=false \
    --no-cache .
```

Run frontend container from created image:

```bash
docker run -d \
    --name packform-fe \
    -p 8080:80 \
    silver-bassoon/frontend
```

#### Using docker-compose (Recommended)

By using `docker compose`, you are not required to have PostgreSQL running, and you are not required to run the migration as well. Everything will be done automatically, encapsulated.

```bash
docker compose up -d
```

## Testing 🧪

### Backend Unit Testing

```bash
go test -v ./...
```

### Frontend Unit Testing

```bash
yarn test
```
