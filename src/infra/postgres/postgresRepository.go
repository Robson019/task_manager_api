package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"time"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectionManager)(nil)

type DatabaseConnectionManager struct{}

func (dcm DatabaseConnectionManager) getConnection() (*sqlx.DB, error) {
	// Load vars env
	uri := getDatabaseURI()

	// init database
	db, err := sqlx.Open("postgres", uri)

	if err != nil {
		log.Print("Error while accessing database: " + err.Error())
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func (dcm DatabaseConnectionManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()
	if err != nil {
		log.Error().Err(err)
	}
}
