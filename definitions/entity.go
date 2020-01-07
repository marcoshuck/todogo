package definitions

import (
	"github.com/marcoshuck/todogo/database"
)

// Entity represents a generic entity to be used by Controllers and Services
type Entity struct {
	database.Model
	UUID string
}
