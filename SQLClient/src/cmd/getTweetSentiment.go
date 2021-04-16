package cmd

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/queryMaker"
	"github.com/spf13/cobra"
)

const getTweetSentiment = "GetTweetSentiment"

var candidateToGetSentiment string
var stateToGetSentiment string
var countyToGetSentiment string

var getTweetSentimentCmd = &cobra.Command{
	Use:   getTweetSentiment,
	Short: "Determine % positive or negative tweets based on sentiment analysis results by county, state, or the whole USA",
	RunE: func(cmd *cobra.Command, args []string) error {
		qm := queryMaker.NewQueryMaker()
		defer qm.Db.Close()

		percentSentimentScore, attemptedQuery, err := qm.GetTweetSentiment(candidateToGetSentiment, stateToGetSentiment, countyToGetSentiment)
		if Verbose {
			fmt.Printf("Ran: %s\n", attemptedQuery)
		}
		if err != nil {
			return err
		}
		fmt.Printf("recieved %g%% positive sentiment towards candidate\n", percentSentimentScore)
		return nil
	},
}

func init() {

	candidateNameFlag := "candidateName"
	getTweetSentimentCmd.Flags().StringVarP(&candidateToGetSentiment, candidateNameFlag, "n", "", "candidate for sentiment analysis. Note only \"Donald Trump\" and \"Joe Biden\" have sentiment data")
	getTweetSentimentCmd.MarkFlagRequired(candidateNameFlag)

	stateFlag := "state"
	getTweetSentimentCmd.Flags().StringVarP(&stateToGetSentiment, stateFlag, "s", "", "state to get # of tweets for")

	countyFlag := "county"
	getTweetSentimentCmd.Flags().StringVarP(&countyToGetSentiment, countyFlag, "c", "", "county to get # of tweets for")

	RootCMD.AddCommand(getTweetSentimentCmd)
}
