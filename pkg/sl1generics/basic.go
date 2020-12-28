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

package sl1generics

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/marco-ostaska/sl1cmd/pkg/httpcalls"
)

// Basic is an abstraction for baisc api results
type Basic []struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

// GetIDs get user IDs from /api/account
func (bStruct *Basic) GetIDs(secAPI string) error {
	var api httpcalls.APIData
	err := api.NewRequest(&bStruct, "/api/"+secAPI)
	if err != nil {
		return err
	}
	return nil
}

// PrintBasic print basic sl1 api returns
func (bStruct *Basic) PrintBasic(args []string) {
	if len(args) == 0 {
		for _, u := range *bStruct {
			fmt.Printf("sl1id=%s(%s)\n", filepath.Base(u.URI), u.Description)
		}
	}

	if len(args) > 0 {
		for _, a := range args {
			id, err := bStruct.ID(a)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("sl1id=%s(%s)\n", filepath.Base(((*bStruct)[id].URI)), ((*bStruct)[id].Description))
		}
	}
}

// SearchByURI search description useing URI
func (bStruct *Basic) SearchByURI(uri string) (string, error) {

	for _, u := range *bStruct {
		if uri == u.URI {
			return u.Description, nil
		}
	}
	return uri, fmt.Errorf("user not found for: %s", uri)

}

// ID returns a specific user ID index from Basic
func (bStruct *Basic) ID(user string) (int, error) {

	for i, u := range *bStruct {
		if user == u.Description {
			return i, nil
		}
	}
	return 0, fmt.Errorf("id: %s: no such user", user)

}

// Sl1UserID returns sl1id from user
func (bStruct *Basic) Sl1UserID(user string) (string, error) {
	id, err := bStruct.ID(user)
	if err != nil {
		return "", err
	}

	return path.Base((*bStruct)[id].URI), nil
}
