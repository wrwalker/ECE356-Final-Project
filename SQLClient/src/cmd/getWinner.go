package cmd

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/queryMaker"
	"github.com/spf13/cobra"
)

const getWinner = "GetWinner"

var stateWhoWon string
var countyWhoWon string

var getWinnerCmd = &cobra.Command{
	Use:   getWinner,
	Short: "Get who won for a given county, state, or all of the USA",
	RunE: func(cmd *cobra.Command, args []string) error {
		qm := queryMaker.NewQueryMaker()
		defer qm.Db.Close()

		candidateWhoWon, attemptedQuery, err := qm.GetWinner(stateWhoWon, countyWhoWon, qm)
		if Verbose {
			fmt.Printf("Ran: %s\n", attemptedQuery)
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s won\n", candidateWhoWon)
		return nil
	},
}

func init() {

	stateFlag := "state"
	getWinnerCmd.Flags().StringVarP(&stateWhoWon, stateFlag, "s", "", "state to see who won")

	countyFlag := "county"
	getWinnerCmd.Flags().StringVarP(&countyWhoWon, countyFlag, "c", "", "county to see who won")

	RootCMD.AddCommand(getWinnerCmd)
}
