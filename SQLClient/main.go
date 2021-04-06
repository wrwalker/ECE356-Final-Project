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

//TODO: Make a database object struct that we can pass around and shit

//TODO: Make an interface with all our methods

//TODO: Make a connect method
func connectToDB() {

}

func main() {
	fmt.Println("------------ 2020 Elections Sentiment Analysis Results Database ------------")
	fmt.Println("Connecting to DB...")

	//TODO: Move all this shit out of here into a func
	connString := "root:root@tcp(127.0.0.1:3307)/Election"
	db, err := sql.Open("mysql", connString)
	defer db.Close()

	if err != nil {
		fmt.Println("Failed to connect to DB! Panik")
		panic(err.Error())
	}
	fmt.Println("Success")

	results, err := db.Query("SELECT * FROM VotesByState")
	if err != nil {
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
}
