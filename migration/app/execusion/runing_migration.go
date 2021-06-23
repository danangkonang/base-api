package execusion

import (
	"io/ioutil"
	"os"

	"reflect"
	"strings"

	"github.com/danangkonang/crud-rest/migration/database/migration"
)

func RuningMigration(tbl *Tables) {
	files, err := ioutil.ReadDir("migration/database/migration")
	if err != nil {
		os.Exit(0)
	}
	if len(tbl.NameTable) == 0 {
		newFile := []string{}
		for _, file := range files {
			filename := file.Name()
			list := strings.Split(filename, "_migration_")
			if list[0] != "0.core_type_migration.go" {
				name := list[1]
				tb_name := strings.Split(name, ".go")
				newFile = append(newFile, tb_name[0])
			}
		}
		tbl.NameTable = newFile
	}
	m := migration.MyMigration{}
	for _, migrate := range tbl.NameTable {
		meth := reflect.ValueOf(m).MethodByName(strings.Title(migrate))
		meth.Call(nil)
	}
}
