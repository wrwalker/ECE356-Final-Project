package dbConnector

import "database/sql"

// DBConnector interface
type DBConnector interface {
	Close() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
