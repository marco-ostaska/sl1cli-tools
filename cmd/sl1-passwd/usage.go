package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const version string = "v1.0.0 dec-2020"

func usage() {
	usage := `Change password for the specified USER.
[WARNING] User configured by sl1-vault must be administrator to execute this task.

  -v             version
  -h             display this and exit
`
	fmt.Fprintf(os.Stderr, "Usage: %s passwd [USER]\n%v", filepath.Base(os.Args[0]), usage)
}

func chkArgs() bool {
	if len(os.Args) == 1 {
		usage()
		return false
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
