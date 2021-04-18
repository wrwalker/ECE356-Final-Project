package main

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/queryMaker"
	"log"
	"os"
	"strconv"
)

const testing = false // testing flag to do a limit on query return size
const likesMultiplyer = float64(0.1)
const retweetMultiplyer = float64(0.5)

const dir = "../dataMiningDataSets/"

var candidates = []string{"biden", "trump"}

type TweetData struct {
	county                    string
	state                     string
	weightedSentimentOfTweets float64
	won                       string
	percentOfTweets           float64
}

func main() {
	qm := queryMaker.NewQueryMaker()
	defer qm.Db.Close()
	//countiesAndStates, _, err := qm.DoRawQuery("select distinct county ,state from VotesByCountyCandidate")
	//if err != nil {
	//	log.Fatal(err)
	//}
	// map of state-county -> tweetData
	countiesAndStatesList := []map[string]*TweetData{}
	countiesAndStatesList = append(countiesAndStatesList, map[string]*TweetData{})
	countiesAndStatesList = append(countiesAndStatesList, map[string]*TweetData{})

	// set up list of states and counties
	//for _,row := range countiesAndStates {
	//	state := row["state"]
	//	county := row["county"]
	//	key := fmt.Sprintf("%s-%s", state, county)
	//	countiesAndStatesList[0][key] = &TweetData{
	//		state: fmt.Sprintf("%s", state),
	//		county: fmt.Sprintf("%s", county),
	//	}
	//	countiesAndStatesList[1][key] = &TweetData{
	//		state: fmt.Sprintf("%s", state),
	//		county: fmt.Sprintf("%s", county),
	//	}
	//}

	query := "select ElectionTweets.trumpOrBiden,sentimentScore,county_name, state_name, likes, retweetCount from Sentiment join Location on Sentiment.tweetID=Location.TweetID join ElectionTweets on ElectionTweets.TweetID=Location.TweetID"

	if testing {
		query = fmt.Sprintf("%s %s", query, "limit 10000")
	}

	tweets, _, err := qm.DoRawQuery(query)
	if err != nil {
		log.Fatal(err)
	}
	totalTweetsForCounty := map[string]int{}
	for _, row := range tweets {
		trumpOrBiden := fmt.Sprintf("%s", row["trumpOrBiden"])
		sentimentScore, _ := strconv.Atoi(string(row["sentimentScore"].([]byte)))
		county_name := fmt.Sprintf("%s", row["county_name"])
		state_name := fmt.Sprintf("%s", row["state_name"])
		likes, _ := strconv.Atoi(string(row["likes"].([]byte)))
		retweetCount, _ := strconv.Atoi(string(row["retweetCount"].([]byte)))

		key := fmt.Sprintf("%s-%s", state_name, county_name)
		if sentimentScore == 0 { // adjust sentiment to have neutral centered at 0
			sentimentScore = -1
		}
		score := float64(float64(sentimentScore) * (1 + float64(likes)*likesMultiplyer + float64(retweetCount)*retweetMultiplyer))
		if _, ok := countiesAndStatesList[0][key]; !ok {
			totalTweetsForCounty[key] = 0
			countiesAndStatesList[0][key] = &TweetData{
				county: county_name,
				state:  state_name,
			}
			countiesAndStatesList[1][key] = &TweetData{
				county: county_name,
				state:  state_name,
			}

			//getWinner
			winner, _, err := qm.GetWinner(state_name, county_name, qm)
			if err == nil {
				if winner == "Joe Biden" {
					countiesAndStatesList[0][key].won = "y"
					countiesAndStatesList[1][key].won = "n"
				} else {
					countiesAndStatesList[0][key].won = "n"
					countiesAndStatesList[1][key].won = "y"
				}
			} else { // try again with county
				swinner, _, err := qm.GetWinner(state_name, fmt.Sprintf("%s county", county_name), qm)
				if err == nil {
					if swinner == "Joe Biden" {
						countiesAndStatesList[0][key].won = "y"
						countiesAndStatesList[1][key].won = "n"
					} else {
						countiesAndStatesList[0][key].won = "n"
						countiesAndStatesList[1][key].won = "y"
					}
				} else { // default to trump win
					countiesAndStatesList[0][key].won = "n"
					countiesAndStatesList[1][key].won = "y"
				}
			}

		}
		totalTweetsForCounty[key] += 1
		// accumulate weightedSentiment
		if trumpOrBiden == "B" {
			countiesAndStatesList[0][key].weightedSentimentOfTweets += score
			countiesAndStatesList[0][key].percentOfTweets += 1
		} else { // trump
			countiesAndStatesList[1][key].weightedSentimentOfTweets += score
			countiesAndStatesList[1][key].percentOfTweets += 1
		}
	}

	// get percentage of Tweets
	for k, _ := range countiesAndStatesList[0] {
		countiesAndStatesList[0][k].percentOfTweets = 1.0 * countiesAndStatesList[0][k].percentOfTweets / float64(totalTweetsForCounty[k]) * 100.0
		countiesAndStatesList[1][k].percentOfTweets = 1.0 * countiesAndStatesList[0][k].percentOfTweets / float64(totalTweetsForCounty[k]) * 100.0
	}

	// print to csv
	for i, candidate := range candidates {
		path := dir + fmt.Sprintf("%sprocessedTweetData.csv", candidate)
		err := os.Remove(path)
		if err != nil {
			log.Fatal("couldn't delete file")
			//return
		}
		f, err := os.Create(path)
		if err != nil {
			log.Fatal("couldn't write file")
			return
		}
		defer f.Close()

		//county string
		//state string
		//weightedSentimentOfTweets float64
		//won string
		//percentOfTweets float64
		f.WriteString(fmt.Sprintf("county,state,weightedSentimentOfTweets,won,percentOfTweets\n"))
		for _, tweet := range countiesAndStatesList[i] {
			f.WriteString(fmt.Sprintf("%s,%s,%.4f,%s,%.2f\n", tweet.county, tweet.state, tweet.weightedSentimentOfTweets, tweet.won, tweet.percentOfTweets))
		}

	}

	fmt.Sprintf("done")
}
