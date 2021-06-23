package migration

import (
	"fmt"
	"os"

	"github.com/danangkonang/crud-rest/migration/app/config"
)

func (m *MyMigration) Animals() {
	db := config.Connect()
	_, err := db.Exec(`
		CREATE TABLE animals(
			animal_id VARCHAR (32) NOT NULL PRIMARY KEY,
			name        VARCHAR (225) NOT NULL,
			color       VARCHAR (225) NOT NULL,
			description VARCHAR (225) NOT NULL,
			image       VARCHAR (225) NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);
	`)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Println("success create table 2021_06_23_201908_migration_animals.go")
}
