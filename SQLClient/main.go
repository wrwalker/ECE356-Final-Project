package main

import (
	"github.com/ECE356-Final-Project/SQLClient/cmd"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cmd.Execute()
}
