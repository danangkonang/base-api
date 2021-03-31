package main

import (
	"os"

	"github.com/danangkonang/crud-rest/migration/app/helper"
	"github.com/danangkonang/crud-rest/migration/app/execusion"
)

func main() {
	arrCmd := os.Args[1:]
	if len(arrCmd) == 0 {
		helper.PrintHelper()
		return
	}
	runCmd()
}

func runCmd() {
	usrCmd := os.Args[1]
	switch usrCmd {
	case "migration":
		execusion.RuningMigration()
		break
	case "seeder":
		execusion.RuningSeeder() 
		break
	default:
		helper.PrintHelper()
	}
}
