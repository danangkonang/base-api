package migration

import (
	"fmt"
	"os"
)

func (m *Migration) Users() {
	query := `
		CREATE TABLE users(
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR (225) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := Connection().Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
  fmt.Println(string(Green), "success", string(Reset), "create table 0002_migration_users.go")
}
