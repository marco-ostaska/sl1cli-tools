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
		var usr sl1user.UserAcct
		if err := usr.GetIDs(); err != nil {
			fmt.Println(err)
			return
		}

		id, err := usr.Sl1UserID(os.Args[1])
		if err != nil {
			panic(err)
		}

		var ud sl1user.UserDetails
		if err := ud.LoadUserDetails(id); err != nil {
			panic(err)
		}
		ud.PrintUserDetails()
	}

}
