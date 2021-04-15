package queryMaker

import (
	"errors"
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/dbConnector"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/utils"
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

func getStringForGetVotesForCandidate(candidate, county string, states []string, annotationsFlag bool) string {
	joinString := ""
	if annotationsFlag {
		joinString = " join County on VotesByCountyCandidate.state=County.state and VotesByCountyCandidate.county=County.county"
	}
	qString := fmt.Sprintf("select sum(votes) from VotesByCountyCandidate%s where candidate = %q", joinString, candidate)
	if annotationsFlag {
		joinString = " and County.annotations != \"\""
		qString = fmt.Sprintf("%s%s", qString, joinString)
	}

	if county != "" || len(states) > 0 {
		if county != "" {
			qString = fmt.Sprintf("%s and VotesByCountyCandidate.county=%q", qString, county)
		}
		if len(states) > 0 {
			qStates := ""
			for _, state := range states {
				qStates = fmt.Sprintf("%s VotesByCountyCandidate.state=%q or", qStates, state)
			}
			qStates = strings.Trim(qStates, " or")
			qString = fmt.Sprintf("%s and (%s)", qString, qStates)
		}
	}
	return qString
}

func (q *QueryMaker) GetVotesForCandidate(candidate, county string, states []string, annotationsFlag bool) (int, string, error) {
	query := getStringForGetVotesForCandidate(candidate, county, states, annotationsFlag)
	rows, colNames, err := q.DoRawQuery(query)
	if err != nil {
		return 0, query, err
	}
	if len(rows) < 1 || rows[0][colNames[0]] == nil {
		return 0, query, errors.New("could not find any matches")
	}
	bytes := rows[0][colNames[0]].([]byte)
	byteToInt, _ := strconv.Atoi(string(bytes)) // hack to convert from byteslice ([]uint8) to int
	return byteToInt, query, nil
}

func getStringForGetCountyAnnotations(state string, county string) string {
	qString := fmt.Sprintf("select state, county, annotations from County where annotations != \"\"")

	if state != ""{
		qString = fmt.Sprintf("%s and state=%q", qString, state)
	}
	if county != "" {
		qString = fmt.Sprintf("%s and county=%q", qString, county)
	}
	return qString
}

func (q *QueryMaker) GetCountyAnnotations(state string, county string) ([]map[string]interface{}, []string, string, error) {
	query := getStringForGetCountyAnnotations(state, county)
	rows, colNames, err := q.DoRawQuery(query)
	if err != nil {
		return nil, nil, query, err
	}
	if len(rows) < 1 || rows[0][colNames[0]] == nil {
		return nil, nil, query, errors.New("could not find any matches")
	}
	return rows, colNames, query, nil
}

func getStringForCheckCountyExists(state string, county string) string {
	return fmt.Sprintf("select state, county from County where state = %q and county = %q limit 1", state, county)
}

func (q *QueryMaker) CheckCountyExists(state string, county string) (bool, string, error) {
	query := getStringForCheckCountyExists(state, county)
	rows, colNames, err := q.DoRawQuery(query)
	if err != nil {
		return false, query, err
	}
	if len(rows) < 1 || rows[0][colNames[0]] == nil {
		return false, query, errors.New("could not find any matches")
	}

	return true, query, nil
}

func getStringForAddCountyAnnotation(state string, county string, annotation string) string {
	return fmt.Sprintf("update County set annotations = %q where state = %q and county = %q", annotation, state, county)
}

func (q *QueryMaker) AddCountyAnnotation(state string, county string, annotation string) (string, error) {
	query := getStringForAddCountyAnnotation(state, county, annotation)
	_, _, err := q.DoRawQuery(query)
	if err != nil {
		return query, err
	}

	return query, nil
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
