// Copyright © 2015 - Rémy MATHIEU

package app

import (
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Addr      string `envconfig:"ADDR,default=:9000,optional"`
	Key       string `envconfig:"KEY,optional"`
	PublicDir string `envconfig:"DIR,default=public/,optional"`
	Conn      string `envconfig:"CONN,default=host=/var/run/postgresql sslmode=disable user=hss dbname=hss`
}

func ReadConfig() Config {
	var conf Config
	envconfig.Init(&conf)
	return conf
}
