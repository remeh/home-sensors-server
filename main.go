// Copyright © 2015 - Rémy MATHIEU

package main

import (
	. "github.com/remeh/home-sensors-server/app"
)

func main() {
	app := NewApp()
	declareRoutes(app)
	app.Start()
}

func declareRoutes(app *App) {
}
