package migration

import (
	"github.com/glendecado/template/model"
)

// migrate holds pointers to your model structs
var Migrate = []interface{}{
	&model.User{},
	// Add other models here as needed
}
