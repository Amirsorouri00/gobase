package app

import (
	"database/sql"
	"fmt"
	stdlog "log"
	"portfolio/app/migration"
	conf "portfolio/services/infrastructure/config/auth"
	"portfolio/services/infrastructure/log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type dbLogger struct{}

func (dbLogger) Fatal(v ...interface{}) {
	stdlog.Fatal(v...)
}

func (dbLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (dbLogger) Print(v ...interface{}) {
	stdlog.Print(v...)
}

func (dbLogger) Println(v ...interface{}) {
	stdlog.Println(v...)
}

func (dbLogger) Printf(format string, v ...interface{}) {
	log.Infof(format, v...)
}

func InitDB(cfg conf.DB) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Driver,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SslMode,
	))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	goose.SetLogger(dbLogger{})
	goose.SetTableName("migrations")
	err = goose.SetDialect("postgres")
	if err != nil {
		log.Debugf("failed to goose SetDialect: %v", err)
	}
	goose.SetBaseFS(migration.FS)
	if err = goose.Up(db, "scripts"); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
