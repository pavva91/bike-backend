package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pavva91/bike-backend/server/config"
)

var (
	DB *sql.DB
)

func MustConnectToDB(cfg config.ServerConfig) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.Timezone,
	)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err.Error())
		panic(fmt.Errorf("error connecting db: %w", err))
	}

	DB = database
}
