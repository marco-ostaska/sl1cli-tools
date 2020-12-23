package main

import (
	"fmt"
	"os"

	"github.com/marco-ostaska/sl1cli-tools/pkg/sl1user"
)

const version string = "v1.0.0 dec-2020"

type userStruct interface {
	getUsers()
}

func main() {

	if chkArgs() {
		var ud sl1user.UserDetails
		if err := ud.GetUserDetails(os.Args[1]); err != nil {
			fmt.Println(err)
			return
		}
		ud.PrintUserDetails()
	}

}
