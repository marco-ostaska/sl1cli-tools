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

// Package sl1user have the routines for /api/account
package sl1user

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/marco-ostaska/sl1cmd/pkg/httpcalls"
)

// UserAcct is an abstraction to /api/account
type UserAcct []struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

// GetIDs get user IDs from /api/account
func (ua *UserAcct) GetIDs() error {
	var api httpcalls.APIData
	api.API = "/api/account/"
	err := api.NewRequest(&ua)
	if err != nil {
		return err
	}
	return nil
}

// PrintID print userId from /api/account
func (ua *UserAcct) PrintID(args []string) {
	if len(args) == 0 {
		for _, u := range *ua {
			fmt.Printf("sl1id=%s(%s)\n", filepath.Base(u.URI), u.Description)
		}
	}

	if len(args) > 0 {
		for _, a := range args {
			id, err := ua.ID(a)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("sl1id=%s(%s)\n", filepath.Base(((*ua)[id].URI)), ((*ua)[id].Description))
		}
	}
}

// SearchByURI search user name by provided URI
func (ua *UserAcct) SearchByURI(uri string) (string, error) {

	for _, u := range *ua {
		if uri == u.URI {
			return u.Description, nil
		}
	}
	return uri, fmt.Errorf("user not found for: %s", uri)

}

// ID returns a specific user ID index from UserAcct
func (ua *UserAcct) ID(user string) (int, error) {

	for i, u := range *ua {
		if user == u.Description {
			return i, nil
		}
	}
	return 0, fmt.Errorf("id: %s: no such user", user)

}

// Sl1UserID returns sl1id from user
func (ua *UserAcct) Sl1UserID(user string) (string, error) {
	id, err := ua.ID(user)
	if err != nil {
		return "", err
	}

	return path.Base((*ua)[id].URI), nil
}
