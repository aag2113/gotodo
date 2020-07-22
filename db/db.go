package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var DataBase *sql.DB

func Load() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error

	// TODO: Refactor this to be less clunky
	cnt := 0
	for cnt < 20 {
		DataBase, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			time.Sleep(5 * time.Second)
			cnt++
		} else {
			cnt = 100
		}
	}
	if err != nil {
		panic(err)
	}

	err = DataBase.Ping()
	if err != nil {
		panic(err)
	}

	createSchema()
}

func createSchema() error {
	_, err := DataBase.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id VARCHAR(50) UNIQUE NOT NULL, 
		title VARCHAR(100) NOT NULL, 
		created_at TIMESTAMP NOT NULL, 
		status VARCHAR(50) NOT NULL);`)
	return err
}

func Kill() {
	DataBase.Close()
}
