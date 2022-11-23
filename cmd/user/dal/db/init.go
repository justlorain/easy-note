package db

import (
	"easy-note/pkg/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt: true,
			// use transaction
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	// set tracing and metrics
	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&User{}) {
		return
	}
	if err = m.CreateTable(&User{}); err != nil {
		panic(err)
	}
}
