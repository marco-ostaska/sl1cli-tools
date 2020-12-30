/*
Copyright © 2020 Marco Ostaska

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

// Package wrapper wraps some boring tasks.
package wrapper

import (
	"strconv"
	"time"
)

// EpochToUnix converts epoch time to Unix time.
func EpochToUnix(epoch string) (time.Time, error) {
	var t time.Time
	i, err := strconv.ParseInt(epoch, 10, 64)
	if err != nil {
		return t, err
	}
	t = time.Unix(i, 0)
	return t, nil
}

// ReturnError checks for multiple errors at once.
func ReturnError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
