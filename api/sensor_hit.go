// Copyright © 2015 - Rémy MATHIEU

package api

import (
	"log"
	"net/http"
	"strconv"
	"time"

	. "github.com/remeh/home-sensors-server/app"
	"github.com/remeh/home-sensors-server/service"

	"github.com/gorilla/mux"
)

type SensorHit struct {
	App *App
}

func (c SensorHit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	r.ParseForm()

	sensorId := vars["sensor"]
	typ := vars["type"]
	v := r.Form.Get("v")

	if len(sensorId) == 0 || len(typ) == 0 || len(v) == 0 {
		w.WriteHeader(400)
		return
	}

	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = service.StoreSensorValue(c.App, sensorId, typ, time.Now(), value)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
}
