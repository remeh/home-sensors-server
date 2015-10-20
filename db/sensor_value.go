// Copyright © 2015 - Rémy MATHIEU

package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type SensorValueDAO struct {
	db *sql.DB
}

const (
	TEMP_FIELDS = `
		"sensor_value"."id",
		"sensor_value"."type",
		"sensor_value"."time",
		"sensor_value"."value"
	`
)

type SensorValue struct {
	Id    string
	Type  string
	Time  time.Time
	Value float32
}

func NewSensorValueDAO(db *sql.DB) (*SensorValueDAO, error) {
	dao := &SensorValueDAO{
		db: db,
	}
	err := dao.initStmt()
	return dao, err
}

func (d *SensorValueDAO) initStmt() error {
	var err error

	// TODO(remy)

	return err
}

// sensorValueFromRow reads an parking model from the current row.
func sensorValueFromRow(rows *sql.Rows) (SensorValue, error) {
	var id,
		typ string
	var t time.Time
	var value float32

	err := rows.Scan(
		&id,
		&typ,
		&t,
		&value)

	return SensorValue{
		Id:    id,
		Type:  typ,
		Time:  t,
		Value: value,
	}, err
}
