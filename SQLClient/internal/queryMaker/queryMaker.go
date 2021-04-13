package queryMaker

import (
	"errors"
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/internal/dbConnector"
	"github.com/ECE356-Final-Project/SQLClient/internal/utils"
	_ "github.com/go-sql-driver/mysql"
	sql "github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"strings"
)

// QueryMaker struct{Db dbInterface, sql }
type QueryMaker struct {
	Db dbConnector.DBConnector
}

func (q *QueryMaker) connectToDB() (dbConnector.DBConnector, error) {
	//fmt.Println("------------ 2020 Elections Sentiment Analysis Results Database ------------")
	//fmt.Println("Connecting to DB...")

	connString := "root:root@tcp(127.0.0.1:3307)/Election"
	return sql.Open("mysql", connString)
}

func (q *QueryMaker) doQuery(input string) (*sql.Rows, error) {
	results, err := q.Db.Queryx(input)
	if err != nil {
		log.Printf("error doing query %q: %s", input, err.Error())
		return nil, err
	}
	return results, nil
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

func getStringForGetVotesForCandidate(candidate, county string, states []string) string {
	qString := fmt.Sprintf("select sum(votes) from VotesByCountyCandidate where candidate = %q", candidate)

	if county != "" || len(states) > 0 {
		if county != "" {
			qString = fmt.Sprintf("%s and county=%q", qString, county)
		}
		if len(states) > 0 {
			qStates := ""
			for _, state := range states {
				qStates = fmt.Sprintf("%s state=%q or", qStates, state)
			}
			qStates = strings.Trim(qStates, " or")
			qString = fmt.Sprintf("%s and (%s)", qString, qStates)
		}
	}
	return qString
}

func (q *QueryMaker) GetVotesForCandidate(candidate, county string, states []string) (int, string, error) {
	query := getStringForGetVotesForCandidate(candidate, county, states)
	rows, colNames, err := q.DoRawQuery(query)
	if err != nil {
		return 0, "", err
	}
	if len(rows) < 1 || rows[0][colNames[0]] == nil {
		return 0, "", errors.New("could not find any matches")
	}
	bytes := rows[0][colNames[0]].([]byte)
	byteToInt, _ := strconv.Atoi(string(bytes)) // hack to convert from byteslice ([]uint8) to int
	return byteToInt, query, nil
}

func (q *QueryMaker) DoRawQuery(input string) ([]map[string]interface{}, []string, error) {
	res, err := q.doQuery(input)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not execute raw query: %s", err.Error())
	}
	return utils.DeserializeRowsToMappedInterface(res)
}

func (q *QueryMaker) DoRawQueryWithPrint(input string) error {
	rows, colNames, err := q.DoRawQuery(input)
	if err != nil {
		return err
	}
	utils.PrintMap(rows, colNames)
	return nil
}
