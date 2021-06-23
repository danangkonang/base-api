package main

import (
	"os"
	"strings"

	"github.com/danangkonang/crud-rest/migration/app/helper"
	"github.com/danangkonang/crud-rest/migration/app/execusion"
)

func main() {
	if len(os.Args[1:]) == 0 {
		helper.PrintHelper()
		return
	}
	runCmd()
}

func runCmd() {
	switch os.Args[1] {
	case "run":
		migrationOrSeeder()
	case "reset":
		var t execusion.Tables
		if os.Args[2] != "" {
			t.NameTable = strings.Split(os.Args[2], ",")
		}
		execusion.ResetTables(&t)
	case "drop":
		var t execusion.Tables
		if os.Args[2] != "" {
			t.NameTable = strings.Split(os.Args[2], ",")
		}
		execusion.DropTables(&t)
	default:
		helper.PrintHelper()
	}
}

func migrationOrSeeder() {
	switch os.Args[2] {
	case "migration":
		var t execusion.Tables
		if os.Args[3] == "" {
			execusion.RuningMigration(&t)
		} else {
			t.NameTable = strings.Split(os.Args[3], ",")
			execusion.RuningMigration(&t)
		}
	case "seeder":
		var t execusion.Tables
		if os.Args[3] == "" {
			execusion.RuningSeeder(&t)
		} else {
			t.NameTable = strings.Split(os.Args[3], ",")
			execusion.RuningSeeder(&t)
		}
	default:
		helper.PrintHelper()
	}
}
