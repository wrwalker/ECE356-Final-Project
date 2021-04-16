package queryMaker

import (
	"errors"
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/dbConnector"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/utils"
	_ "github.com/go-sql-driver/mysql"
	sql "github.com/jmoiron/sqlx"
	"log"
	"math"
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

func getStringForGetWinner(state string, county string, candidate string) string {
	qString := fmt.Sprintf("select count(*) from VotesByCountyCandidate where won = true and candidate = %q", candidate)
	if county != "" {
		qString = fmt.Sprintf("%s and county=%q", qString, county)
	}
	if state != "" {
		qString = fmt.Sprintf("%s and state=%q", qString, state)
	}
	return qString
}

func (q *QueryMaker) GetWinner(state string, county string, maker *QueryMaker) (string, string, error) {
	if state != "" && county != ""{
		query := getStringForGetWinner(state, county, "Donald Trump")
		rows, colNames, err := q.DoRawQuery(query)
		if err != nil {
			return "", query, err
		}
		if len(rows) < 1 || rows[0][colNames[0]] == nil {
			return "", query, errors.New("could not find any matches")
		}
		bytes := rows[0][colNames[0]].([]byte)
		trump, _ := strconv.Atoi(string(bytes))
		fmt.Printf("Ran: %s\n", query)

		query = getStringForGetWinner(state, county, "Joe Biden")
		rows, colNames, err = q.DoRawQuery(query)
		if err != nil {
			return "", query, err
		}
		if len(rows) < 1 || rows[0][colNames[0]] == nil {
			return "", query, errors.New("could not find any matches")
		}
		bytes = rows[0][colNames[0]].([]byte)
		biden, _ := strconv.Atoi(string(bytes)) // hack to convert from byteslice ([]uint8) to int

		if trump > biden {
			return "Donald Trump", query, nil
		} else if trump < biden {
			return "Joe Biden", query, nil
		} else {
			return "Tie?", query, errors.New("could not find any matches")
		}
	} else if state != "" && county == "" {
		x := []string{state}
		trump, _, _ := maker.GetVotesForCandidate("Donald Trump", "", x, false)

		biden, _, _ := maker.GetVotesForCandidate("Joe Biden", "", x, false)

		if trump > biden {
			return "Donald Trump", "", nil
		} else if trump < biden {
			return "Joe Biden", "", nil
		} else {
			return "Tie?", "", errors.New("could not find any matches")
		}
	} else {
		return "Joe Biden", "", nil
	}
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

func getStringForGetNumberOfTweets(state string, county string) string {
	qString := fmt.Sprintf("select count(ElectionTweets.tweetID) from ElectionTweets")

	if county != "" || state != "" {
		qString = fmt.Sprintf("%s join Location on Location.tweetID = ElectionTweets.tweetID where", qString)

		if state != "" {
			qString = fmt.Sprintf("%s Location.state_name=%q", qString, state)
		}
		if county != "" && state != "" {
			qString = fmt.Sprintf("%s and", qString)
		}
		if county != "" {
			qString = fmt.Sprintf("%s Location.county_name=%q", qString, county)
		}
	}
	return qString
}

func (q *QueryMaker) GetNumberOfTweets(state string, county string) (int, string, error) {
	query := getStringForGetNumberOfTweets(state, county)
	rows, colNames, err := q.DoRawQuery(query)
	if err != nil {
		return -1, query, err
	}
	if len(rows) < 1 || rows[0][colNames[0]] == nil {
		return -1, query, errors.New("could not find any matches")
	}
	bytes := rows[0][colNames[0]].([]byte)
	byteToInt, _ := strconv.Atoi(string(bytes)) // hack to convert from byteslice ([]uint8) to int
	return byteToInt, query, nil
}

func getStringForGetTweetPositiveSentimentCount(trumpOrBiden string, state string, county string) string {
	qString := fmt.Sprintf("select count(ElectionTweets.tweetID) from ElectionTweets join Sentiment on ElectionTweets.tweetID=Sentiment.tweetID join Location on ElectionTweets.tweetID=Location.tweetID")
	qString = fmt.Sprintf("%s where ElectionTweets.trumpOrBiden=%q and sentimentScore=true", qString, trumpOrBiden)

	// select count(*) from ElectionTweets join Sentiment on ElectionTweets.tweetID=Sentiment.tweetID where ElectionTweets.trumpOrBiden='T' and sentimentScore=true

	if county != "" || state != "" {
		if state != "" {
			qString = fmt.Sprintf("%s and state_name=%q", qString, state)
		}
		if county != "" {
			qString = fmt.Sprintf("%s and county_name=%q", qString, county)
		}
	}
	return qString
}

func getStringForGetTweetTotalSentimentCount(trumpOrBiden string, state string, county string) string {
	qString := fmt.Sprintf("select count(ElectionTweets.tweetID) from ElectionTweets join Sentiment on ElectionTweets.tweetID=Sentiment.tweetID join Location on ElectionTweets.tweetID=Location.tweetID")
	qString = fmt.Sprintf("%s where ElectionTweets.trumpOrBiden=%q", qString, trumpOrBiden)

	// select count(*) from ElectionTweets join Sentiment on ElectionTweets.tweetID=Sentiment.tweetID where ElectionTweets.trumpOrBiden='T' and sentimentScore=true

	if county != "" || state != "" {
		if state != "" {
			qString = fmt.Sprintf("%s and state_name=%q", qString, state)
		}
		if county != "" {
			qString = fmt.Sprintf("%s and county_name=%q", qString, county)
		}
	}
	return qString
}

func (q *QueryMaker) GetTweetSentiment(candidate string, state string, county string) (float64, string, error) {
	if candidate != "Donald Trump" && candidate != "Joe Biden" {
		return math.NaN(), "", errors.New("candidate must be \"Donald Trump\" or \"Joe Biden\"")
	}
	trumpOrBiden := ""
	if candidate == "Donald Trump" {
		trumpOrBiden = "T"
	} else {
		trumpOrBiden = "B"
	}

	query := getStringForGetTweetPositiveSentimentCount(trumpOrBiden, state, county)
	rows, colNames, err := q.DoRawQuery(query)
	if err != nil {
		return -1, query, err
	}
	if len(rows) < 1 || rows[0][colNames[0]] == nil {
		return -1, query, errors.New("could not find any matches")
	}
	bytes := rows[0][colNames[0]].([]byte)
	positiveSentimentCount, _ := strconv.Atoi(string(bytes)) // hack to convert from byteslice ([]uint8) to int
	fmt.Printf("Ran: %s\n", query)


	query = getStringForGetTweetTotalSentimentCount(trumpOrBiden, state, county)
	rows, colNames, err = q.DoRawQuery(query)
	if err != nil {
		return -1, query, err
	}
	if len(rows) < 1 || rows[0][colNames[0]] == nil {
		return -1, query, errors.New("could not find any matches")
	}
	bytes = rows[0][colNames[0]].([]byte)
	totalCount, _ := strconv.Atoi(string(bytes)) // hack to convert from byteslice ([]uint8) to int
	return float64(positiveSentimentCount)/float64(totalCount) * 100.0, query, nil
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
