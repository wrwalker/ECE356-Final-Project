package main

import (
	"github.com/ECE356-Final-Project/SQLClient/internal/queryMaker"
	"github.com/ECE356-Final-Project/SQLClient/internal/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	qm := queryMaker.NewQueryMaker()
	defer qm.Db.Close()

	res, _ := qm.DoQuery("SELECT * FROM VotesByState")
	rows, colNames, _ := utils.DeserializeRowsToMappedInterface(res)
	utils.PrintMap(rows, colNames)
}
