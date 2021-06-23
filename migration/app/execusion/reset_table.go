package execusion

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/danangkonang/crud-rest/migration/app/config"
)

func ResetTables(tb *Tables) {
	files, err := ioutil.ReadDir("migration/database/migration")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	db := config.Connect()
	if len(tb.NameTable) > 0 {
		for _, ntb := range tb.NameTable {
			query := "TRUNCATE " + ntb + ";"
			_, err := db.Exec(query)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			fmt.Println("success delete row")
		}
	} else {
		for _, file := range files {
			filename := file.Name()
			list := strings.Split(filename, "_migration_")
			name := list[0]
			if name != "0.core_type_migration.go" {
				query := "TRUNCATE " + name + ";"
				_, err := db.Exec(query)
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}
				fmt.Println("success delete row")
			}
		}
	}
}
