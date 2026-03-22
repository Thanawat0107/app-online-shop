package config

import (
	"log"

	"github.com/Thanawat0107/app-online-shop/internal/infra/database"
)

type Config struct {
	Env      *Env
	database *DbConn
}

func NewConfig() *Config {
	env := NewEnv()
	dbs := NewDbConn(env)
	return &Config{
		Env:      env,
		database: dbs,
	}
}

func (c *Config) GetDb(key string) database.Database {
	if db, exist := c.database.dbs[key]; exist {
		return db
	}
	log.Fatalf("Error: database with name=%s not found", key)
	return nil
}
