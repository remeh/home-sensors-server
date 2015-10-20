package db

import (
	"strings"
)

// insertFields must be used to inject fields while
// writing an insert query.
// This is because PostgreSQL doesn't support the notation :
// INSERT INTO "table" ( "table"."field" ) VALUES ( $1 )
// but only supports :
// INSERT INTO "table" ( "field" ) VALUES ( $1 )
func insertFields(tableName string, fields string) string {
	return strings.Replace(fields, "\""+tableName+"\".", "", -1)
}
