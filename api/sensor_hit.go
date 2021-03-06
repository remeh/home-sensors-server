// Copyright © 2015 - Rémy MATHIEU

package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	. "github.com/remeh/home-sensors-server/app"
	"github.com/remeh/home-sensors-server/service"

	"github.com/gorilla/mux"
)

type SensorHit struct {
	App *App
}

type sensorHitBody struct {
	Values []struct {
		Type  string  `json:"type"`
		Value float64 `json:"value"`
	} `json:"values"`
}

func (c SensorHit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	r.ParseForm()

	sensorId := vars["sensor"]

	if len(sensorId) == 0 {
		w.WriteHeader(400)
		return
	}

	// read the body

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Println("err:", err)
		return
	}

	// unmarshal the body

	var body sensorHitBody
	err = json.Unmarshal(data, &body)
	if err != nil {
		w.WriteHeader(500)
		log.Println("err:", err)
		return
	}

	// store the values

	for _, v := range body.Values {
		err = service.StoreSensorValue(c.App, sensorId, v.Type, time.Now(), v.Value, extractIp(r))
		if err != nil {
			w.WriteHeader(500)
			log.Println("err: ", err)
			return
		}
	}
}
