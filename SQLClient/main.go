package main

import (
	"github.com/ECE356-Final-Project/SQLClient/internal/queryMaker"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	qm := queryMaker.NewQueryMaker()
	defer qm.Db.Close()

	res, _ := qm.DoQuery("SELECT * FROM VotesByState")
	queryMaker.DeserializeRows(res)

}
