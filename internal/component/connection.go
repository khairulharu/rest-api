package component

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/khairulharu/restapi/internal/config"
	_ "github.com/lib/pq"
)

func NewDatabase(conf *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s "+"port=%s "+"user=%s "+"password=%s "+"dbname=%s "+"sslmode=disable",
		conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Pass, conf.Database.Name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("database connected")
	return db
}
