// Copyright © 2015 - Rémy MATHIEU

package main

import (
	. "github.com/remeh/home-sensors-server/api"
	. "github.com/remeh/home-sensors-server/app"
)

func main() {
	app := NewApp()
	declareRoutes(app)
	app.Start()
}

func declareRoutes(app *App) {
	app.Add("/api/hit/{sensor}", KeyHandler{app, SensorHit{app}}, "POST")
	app.Add("/api/graph/{type}.json", KeyHandler{app, GetGraph{app}}, "GET")
}
