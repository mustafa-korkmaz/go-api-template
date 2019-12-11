// Package helper provides some common methods which can be used anywhere in application
package helper

import "time"

//UtcNow returns UTC.now
func UtcNow() time.Time {
	loc, _ := time.LoadLocation("UTC")

	return time.Now().In(loc)
}
