package handler

import (
	"github.com/kayky18/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHanler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
