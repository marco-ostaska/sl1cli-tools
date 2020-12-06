package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/marco-ostaska/sl1cli-tools/internal/pkg/apirequest.go"
)

const version string = "v1.0.0 dec-2020"

type sl1UserLst []struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

func usage() {
	usage := `Print sl1 user information for the specified USERS,
or (when USER omitted) prints a list of all users.

  -v             version
  -h             display this and exit
`
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [USERS]\n%v", filepath.Base(os.Args[0]), usage)
}

func chkArgs() bool {
	if len(os.Args) == 1 {
		return true
	}

	switch os.Args[1] {
	case "-v":
		fmt.Printf("%s: %s\n", filepath.Base(os.Args[0]), version)
		return false
	case "-h":
		usage()
		return false
	default:
		return true
	}

}

func listIDs() {
	var api apirequest.APIData
	api.API = "/api/account/"

	if err := api.APIRequest(); err != nil {
		fmt.Println(err)
	}

	var dat sl1UserLst

	if err := json.Unmarshal(api.Result, &dat); err != nil {
		fmt.Println(err)
	}

	if len(os.Args) == 1 {
		for _, u := range dat {
			fmt.Printf("sl1id=%s(%s)\n", filepath.Base(u.URI), u.Description)
		}
		return
	}

	if len(os.Args) > 1 {
		for _, u := range dat {
			for _, a := range os.Args {
				if u.Description == a {
					fmt.Printf("sl1id=%s(%s)\n", filepath.Base(u.URI), u.Description)
				}
			}
		}
		return
	}

}

func main() {

	if chkArgs() {
		listIDs()
	}

}
