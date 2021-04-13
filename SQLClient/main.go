package main

import (
	"fmt"
	"github.com/ECE356-Final-Project/SQLClient/internal/queryMaker"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	qm := queryMaker.NewQueryMaker()
	defer qm.Db.Close()

	res, _ := qm.DoQuery("SELECT * FROM VotesByState")
	rows, colNames, _ := queryMaker.DeserializeRows(res) //use this so order is deterministic
	for _, row := range rows {
		for _, k := range colNames {
			fmt.Printf("%s: %s, ", k, row[k])
		}
		fmt.Println()
	}

}
