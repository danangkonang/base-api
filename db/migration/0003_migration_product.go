package migration

import (
	"fmt"
	"os"
)

func (m *Migration) Product() {
	query := `
		CREATE TABLE product(
			id serial PRIMARY KEY,
			name VARCHAR (225) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`
	_, err := Connection().Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
  fmt.Println(string(Green), "success", string(Reset), "create table 0003_migration_product.go")
}
