package execusion

import (
	"github.com/danangkonang/crud-rest/migration/database/migration"
)

func RuningMigration() {
	migration.Animals()
}
