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

// Package sl1 Package account implements utility routines for manipulating /api
package sl1

import (
	"fmt"

	"github.com/marco-ostaska/sl1cmd/pkg/sl1/httpcalls"
)

// These are predefined shortcuts for sl1 APIs sections.
const (
	AccountAPI = "/api/account"
)

// BasicInfo is an abstraction for primitives api results
type BasicInfo []struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
	ID          string
}

// Load loads BasicInfo from /api/xyz
func (bInfo *BasicInfo) Load(api string) error {
	var a httpcalls.APIData
	a.API = api
	err := a.NewRequest(&bInfo)
	if err != nil {
		return err
	}

	for i := 0; i < len(*bInfo); i++ {
		id, err := bInfo.LookupIDbyDesc((*bInfo)[i].Description)
		if err != nil {
			return err
		}
		(*bInfo)[i].ID = id
	}

	return nil
}

//ListByDesc returns the list of index for provided description(s) lookup
func (bInfo *BasicInfo) ListByDesc(args []string) ([]int, []error) {
	var found []int
	var notFound []error

	for _, a := range args {
		i, err := bInfo.LookupIdxByDesc(a)
		if err != nil {
			notFound = append(notFound, err)
			continue
		}
		found = append(found, i)
	}

	return found, notFound

}

// FmtMsg returns an user friendly output format based on provided index
func (bInfo *BasicInfo) FmtMsg(i int) string {
	return fmt.Sprintf("sl1id=%s(%s)", (*bInfo)[i].ID, (*bInfo)[i].Description)
}
