# Todo API

This is a simple RESTful API for managing tasks (Todo) built with the [Gin](https://github.com/gin-gonic/gin) framework and the ORM [GORM](https://gorm.io/), using SQLite as the database.

## Features

- Create a new task
- Retrieve a list of all tasks
- Retrieve a task by ID
- Update a task by ID
- Delete a task by ID

## Task (Todo) Structure

A task includes the following fields:

- `ID` (auto-increment identifier)
- `Title` (task title)
- `Description` (task description)
- `CreatedAt` (creation timestamp)
- `UpdatedAt` (update timestamp)

## Installation and Launch

### Prerequisites

- Installed [Go](https://golang.org/dl/)

### Installation Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/ваше_имя_пользователя/todo-api.git
   cd todo-api
   ```
2. Install dependencies (Gin and GORM):
    ```bash
    go mod tidy
    ```
3. Start the server:
    ```bash
    go run main.go
    ```


The server will start on port '8080'. To change the port, edit the line in `main.go`:
 ```bash
    router.Run(":8080")
```

## API usage Examples
Use `curl` or any API testing tool (for example, Postman) to interact with the server.

### Creating a new task
**Request:**

```bash
  curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "My First Todo", "description": "This is a test todo"}'
```
**Response:**
```json
{
  "ID": 1,
  "Title": "My First Todo",
  "Description": "This is a test todo",
  "CreatedAt": "2023-10-01T12:00:00Z",
  "UpdatedAt": "2023-10-01T12:00:00Z"
}
```
### Getting a list of all tasks
**Request:**
```bash
  curl -X GET http://localhost:8080/todos
```
**Response:**
```json
[
  {
    "ID": 1,
    "Title": "My First Todo",
    "Description": "This is a test todo",
    "CreatedAt": "2023-10-01T12:00:00Z",
    "UpdatedAt": "2023-10-01T12:00:00Z"
  }
]
```
### Getting a task by ID
**Request:**
```bash
  curl -X GET http://localhost:8080/todos/1
```
**Response:**
```json
{
  "ID": 1,
  "Title": "My First Todo",
  "Description": "This is a test todo",
  "CreatedAt": "2023-10-01T12:00:00Z",
  "UpdatedAt": "2023-10-01T12:00:00Z"
}
```
### Updating the task by ID
**Request:**
```bash
  curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated Title", "description": "Updated description"}'
```
**Response:**
```json
{
  "ID": 1,
  "Title": "Updated Title",
  "Description": "Updated description",
  "CreatedAt": "2023-10-01T12:00:00Z",
  "UpdatedAt": "2023-10-01T12:30:00Z"
}
```
### Deleting a task by ID
**Request:**
```bash
 curl -X DELETE http://localhost:8080/todos/1
```
**Response:**
```json
{
  "message": "Todo with ID 1 deleted"
}
```

## Зависимости

- [Gin Web Framework](https://github.com/gin-gonic/gin ) is an HTTP web framework for Go.
- [GORM ORM Library](https://gorm.io/) — ORM for working with databases in Go.
- SQLite is a lightweight relational database.


 
