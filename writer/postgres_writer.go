package writer

import (
	"database/sql"
	"ekyc/config"
	"fmt"
	"time"
)

// PostgresWriter implements the SQLWriter interface for writing SQL commands to the postgres DB.
type PostgresWriter struct {
	DBConfig config.DBConfig
}

// NewPostgresWriter creates a new instance of PostgresWriter.
func NewPostgresWriter(dbConfig config.DBConfig) SQLWriter {
	return &PostgresWriter{
		DBConfig: dbConfig,
	}
}

const (
	POSTGRES string = "postgres"
)

func (p *PostgresWriter) WriteSQL() {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		p.DBConfig.DBUserName,
		p.DBConfig.DBPassword,
		p.DBConfig.DBHost,
		p.DBConfig.DBPort,
		p.DBConfig.DBName,
	)
	postgresConn, err := sql.Open(POSTGRES, connStr)
	if err != nil {
		panic(err)
	}
	defer postgresConn.Close()

	// Set connection pool properties
	postgresConn.SetMaxOpenConns(10)
	postgresConn.SetMaxIdleConns(5)
	postgresConn.SetConnMaxLifetime(5 * time.Minute)

	err = postgresConn.Ping()
	if err != nil {
		panic(err)
	}

	tx, err := postgresConn.Begin()
	if err != nil {
		panic(err)
	}

	// Commit the remaining queries
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
