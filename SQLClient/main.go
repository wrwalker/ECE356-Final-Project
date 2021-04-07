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
	level      string   `json:"totalVotes"`
}

// DBConnector interface
type DBConnector interface {
	Close() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// CLI struct{db dbInterface, sql }
type CLI struct {
	db DBConnector
	//sql SQLDriver
}


func (c *CLI)connectToDB()(DBConnector, error) {
	fmt.Println("------------ 2020 Elections Sentiment Analysis Results Database ------------")
	fmt.Println("Connecting to DB...")

	connString := "root:root@tcp(127.0.0.1:3307)/Election"
	return sql.Open("mysql", connString)
}

func (c *CLI)DoQuery(input string)(*sql.Rows, error){
	results, err := c.db.Query(input)
	if err != nil {
		return nil, err
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var votesByState VotesByState
		// for each row, scan the result into our votesByState composite object
		err = results.Scan(&votesByState.state, &votesByState.totalVotes, &votesByState.level)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the votesByState's Name attribute
		log.Printf(votesByState.state)
		log.Printf("%v", votesByState.totalVotes)
		log.Printf(votesByState.level)
		log.Println()
	}
	return nil, nil
}

func NewCLI(dbs ...DBConnector) *CLI{
	cli := &CLI{}

	if len(dbs) > 0 {
		cli.db = dbs[0]
	} else {
		db, err := cli.connectToDB()
		if err != nil {
			fmt.Println("Failed to connect to DB! Panik")
			panic(err.Error())
		}
		cli.db = db
	}
	return cli
}

func main() {
	cli := NewCLI()
	defer cli.db.Close()


}
