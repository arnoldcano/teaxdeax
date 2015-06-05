package infrastructure

import (
	"database/sql"

	"github.com/arnoldcano/teaxdeax/interfaces"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteHandler struct {
	conn *sql.DB
}

func NewSqliteHandler(fileName string) *SqliteHandler {
	conn, _ := sql.Open("sqlite3", fileName)
	return &SqliteHandler{
		conn: conn,
	}
}

func (handler *SqliteHandler) Execute(query string) error {
	if _, err := handler.conn.Exec(query); err != nil {
		return err
	}
	return nil
}

func (handler *SqliteHandler) Query(query string) (interfaces.Rows, error) {
	rows, err := handler.conn.Query(query)
	if err != nil {
		return nil, err
	}
	return NewSqliteRows(rows), nil
}

type SqliteRows struct {
	rows *sql.Rows
}

func NewSqliteRows(rows *sql.Rows) *SqliteRows {
	return &SqliteRows{
		rows: rows,
	}
}

func (rows *SqliteRows) Scan(dest ...interface{}) error {
	return rows.rows.Scan(dest...)
}

func (rows *SqliteRows) Next() bool {
	return rows.rows.Next()
}

func (rows *SqliteRows) Close() error {
	return rows.rows.Close()
}
