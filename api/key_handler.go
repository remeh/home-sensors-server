// Copyright © 2015 - Rémy MATHIEU

package api

import (
	"net/http"

	. "github.com/remeh/home-sensors-server/app"
)

type KeyHandler struct {
	App     *App
	Handler http.Handler
}

func (c KeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(c.App.Config.Key) > 0 {
		if len(r.Header.Get(c.App.Config.Key)) == 0 {
			w.WriteHeader(401)
			return
		}
	}

	c.Handler.ServeHTTP(w, r)
}
