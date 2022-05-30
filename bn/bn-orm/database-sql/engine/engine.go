package engine

import (
	"GolandCode/bn/bn-orm/database-sql/log"
	"GolandCode/bn/bn-orm/database-sql/session"
	"database/sql"
	"fmt"
)

type Engine struct {
	db *sql.DB
}

func New(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(fmt.Errorf("error open %s from %s, %v", driver, source, err))
		return
	}
	if err := db.Ping(); err != nil {
		log.Error(fmt.Errorf("error connection, %v", err))
	}
	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
