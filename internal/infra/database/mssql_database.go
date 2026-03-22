package database

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type sqlServerDatabase struct {
	*gorm.DB
}

var (
	instance *sqlServerDatabase
	once     sync.Once
)

func NewSqlServerDatabase(host, username, password, db string) *sqlServerDatabase {
	once.Do(func() {
		connString := fmt.Sprintf(
			"server=%s;database=%s;trusted_connection=yes",
			host,
			db,
		)

		conn, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Database connect to SQL server success")
		instance = &sqlServerDatabase{conn}
	})
	return instance
}

func (s *sqlServerDatabase) Connect() *gorm.DB {
	return s.DB
}
