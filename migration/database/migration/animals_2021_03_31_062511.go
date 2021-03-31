package migration

import (
	"fmt"
	"log"

	"github.com/danangkonang/crud-rest/migration/app/config"
)

func Animals() {
	db := config.Connect()
	db.Exec(`DROP TABLE animals`)
	_, err := db.Exec(`
	CREATE TABLE animals(
		animal_id VARCHAR (32) NOT NULL PRIMARY KEY,
		name        VARCHAR (225) NOT NULL,
		color       VARCHAR (225) NOT NULL,
		description VARCHAR (225) NOT NULL,
		image       VARCHAR (225) NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success create table animals_2021_03_31_062511.go")
}
