package db

import (
	"context"
	"fmt"
	"shrinklink/internal/logger"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Sqlx     *sqlx.DB
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewPostgresDB(host string, port string, username string, password string, database string) *PostgresDB {
	return &PostgresDB{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}

func (p *PostgresDB) DB() *sqlx.DB {
	return p.Sqlx
}

func (p *PostgresDB) Connect(ctx context.Context) error {
	var err error
	log := logger.CreateLoggerWithCtx(ctx)

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.Username, p.Password, p.Host, p.Port, p.Database,
	)

	p.Sqlx, err = sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Errorw("erorr connecting postgres", "err", err)
		return err
	}
	log.Infof("connected to postgres  %s", dbUrl)

	p.Sqlx.SetMaxIdleConns(1000)
	p.Sqlx.SetMaxOpenConns(5000)
	p.Sqlx.SetConnMaxLifetime(2 * time.Minute)
	return err
}

func (p *PostgresDB) Disconnect() error {
	return p.Sqlx.Close()
}
