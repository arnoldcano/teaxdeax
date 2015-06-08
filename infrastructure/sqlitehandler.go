package infrastructure

import (
	"database/sql"

	"github.com/arnoldcano/teaxdeax/interfaces"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteHandler struct {
	conn *sql.DB
}

func NewSqliteHandler(file string) *SqliteHandler {
	conn, _ := sql.Open("sqlite3", file)
	return &SqliteHandler{
		conn: conn,
	}
}

func (h *SqliteHandler) Execute(query string) error {
	if _, err := h.conn.Exec(query); err != nil {
		return err
	}
	return nil
}

func (h *SqliteHandler) Query(query string) (interfaces.Rows, error) {
	rows, err := h.conn.Query(query)
	if err != nil {
		return nil, err
	}
	return NewSqliteRows(rows), nil
}

type SqliteRows struct {
	rows *sql.Rows
}

func NewSqliteRows(r *sql.Rows) *SqliteRows {
	return &SqliteRows{
		rows: r,
	}
}

func (r *SqliteRows) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest...)
}

func (r *SqliteRows) Next() bool {
	return r.rows.Next()
}

func (r *SqliteRows) Close() error {
	return r.rows.Close()
}
