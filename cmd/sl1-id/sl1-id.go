package main

import (
	"fmt"
	"os"

	"github.com/marco-ostaska/sl1cli-tools/pkg/sl1user"
)

func main() {

	if chkArgs() {
		var usr sl1user.UserAcct
		if err := usr.GetIDs(); err != nil {
			fmt.Println(err)
			return
		}
		usr.PrintID(os.Args)
	}

}
