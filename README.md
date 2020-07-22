# ToDo API

Simple "To Do" application API written in Golang

## How to run this project
```docker-compose up``` will expose the server on localhost:8000

## Endpoints
- `GET, POST /tasks`
- `GET /tasks/{task.ID}`
- `PUT /tasks/{task.ID}`

## Sample POST Request Body
```json
{
    "title": "Task 1"
}
```

## Sample GET Response Body
```json
{
    "id": "1595458521160191726",
    "title": "Task 1",
    "createdAt": "2020-07-22T22:55:21.160196Z",
    "status": "New"
}
```

## Sample PUT Request Body
The full task struct is currently required for PUT
```json
{
    "id": "1595458521160191726",
    "title": "Task 1",
    "createdAt": "2020-07-22T22:55:21.160196Z",
    "status": "InProgress"
}
```

## TODO
- [ ] Deployments
    - [x] dockerize app
    - [x] dockerize db
    - [ ] dockerize nginx
    - [ ] deploy to server
- [ ] Bulk Ops
- [ ] Tests