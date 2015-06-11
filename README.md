# teaxdeax
A RESTful todo app written in Go and based on the clean architecture. Tested using dependency injection and the "testing" standard library package.

1. sqlite3 todos.sqlite < setup.sql
2. FILE=todos.sqlite go run main.go

* GET /todos
* GET /todos/{id}
* POST /todos
* PUT /todos/{id}
* DELETE /todos/{id}
