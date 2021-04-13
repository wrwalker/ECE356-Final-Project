package dbConnector

import sql "github.com/jmoiron/sqlx"

// DBConnector interface
type DBConnector interface {
	Close() error
	Queryx(query string, args ...interface{}) (*sql.Rows, error)
}
