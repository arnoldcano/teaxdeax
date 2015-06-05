# teaxdeax
a restful todo app written in Go and based on the clean architecture

1. sqlite3 todos.sqlite
2. bootstrap with setup.sql
3. FILE=todos.sqlite go run main.go

* GET /todos
* GET /todos/{id}
* POST /todos
* PUT /todos/{id}
* DELETE /todos/{id}