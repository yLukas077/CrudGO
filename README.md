
# Simple Go CRUD with MySQL Database

This project is a simple CRUD written in Go using the Gorilla Mux framework and MySQL for database operations.

## Requirements

- Go (version 1.16 or higher)
- Gorilla Mux
- MySQL Go driver (`github.com/go-sql-driver/mysql`)

## Installation

1. Clone this repository:
```bash
git clone <REPOSITORY_URL>
```

2. Navigate to the project folder:
```bash
cd <REPOSITORY_FOLDER_NAME>
```

3. Install the dependencies:
```bash
go get -u github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql
```

## Setup Database

Ensure you have a running MySQL instance. Set up the `DB_CONNECTION_STRING` environment variable with your database connection string.

Example:
```bash
export DB_CONNECTION_STRING=""user:password@/table?charset=utf8&parseTime=True"
```

## Usage

Run the server:
```bash
go run cmd/crudapp/main.go
```
or

```bash
go build -o crud-go
```
and execute



After executing, the server will be running on port 8080.

### Endpoints

1. **Create User**
- **Endpoint:** `/users`
- **Method:** POST

2. **Fetch All Users**
- **Endpoint:** `/users`
- **Method:** GET

3. **Fetch User by ID**
- **Endpoint:** `/users/{id}`
- **Method:** GET

4. **Update User**
- **Endpoint:** `/users/{id}`
- **Method:** PUT

5. **Delete User**
- **Endpoint:** `/users/{id}`
- **Method:** DELETE

## License

[MIT](https://choosealicense.com/licenses/mit/)
