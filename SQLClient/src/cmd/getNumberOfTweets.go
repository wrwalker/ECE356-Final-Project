package cmd

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/queryMaker"
	"github.com/spf13/cobra"
)

const getNumberOfTweets = "GetNumberOfTweets"

var stateToGetNumberOfTweets string
var countyToGetNumberOfTweets string

var getNumberOfTweetsCmd = &cobra.Command{
	Use:   getNumberOfTweets,
	Short: "Get number of tweets by county, by state, or just in all of the USA",
	RunE: func(cmd *cobra.Command, args []string) error {
		qm := queryMaker.NewQueryMaker()
		defer qm.Db.Close()

		numberOfTweets, attemptedQuery, err := qm.GetNumberOfTweets(stateToGetNumberOfTweets, countyToGetNumberOfTweets)
		if Verbose {
			fmt.Printf("Ran: %s\n", attemptedQuery)
		}
		if err != nil {
			return err
		}
		fmt.Printf("recieved %d tweets\n", numberOfTweets)
		return nil
	},
}

func init() {

	stateFlag := "state"
	getNumberOfTweetsCmd.Flags().StringVarP(&stateToGetNumberOfTweets, stateFlag, "s", "", "state to get # of tweets for")

	countyFlag := "county"
	getNumberOfTweetsCmd.Flags().StringVarP(&countyToGetNumberOfTweets, countyFlag, "c", "", "county to get # of tweets for")

	RootCMD.AddCommand(getNumberOfTweetsCmd)
}
