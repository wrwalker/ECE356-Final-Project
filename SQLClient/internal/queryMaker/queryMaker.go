package queryMaker

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/internal/dbConnector"
	_ "github.com/go-sql-driver/mysql"
	sql "github.com/jmoiron/sqlx"
	"log"
)

// STRUCTURES:
type VotesByState struct {
	State      string `db:"state"`
	TotalVotes int    `db:"totalVotes"`
	Level      string `db:"level"`
}

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

func DeserializeRows(r *sql.Rows) error {
	for r.Next() {
		var votesByState VotesByState

		err := r.StructScan(&votesByState)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the votesByState's Name attribute
		log.Printf(votesByState.State)
		log.Printf("%v", votesByState.TotalVotes)
		log.Printf("%v", votesByState.Level)
		log.Println()
	}
	return nil
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
