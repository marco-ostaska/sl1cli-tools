package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func usage() {
	usage := `Print sl1 user information details for the specified USER,

  -v             version
  -h             display this and exit
`
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [USER]\n%v", filepath.Base(os.Args[0]), usage)
}

func chkArgs() bool {
	if len(os.Args) < 2 {
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
