package boot

import (
	"database/sql"
	"fmt"

	"github.com/tren03/tr-coreutils/tr-todo/pkg/config"
)

func setupDatabase(cfg config.DB) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Connection pool settings
	db.SetMaxOpenConns(cfg.MaxOpen)
	db.SetMaxIdleConns(cfg.MaxIdle)

	return db, nil
}
