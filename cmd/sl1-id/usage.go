package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const version string = "v1.0.0 dec-2020"

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
