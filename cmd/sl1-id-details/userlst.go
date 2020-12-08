package main

import (
	"fmt"

	"github.com/marco-ostaska/sl1cli-tools/pkg/apirequest.go"
)

type sl1UserLst []struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

func (dat *sl1UserLst) getUsers() {
	var api apirequest.APIData
	err := api.NewRequest(dat, "/api/account/")
	if err != nil {
		fmt.Println(err)
	}

	// for _, u := range dat {
	// 	if u.Description == os.Args[1] {
	// 		getDetails(filepath.Base(u.URI))
	// 	}
	// }
	// return

}
