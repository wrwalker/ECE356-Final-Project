package utils

import (
	"fmt"
	sql "github.com/jmoiron/sqlx"
)

func DeserializeRowsToMappedInterface(r *sql.Rows) ([]map[string]interface{}, []string, error) {
	var allRows []map[string]interface{}
	var colHeaders []string
	for r.Next() {
		results := make(map[string]interface{})

		err := r.MapScan(results) // this can cause errors with non-fully qualified names
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		colHeaders, err = r.Columns()
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		allRows = append(allRows, results)
	}
	return allRows, colHeaders, nil
}

func PrintMap(m []map[string]interface{}, colNames []string) {
	for _, k := range colNames {
		fmt.Printf("%20s| ", k)
	}
	fmt.Println()
	for range colNames {
		fmt.Print("------------------------")
	}
	fmt.Println()

	for _, row := range m {
		for _, k := range colNames {
			fmt.Printf("%20s| ", row[k])
		}
		fmt.Println()
	}
}
