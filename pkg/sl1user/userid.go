// Package sl1user have the routines for /api/account
package sl1user

import (
	"fmt"
	"path/filepath"

	"github.com/marco-ostaska/sl1cli-tools/pkg/apirequest.go"
)

// UserAcct is an abstraction to /api/account
type UserAcct []struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

// GetIDs get user IDs from /api/account
func (ua *UserAcct) GetIDs() error {
	var api apirequest.APIData
	err := api.NewRequest(&ua, "/api/account/")
	if err != nil {
		return err
	}
	return nil
}

// PrintID print userId from /api/account
func (ua *UserAcct) PrintID(args []string) {
	if len(args) == 1 {
		for _, u := range *ua {
			fmt.Printf("sl1id=%s(%s)\n", filepath.Base(u.URI), u.Description)
		}
	}

	if len(args) > 1 {
		for _, a := range args[1:] {
			id, err := ua.ID(a)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("sl1id=%s(%s)\n", filepath.Base(((*ua)[id].URI)), ((*ua)[id].Description))
		}
	}
}

// ID returns a specific user ID index on UserAcct
func (ua *UserAcct) ID(user string) (int, error) {

	for i, u := range *ua {
		if user == u.Description {
			return i, nil
		}
	}
	return 0, fmt.Errorf("id: %s: no such user", user)

}
