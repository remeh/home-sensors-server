// Copyright © 2015 - Rémy MATHIEU

package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

const (
	SENSOR_VALUE_FIELDS = `
		"sensor_value"."sensor_id",
		"sensor_value"."type",
		"sensor_value"."time",
		"sensor_value"."value"
	`
)

type SensorValueDAO struct {
	db *sql.DB

	findLast  *sql.Stmt
	findRange *sql.Stmt
	insert    *sql.Stmt
}

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

	if d.findRange, err = d.db.Prepare(`
		SELECT ` +
		SENSOR_VALUE_FIELDS + `
		FROM "sensor_value"
		WHERE
			"sensor_value"."time" >= $1
			AND
			"sensor_value"."time" <= $2
			AND
			"sensor_value"."type" = $3
		ORDER BY "sensor_value"."time"
	`); err != nil {
		return err
	}

	if d.findLast, err = d.db.Prepare(`
		SELECT ` +
		SENSOR_VALUE_FIELDS + `
		FROM "sensor_value"
		WHERE
			"sensor_value"."sensor_id" = $1
			AND
			"sensor_value"."type" = $2
		ORDER BY "sensor_value"."time" DESC
		LIMIT 1
	`); err != nil {
		return err
	}

	if d.insert, err = d.db.Prepare(`
		INSERT INTO
		"sensor_value"
		(` + insertFields("sensor_value", SENSOR_VALUE_FIELDS) + `)
		VALUES
		($1, $2, $3, $4)
	`); err != nil {
		return err
	}

	return nil
}

func (d *SensorValueDAO) Insert(sensorValue SensorValue) (sql.Result, error) {
	return d.insert.Exec(
		sensorValue.SensorId,
		sensorValue.Type,
		sensorValue.Time,
		sensorValue.Value,
	)
}

func (d *SensorValueDAO) FindRange(start, end time.Time, typ string) ([]SensorValue, error) {
	return readValues(d.findRange.Query(start, end, typ))
}

func (d *SensorValueDAO) FindLast(sensorId string, typ string) (SensorValue, error) {
	return ReadSensorValueAndReturn(d.findLast.Query(sensorId, typ))
}

func ReadSensorValueAndReturn(rows *sql.Rows, err error) (SensorValue, error) {
	var rv SensorValue

	if err != nil {
		return rv, err
	}

	if rows == nil {
		return rv, nil
	}

	defer rows.Close()

	if rows.Next() {
		rv, err = sensorValueFromRow(rows)
	}

	return rv, err
}

func readValues(rows *sql.Rows, err error) ([]SensorValue, error) {
	rv := make([]SensorValue, 0)
	if rows == nil {
		return rv, nil
	}

	defer rows.Close()

	for rows.Next() {
		var v SensorValue
		var err error
		if v, err = sensorValueFromRow(rows); err != nil {
			return rv, err
		}
		rv = append(rv, v)
	}

	return rv, nil
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
