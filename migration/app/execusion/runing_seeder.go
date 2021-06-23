package execusion

import (
	"io/ioutil"
	"os"

	"reflect"
	"strings"

	"github.com/danangkonang/crud-rest/migration/database/seed"
)

func RuningSeeder(tbl *Tables) {
	files, err := ioutil.ReadDir("migration/database/seed")
	if err != nil {
		os.Exit(0)
	}
	if len(tbl.NameTable) == 0 {
		newFile := []string{}
		for _, file := range files {
			filename := file.Name()
			list := strings.Split(filename, "_seeder_")
			name := list[0]
			if name != "0.core_type_seed.go" {
				newFile = append(newFile, name)
			}
		}
		tbl.NameTable = newFile
	}
	s := seed.MySeed{}
	for _, data_seeder := range tbl.NameTable {
		meth := reflect.ValueOf(s).MethodByName(strings.Title(data_seeder))
		meth.Call(nil)
	}
}
