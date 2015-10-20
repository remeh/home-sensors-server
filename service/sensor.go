// Copyright © 2015 - Rémy MATHIEU

package service

import (
	"time"

	. "github.com/remeh/home-sensors-server/app"
	. "github.com/remeh/home-sensors-server/db"
)

func StoreSensorValue(app *App, sensorId, typ string, t time.Time, value float64) error {
	dao := app.Storage.SensorValueDAO
	last, err := dao.FindLast(sensorId, typ)
	if err != nil {
		return err
	}

	// No need to insert the new value
	if last.Value == value {
		return nil
	}

	_, err = dao.Insert(SensorValue{
		SensorId: sensorId,
		Type:     typ,
		Time:     t,
		Value:    value,
	})
	return err
}
