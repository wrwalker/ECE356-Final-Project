package cmd

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/queryMaker"
	"github.com/spf13/cobra"
)

const rawQuery = "RawQuery"

var inputQuery string

// example usage: RawQueryCmd -i 'SELECT * FROM VotesByCountyCandidate limit 10'
var RawQueryCmd = &cobra.Command{
	Use:   rawQuery,
	Short: "Run a Raw SQL Query",
	RunE: func(cmd *cobra.Command, args []string) error {
		qm := queryMaker.NewQueryMaker()
		defer qm.Db.Close()

		if Verbose {
			fmt.Printf("Ran: %s\n", inputQuery)
		}

		return qm.DoRawQueryWithPrint(inputQuery)
	},
}

func init() {
	inputQueryFlag := "inputQuery"
	RawQueryCmd.Flags().StringVarP(&inputQuery, inputQueryFlag, "i", "", "raw sql query")
	RawQueryCmd.MarkFlagRequired(inputQueryFlag)

	RootCMD.AddCommand(RawQueryCmd)
}
