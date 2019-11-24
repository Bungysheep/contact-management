package db

import (
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/config"
	"github.com/bungysheep/contact-management/pkg/logger"
)

// OpenDbConn opens connection to postgres database
func OpenDbConn() (*sql.DB, error) {
	logger.Log.Info("Database connection is opening...")

	db, err := sql.Open("postgres", config.CONNECTIONSTRING)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
