// Package sl1generics have the generic routines
//to be used throughout the sl1tools
package sl1generics

import (
	"strconv"
	"time"
)

// EpochToUnix converts epoch time to Unix time
func EpochToUnix(epoch string) (time.Time, error) {
	var t time.Time
	i, err := strconv.ParseInt(epoch, 10, 64)
	if err != nil {
		return t, err
	}
	t = time.Unix(i, 0)
	return t, nil
}
