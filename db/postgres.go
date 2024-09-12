package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// postgresConfig Struct that represents the required info for connecting to the postgres db
type postgresConfig struct {
	user     string
	password string
	dbName   string
	host     string
	timeout  int
}

// Converts the configuration to the data source string for `sqlx`
func (c postgresConfig) getDataSourceName() string {
	return fmt.Sprintf("user=%v dbname=%v sslmode=disable connect_timeout=%v password=%v host=%v", c.user, c.dbName, c.timeout, c.password, c.host)
}

// Struct that represents the connection to the database
type Postgres struct {
	connection *sqlx.DB
}

// Gets the database connection
func GetDatabase() Postgres {
	config := postgresConfig{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		dbName:   os.Getenv("DB_NAME"),
		host:     os.Getenv("DB_HOST"),
		timeout:  5,
	}

	pg, err := sqlx.Connect("postgres", config.getDataSourceName())
	if err != nil {
		log.Fatalln(err)
	}

	return Postgres{connection: pg}
}
