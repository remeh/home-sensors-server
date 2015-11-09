// Copyright © 2015 - Rémy MATHIEU

package api

import (
	"time"
)

// parseTime reads the given string to return a valid time.
// A zero time is returned if the string is not parsable.
func parseTime(t string) time.Time {
	rv, err := time.Parse("2006-01-02 15:04", t)
	if err != nil {
		return time.Time{}
	}
	return rv
}
