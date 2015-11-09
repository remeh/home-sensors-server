// Copyright © 2015 - Rémy MATHIEU

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	. "github.com/remeh/home-sensors-server/app"
	"github.com/remeh/home-sensors-server/service"

	"github.com/gorilla/mux"
)

type GetGraph struct {
	App *App
}

type getGraphResponse struct {
	Values []timedValue `json:"values"`
}

type timedValue struct {
	T time.Time `json:"t"`
	V float64   `json:"v"`
}

func (c GetGraph) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read the parameters

	vars := mux.Vars(r)
	r.ParseForm()

	typ := vars["type"]

	if len(typ) == 0 {
		w.WriteHeader(400)
		return
	}

	pStart, pEnd := r.Form.Get("start"), r.Form.Get("end")

	if len(pStart) == 0 || len(pEnd) == 0 {
		w.WriteHeader(400)
		return
	}

	// parse the parameters

	start, end := parseTime(pStart), parseTime(pEnd)

	if start.IsZero() || end.IsZero() {
		w.WriteHeader(400)
		return
	}

	// get the data

	values, err := service.GetValuesRange(c.App, start, end, typ)
	if err != nil {
		log.Println("error: while getting values range:", err.Error())
	}

	// render the data

	var resp getGraphResponse
	resp.Values = make([]timedValue, len(values))

	for i, v := range values {
		resp.Values[i] = timedValue{
			T: v.Time,
			V: v.Value,
		}
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("error: while marshaling the GetGraph response: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
