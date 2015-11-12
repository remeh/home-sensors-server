// Copyright © 2015 - Rémy MATHIEU

package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	Conn *sql.DB

	SensorValueDAO *SensorValueDAO
	SystemEventDAO *SystemEventDAO
}

// Init opens a PostgreSQL connection with the given connectionString.
func (s *Storage) Init(connectionString string) (*sql.DB, error) {
	dbase, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	s.Conn = dbase

	// Creates all the DAOs of this storage.
	err = s.createDAOs()
	if err != nil {
		return nil, err
	}

	return dbase, s.Conn.Ping()
}

func (s *Storage) createDAOs() error {
	var err error
	if s.SensorValueDAO, err = NewSensorValueDAO(s.Conn); err != nil {
		return err
	}
	if s.SystemEventDAO, err = NewSystemEventDAO(s.Conn); err != nil {
		return err
	}
	return nil
}
