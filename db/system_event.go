// Copyright © 2015 - Rémy MATHIEU

package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

const (
	SYSTEM_EVENT_FIELDS = `
		"system_event"."system_id",
		"system_event"."time",
		"system_event"."message"
	`
)

type SystemEventDAO struct {
	db *sql.DB

	insert *sql.Stmt
}

type SystemEvent struct {
	SystemId string
	Time     time.Time
	Message  string
}

func NewSystemEventDAO(db *sql.DB) (*SystemEventDAO, error) {
	dao := &SystemEventDAO{
		db: db,
	}
	err := dao.initStmt()
	return dao, err
}

func (d *SystemEventDAO) initStmt() error {
	var err error

	if d.insert, err = d.db.Prepare(`
		INSERT INTO
		"system_event"
		(` + insertFields("system_event", SYSTEM_EVENT_FIELDS) + `)
		VALUES
		($1, $2, $3)
	`); err != nil {
		return err
	}

	return nil
}

func (d *SystemEventDAO) Insert(systemEvent SystemEvent) (sql.Result, error) {
	return d.insert.Exec(
		systemEvent.SystemId,
		systemEvent.Time,
		systemEvent.Message,
	)
}
