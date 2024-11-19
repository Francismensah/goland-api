# Todo List API

A simple RESTful API for managing todo items built with Go and PostgreSQL.

## Features

- Create new todo items
- List all todo items
- Search todo items
- RESTful API design
- PostgreSQL database storage
- Unit tests with mocking

## Prerequisites

- Go 1.22 or higher
- PostgreSQL 14 or higher
- Docker (optional)

## Installation

1. Clone the repository:
    ```bash
    git https://github.com/Francismensah/goland-api.git
    cd my-first-api
    ```

2. Install dependencies:
    ```bash
    go mod download
    ```

3. Set up the database:
    ```sql
    CREATE TABLE IF NOT EXISTS todo_items (
        id SERIAL PRIMARY KEY,
        task TEXT NOT NULL,
        status TEXT NOT NULL
    );
    ```

4. Configure the database connection in `main.go` or through environment variables.

## Running the Application

1. Start the server:
    ```bash
    go run main.go
    ```

The server will start on `localhost:8080`

## API Endpoints

### GET /todo
Returns all todo items.

Response:
```json
[
    {
        "task": "Buy groceries",
        "status": "TO_BE_STARTED"
    }
]
```

### POST /todo
Creates a new todo item.

Request:
```json
{
    "item": "Buy groceries"
}
```

### GET /search?q=query
Search for todo items.

Example: `GET /search?q=groceries`

Response:
```json
[
    "Buy groceries"
]
```

## Testing

Run the tests:
```bash
go test ./...
```

## Project Structure
```
.my-first-api
├── internal/
│   ├── db/
│   │   └── db.go
│   ├── todo/
│   │   ├── todo.go
│   │   └── todo_test.go
│   └── transport/
│       └── http.go
├── docker-compose.yml
├── go.mod
├── main.go
├── my-first-api.http
└── README.md
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is not under any license. Feel to clone and use.

## Acknowledgments

- Thanks to the Go community for the amazing tools and libraries
- [jackc/pgx](https://github.com/jackc/pgx) for PostgreSQL driver