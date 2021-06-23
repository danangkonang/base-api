package execusion

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/danangkonang/crud-rest/migration/app/config"
)

func DropTables(tb *Tables) {
	files, err := ioutil.ReadDir("migration/database/migration")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	db := config.Connect()
	if len(tb.NameTable) > 0 {
		for _, ntb := range tb.NameTable {
			query := "DROP TABLE IF EXISTS " + ntb + ";"
			_, err := db.Exec(query)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			fmt.Println("success DROP TABLE " + ntb)
		}
	} else {
		for _, file := range files {
			filename := file.Name()
			list := strings.Split(filename, "_migration_")
			name := list[0]
			if name != "0.core_type_migration.go" {
				tb_name := strings.Split(list[1], ".go")
				query := "DROP TABLE IF EXISTS " + tb_name[0] + ";"
				_, err := db.Exec(query)
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}
				fmt.Println("success DROP TABLE " + tb_name[0])
			}
		}
	}
}
