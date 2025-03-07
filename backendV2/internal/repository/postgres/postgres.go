package postgres

import (
	"fmt"
	"queueAppV2/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func MustCreate(cfg *config.Config) *sqlx.DB {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.Name, cfg.Db.Password, cfg.Db.Dbname, cfg.Db.Sslmode))
	if err != nil {
		panic(fmt.Sprint("error occurred creating database connection: ", err))
	}
	if err = db.Ping(); err != nil {
		panic(fmt.Sprint("error occurred checking connection to database: ", err))
	}
	return db
}





