package cmd

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/queryMaker"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/utils"
	"github.com/spf13/cobra"
)

const getCountyAnnotations = "GetCountyAnnotations"

var stateToGet string
var countyToGet string

var getCountyAnnotationsCmd = &cobra.Command{
	Use:   getCountyAnnotations,
	Short: "Get annotations you've added to a county, annotations you've added to counties in a state, or all annotations you've added",
	RunE: func(cmd *cobra.Command, args []string) error {
		qm := queryMaker.NewQueryMaker()
		defer qm.Db.Close()

		rows, colNames, attemptedQuery, err := qm.GetCountyAnnotations(stateToGet, countyToGet)
		if Verbose {
			fmt.Printf("Ran: %s\n", attemptedQuery)
		}
		if err != nil {
			return err
		}

		utils.PrintMap(rows, colNames)
		return nil
	},
}

func init() {

	countyFlag := "county"
	getCountyAnnotationsCmd.Flags().StringVarP(&countyToGet, countyFlag, "c", "", "county you want to get your annotation for. Optional")

	stateFlag := "state"
	getCountyAnnotationsCmd.Flags().StringVarP(&stateToGet, stateFlag, "s", "", "state you want to get your annotation for. Optional")

	RootCMD.AddCommand(getCountyAnnotationsCmd)
}
