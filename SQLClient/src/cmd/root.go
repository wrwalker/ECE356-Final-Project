package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const rootStr = "SQLClient"

var Verbose bool

var RootCMD = &cobra.Command{
	Use:   rootStr,
	Short: fmt.Sprintf("%s is a command line interface (CLI) to interact with the Elections database", rootStr),
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		log.Printf("err: %s", err.Error())
		os.Exit(1)
	}
}

func init() {
	verboseFlag := "verbose"
	RootCMD.PersistentFlags().BoolVarP(&Verbose, verboseFlag, "v", false, "verbose output")
}
