// Copyright © 2015 - Rémy MATHIEU

package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	. "github.com/remeh/home-sensors-server/app"
	"github.com/remeh/home-sensors-server/service"
)

type SystemHit struct {
	App *App
}

type systemHitBody struct {
	System  string `json:"system"`
	Message string `json:"message"`
}

func (c SystemHit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read the body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Println("err:", err)
		return
	}

	// unmarshal the body

	var body systemHitBody
	err = json.Unmarshal(data, &body)
	if err != nil {
		w.WriteHeader(500)
		log.Println("err:", err)
		return
	}

	// store the event

	err = service.StoreSystemEvent(c.App, body.System, body.Message, extractIp(r))
	if err != nil {
		w.WriteHeader(500)
		log.Println("err: ", err)
		return
	}
}

func extractIp(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
