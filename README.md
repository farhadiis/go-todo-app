
# Go Todo App

An implementation of Todo app with Clean Architecture and GO language

## API Reference

#### You can this apis to work with app

```http
  POST /api/v1/todo - create todo
  GET /api/v1/todo - get all todo
  GET /api/v1/todo/${id} - get todo by id 
  DELETE /api/v1/todo/${id} delete todo by id
```

## Deployment

To deploy this project clone it and run

```bash
  cd go-todo-app
  export MONGODB_URI='<your connection string>'
  go run .
```

## Running Tests

To run tests, run the following command

```bash
  go test ./tests -v
```