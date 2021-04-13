package queryMaker

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/internal/dbConnector"
	_ "github.com/go-sql-driver/mysql"
	sql "github.com/jmoiron/sqlx"
	"log"
)

// QueryMaker struct{Db dbInterface, sql }
type QueryMaker struct {
	Db dbConnector.DBConnector
}

func (q *QueryMaker) connectToDB() (dbConnector.DBConnector, error) {
	fmt.Println("------------ 2020 Elections Sentiment Analysis Results Database ------------")
	fmt.Println("Connecting to DB...")

	connString := "root:root@tcp(127.0.0.1:3307)/Election"
	return sql.Open("mysql", connString)
}

func (q *QueryMaker) DoQuery(input string) (*sql.Rows, error) {
	results, err := q.Db.Queryx(input)
	if err != nil {
		log.Printf("error doing query %q: %s", input, err.Error())
		return nil, err
	}
	return results, nil
}

func DeserializeRows(r *sql.Rows) ([]map[string]interface{}, []string, error) {
	var allRows []map[string]interface{}
	var colHeaders []string
	for r.Next() {
		results := make(map[string]interface{})

		err := r.MapScan(results) // this can cause errors with non-fully qualified names
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		colHeaders, err = r.Columns()
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		allRows = append(allRows, results)
	}
	return allRows, colHeaders, nil
}

func NewQueryMaker(dbs ...dbConnector.DBConnector) *QueryMaker {
	qm := &QueryMaker{}

	if len(dbs) > 0 { // use optional dbConnector for testing
		qm.Db = dbs[0]
	} else {
		db, err := qm.connectToDB()
		if err != nil {
			log.Fatalf("Could not connect to DB: %s", err.Error())
		}
		qm.Db = db
	}
	return qm
}
