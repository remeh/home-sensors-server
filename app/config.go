// Copyright © 2015 - Rémy MATHIEU

package app

import (
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Addr      string `envconfig:"ADDR,default=:9000,optional"`
	Key       string `envconfig:"KEY,optional"`
	PublicDir string `envconfig:"DIR,default=public/,optional"`
}

func ReadConfig() Config {
	var conf Config
	// NOTE(remy): don't read the error because it'll
	// fallback on default values on error
	envconfig.Init(&conf)
	return conf
}
