package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/marco-ostaska/sl1cli-tools/pkg/apipost"
	"github.com/marco-ostaska/sl1cli-tools/pkg/sl1user"
)

func main() {
	if chkArgs() {
		if chkArgs() {
			var usr sl1user.UserAcct
			if err := usr.GetIDs(); err != nil {
				fmt.Println(err)
				return
			}
			i, err := usr.ID(os.Args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s no such user\n", path.Base(os.Args[0]), os.Args[1])
				os.Exit(1)
			}

			fmt.Println("deleting:", (usr)[i].Description)

			var p apipost.APIData
			p.API = (usr)[i].URI

			if err := p.APIDelete(); err != nil {
				log.Fatal(err)
			}

			if string(p.Result) == "" {
				fmt.Println((usr)[i].Description, "deleted successfully")
				return
			}
			fmt.Println()
			fmt.Println(string(p.Result))

		}
	}

}
