// Copyright © 2015 - Rémy MATHIEU

package service

import (
	"time"

	. "github.com/remeh/home-sensors-server/app"
	. "github.com/remeh/home-sensors-server/db"
)

func StoreSystemEvent(app *App, systemId, message string) error {
	dao := app.Storage.SystemEventDAO
	_, err := dao.Insert(SystemEvent{
		SystemId: systemId,
		Time:     time.Now(),
		Message:  message,
	})
	return err
}
