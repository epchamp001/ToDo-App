# Todo API

Это простое RESTful API для управления задачами (Todo) на основе фреймворка [Gin](https://github.com/gin-gonic/gin) и ORM [GORM](https://gorm.io/), с использованием базы данных SQLite.

## Функциональность

- Создание новой задачи
- Получение списка всех задач
- Получение задачи по ID
- Обновление задачи по ID
- Удаление задачи по ID

## Структура задачи (Todo)

Задача включает следующие поля:

- `ID` (автоинкрементный идентификатор)
- `Title` (название задачи)
- `Description` (описание задачи)
- `CreatedAt` (время создания)
- `UpdatedAt` (время обновления)

## Установка и запуск

### Предварительные требования

- Установленный [Go](https://golang.org/dl/)

### Шаги установки

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/ваше_имя_пользователя/todo-api.git
   cd todo-api
   ```
2. Установите зависимости (Gin и GORM):
    ```bash
    go mod tidy
    ```
3. Запустите сервер:
    ```bash
    go run main.go
    ```


Сервер запустится на порту `8080`. Чтобы изменить порт, отредактируйте строку в `main.go`:
   
 ```bash
    router.Run(":8080")
```

## Примеры использования API
Используйте `curl` или любой инструмент для тестирования API (например, Postman) для взаимодействия с сервером.

### Создание новой задачи
**Запрос:**
```bash
  curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "My First Todo", "description": "This is a test todo"}'
```
**Ответ:**
```json
{
  "ID": 1,
  "Title": "My First Todo",
  "Description": "This is a test todo",
  "CreatedAt": "2023-10-01T12:00:00Z",
  "UpdatedAt": "2023-10-01T12:00:00Z"
}
```
### Получение списка всех задач
**Запрос:**
```bash
  curl -X GET http://localhost:8080/todos
```
**Ответ:**
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
### Получение задачи по ID
**Запрос:**
```bash
  curl -X GET http://localhost:8080/todos/1
```
**Ответ:**
```json
{
  "ID": 1,
  "Title": "My First Todo",
  "Description": "This is a test todo",
  "CreatedAt": "2023-10-01T12:00:00Z",
  "UpdatedAt": "2023-10-01T12:00:00Z"
}
```
### Обновление задачи по ID
**Запрос:**
```bash
  curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated Title", "description": "Updated description"}'
```
**Ответ:**
```json
{
  "ID": 1,
  "Title": "Updated Title",
  "Description": "Updated description",
  "CreatedAt": "2023-10-01T12:00:00Z",
  "UpdatedAt": "2023-10-01T12:30:00Z"
}
```
### Удаление задачи по ID
**Запрос:**
```bash
 curl -X DELETE http://localhost:8080/todos/1
```
**Ответ:**
```json
{
  "message": "Todo with ID 1 deleted"
}
```

## Зависимости

- [Gin Web Framework](https://github.com/gin-gonic/gin) — HTTP веб-фреймворк для Go.
- [GORM ORM Library](https://gorm.io/) — ORM для работы с базами данных в Go.
- SQLite — легковесная реляционная база данных.


 
