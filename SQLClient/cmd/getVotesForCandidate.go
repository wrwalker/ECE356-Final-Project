package cmd

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/internal/queryMaker"
	"github.com/spf13/cobra"
)

const getVotesForCandidate = "GetVotesForCandidate"

var candidateName string
var states []string
var county string

var getVotesForCandidateCmd = &cobra.Command{
	Use:   getVotesForCandidate,
	Short: "Get the total votes for a candidate",
	RunE: func(cmd *cobra.Command, args []string) error {
		qm := queryMaker.NewQueryMaker()
		defer qm.Db.Close()

		numVotes, attemptedQuery, err := qm.GetVotesForCandidate(candidateName, county, states)
		if Verbose {
			fmt.Printf("Ran: %s\n", attemptedQuery)
		}
		if err != nil {
			return err
		}
		fmt.Printf("candidateName recieved %d votes\n", numVotes)
		return nil
	},
}

func init() {
	candidateNameFlag := "candidateName"
	getVotesForCandidateCmd.Flags().StringVarP(&candidateName, candidateNameFlag, "n", "", "name of candidate to tally Votes")
	getVotesForCandidateCmd.MarkFlagRequired(candidateNameFlag)

	stateFlag := "state"
	getVotesForCandidateCmd.Flags().StringSliceVarP(&states, stateFlag, "s", []string{}, "states to tally Votes")

	countyFlag := "county"
	getVotesForCandidateCmd.Flags().StringVarP(&candidateName, countyFlag, "c", "", "county to tally Votes")

	RootCMD.AddCommand(getVotesForCandidateCmd)
}
