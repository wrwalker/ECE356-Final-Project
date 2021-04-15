package cmd

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/queryMaker"
	"github.com/spf13/cobra"
)

const addCountyAnnotation = "AddCountyAnnotation"

var state string
var countyToAnnotate string
var annotation string

var addCountyAnnotationCmd = &cobra.Command{
	Use:   addCountyAnnotation,
	Short: "Add an annotation string to a county. Can also be used to mark counties of interest",
	RunE: func(cmd *cobra.Command, args []string) error {
		qm := queryMaker.NewQueryMaker()
		defer qm.Db.Close()

		doesCountyExist, attemptedQuery, err := qm.CheckCountyExists(state, countyToAnnotate)
		if Verbose {
			fmt.Printf("Ran: {%s} to check if the county doesCountyExist\n", attemptedQuery)
			fmt.Printf("County exists: %t\n", doesCountyExist)
		}
		if err != nil {
			return err
		}
		if doesCountyExist == false {
			fmt.Println("This state / county combination does not exist in this database. Exiting...")
			return nil
		}

		attemptedQuery, err = qm.AddCountyAnnotation(state, countyToAnnotate, annotation)
		if Verbose {
			fmt.Printf("Ran: %s\n", attemptedQuery)
		}
		if err != nil {
			return err
		}
		fmt.Println("Updated County.")
		return nil
	},
}

func init() {

	stateFlag := "state"
	addCountyAnnotationCmd.Flags().StringVarP(&state, stateFlag, "s", "", "state the county is in")
	addCountyAnnotationCmd.MarkFlagRequired(stateFlag)

	countyFlag := "county"
	addCountyAnnotationCmd.Flags().StringVarP(&countyToAnnotate, countyFlag, "c", "", "county to to add annotation to")
	addCountyAnnotationCmd.MarkFlagRequired(countyFlag)

	annotationFlag := "annotation"
	addCountyAnnotationCmd.Flags().StringVarP(&annotation, annotationFlag, "a", "", "annotation string to add to county")
	addCountyAnnotationCmd.MarkFlagRequired(annotationFlag)

	RootCMD.AddCommand(addCountyAnnotationCmd)
}
