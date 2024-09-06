package schemas

import (
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Hole    string
	Company string
	Locate  string
	Remote  bool
	Link    string
	Salary  int64
}
