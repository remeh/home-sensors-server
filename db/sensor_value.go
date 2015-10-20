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
		"sensor_value"."sensor_id",
		"sensor_value"."type",
		"sensor_value"."time",
		"sensor_value"."value"
	`
)

type SensorValue struct {
	SensorId string
	Type     string
	Time     time.Time
	Value    float64
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
	var sensorId,
		typ string
	var t time.Time
	var value float64

	err := rows.Scan(
		&sensorId,
		&typ,
		&t,
		&value)

	return SensorValue{
		SensorId: sensorId,
		Type:     typ,
		Time:     t,
		Value:    value,
	}, err
}
