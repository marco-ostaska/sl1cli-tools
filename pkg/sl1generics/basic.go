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

// BasicInfo is an abstraction for baisc api results
type BasicInfo []struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

// Load get basic struct from /api/x
func (bInfo *BasicInfo) Load(api string) error {
	var a httpcalls.APIData
	a.API = api
	err := a.NewRequest(&bInfo)
	if err != nil {
		return err
	}
	return nil
}

// ListBasic basic sl1 api returns
// with error formating
func (bInfo *BasicInfo) ListBasic(args []string, e ...string) []string {
	var result []string
	if len(args) == 0 {
		for _, u := range *bInfo {
			fmtStr := fmt.Sprintf("sl1id=%s(%s)", filepath.Base(u.URI), u.Description)
			result = append(result, fmtStr)
		}
	}

	if len(args) > 0 {
		for _, a := range args {
			id, err := bInfo.SearchByDesc(a)
			if err != nil {
				fmtStr := fmt.Sprintf("%v: %v %v", e[0], a, e[1])
				result = append(result, fmtStr)
				continue
			}
			fmtStr := fmt.Sprintf("sl1id=%s(%s)", filepath.Base(((*bInfo)[id].URI)), ((*bInfo)[id].Description))
			result = append(result, fmtStr)
		}

	}
	return result
}

// SearchByURI search description useing URI
func (bInfo *BasicInfo) SearchByURI(uri string) (string, error) {

	for _, u := range *bInfo {
		if uri == u.URI {
			return u.Description, nil
		}
	}
	return uri, fmt.Errorf("user not found for: %s", uri)

}

// SearchByDesc search from Basic.Description
func (bInfo *BasicInfo) SearchByDesc(s string) (int, error) {

	for i, u := range *bInfo {
		if s == u.Description {
			return i, nil
		}
	}
	return 0, fmt.Errorf("%s not found", s)

}

// Sl1ID returns sl1id
func (bInfo *BasicInfo) Sl1ID(s string) (string, error) {
	id, err := bInfo.SearchByDesc(s)
	if err != nil {
		return "", err
	}

	return path.Base((*bInfo)[id].URI), nil
}
