# Database Initialization

## Whats going on here?

- [Short Mode](#initialize-the-database-short-mode)
- [Advance Mode](#initialize-the-database-short-mode)

## Initialize the Database (Short Mode)

To initiate the Database from `.csv` using script (Golang), run this:

```bash
DATABASE_URL=postgres://user:password@ip:5432/dbname ./initdb.sh
```

## Initialize the Database (Advance Mode - With AI - Recommended)

**_NOTE: If you've done the short mode, you need to cleanup your db first, and then run this:_**

```bash
./cleandb.sh
```

### Start!

The provided `.csv` files are located in the `cmd/gen/files` directory. Additionally, all generated `.sql` files can be found in the `cmd/migrate/migrations` directory.

```bash
├── cmd
│   ├── gen
│   │   ├── files
│   │   │   ├── customer_companies.csv
│   │   │   ├── customers.csv
│   │   │   ├── deliveries.csv
│   │   │   ├── order_items.csv
│   │   │   └── orders.csv
│   │   └── main.go
│   └── migrate
│       ├── main.go
│       └── migrations
│           └── 00_initial.up.sql
```

### Generate `.sql` from `.csv`

To generate the migration (`.sql` file), run:

```bash
go run cmd/gen/main.go
```

At this stage, you will be prompted with the message: `migration file already exists`. This is expected since, by default, the program checks for the existence of the default migration file first.

To create a new migration file from the provided `.csv` files, use:

```bash
go run cmd/gen/main.go -new
```

Once the process done, you will got new file under `cmd/migrate/migrations` directory.

```
├── cmd
│   ├── gen
│   │   ...
│   └── migrate
│       ├── main.go
│       └── migrations
│           ├── 00_initial.up.sql
│           └── 2024012XXXXXXX_initial.up.sql (NEW)
```

If you inspect the result of the generated `.sql` file, you'll notice that all columns are using `VARCHAR(255)` as the data type. This is intentional for simplicity. However, if you prefer to generate **well-defined** and **appropriate** data types, you can do so using OpenAI tools in the next step.

### Generate Well-Defined `.sql`

To enable this feature, first, you need to have an `OPEN AI KEY`. You can generate your own `KEY` from [here](https://platform.openai.com/account/api-keys).

Once you have the `KEY`, run this command:

```bash
OPEN_AI_KEY=your-key go run cmd/gen/main.go -new
```

**_Note: This process may take a longer time as it will generate the provided seed data that will be inserted._**

### Run Migrations

After migration file (`.sql`) created, next you need to run the migration. To run migrations, run:

```bash
DATABASE_URL=postgres://user:password@ip:5432/dbname go run cmd/migrate/main.go
```

### Congrats!

Congratulation! Your database initialization has finished 🚀
