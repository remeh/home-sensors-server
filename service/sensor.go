// Copyright © 2015 - Rémy MATHIEU

package service

import (
	"fmt"
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

// GetValuesRange returns a slice of sensor values of the given type in the
// given interval of time.
func GetValuesRange(app *App, start, end time.Time, typ string) ([]SensorValue, error) {
	// First retrieve the values.
	data, err := app.Storage.SensorValueDAO.FindRange(start, end, typ)

	if err != nil {
		return data, fmt.Errorf("error: can't get range sensor values of type[%v] between start[%v] and end[%v]: %v", start, end, typ, err)
	}

	return data, nil
}
