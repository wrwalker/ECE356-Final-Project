package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// STRUCTURES:
type VotesByState struct {
	state      string `json:"state"`
	totalVotes int    `json:"totalVotes"`
	level      string `json:"totalVotes"`
}

// DBConnector interface
type DBConnector interface {
	Close() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// QueryMaker struct{db dbInterface, sql }
type QueryMaker struct {
	db DBConnector
}

func (q *QueryMaker) connectToDB() (DBConnector, error) {
	fmt.Println("------------ 2020 Elections Sentiment Analysis Results Database ------------")
	fmt.Println("Connecting to DB...")

	connString := "root:root@tcp(127.0.0.1:3307)/Election"
	return sql.Open("mysql", connString)
}

func (q *QueryMaker) DoQuery(input string) (*sql.Rows, error) {
	results, err := q.db.Query(input)
	if err != nil {
		log.Printf("error doing query %q: %s", input, err.Error())
		return nil, err
	}
	return results, nil
}

func DeserializeRows(r *sql.Rows) error {
	for r.Next() {
		var votesByState VotesByState
		// for each row, scan the result into our votesByState composite object
		err := r.Scan(&votesByState.state, &votesByState.totalVotes, &votesByState.level)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the votesByState's Name attribute
		log.Printf(votesByState.state)
		log.Printf("%v", votesByState.totalVotes)
		log.Printf(votesByState.level)
		log.Println()
	}
	return nil
}

func NewQueryMaker(dbs ...DBConnector) *QueryMaker {
	qm := &QueryMaker{}

	if len(dbs) > 0 { // use optional dbConnector for testing
		qm.db = dbs[0]
	} else {
		db, err := qm.connectToDB()
		if err != nil {
			log.Fatalf("Could not connect to DB: %s", err.Error())
		}
		qm.db = db
	}
	return qm
}

func main() {
	qm := NewQueryMaker()
	defer qm.db.Close()

	res, _ := qm.DoQuery("SELECT * FROM VotesByState")
	DeserializeRows(res)

}
