package main

import (
	"flag"
	"fmt"
	"os"
)

type flArgs struct {
	username *string
	passwd   *string
	url      *string
}

const version string = "%s: v1.0.0-2020-Dec\n"
const usage string = `Configure or update credentials for sl1api

Options:
  -new     setup a new configuration
  -update  update user and Password to existing cofiguration
  -h       display this help and exit
  -v       display version

Arguments:
  -u      username
  -p      password
  -url    sl1 api URL
`

const usageNewCfg string = `Configure new credentials for sl1api
This can also be used to update all parameters to existing vault

Mandatory Arguments: 
  -u      username
  -p      password
  -url    sl1 api URL

Example:
  %s -new -u "myuser" -p "pass1234" -url "https://sl1api.com"
`

const usageUpdateCfg string = `Update existing credentials for sl1api

Mandatory Arguments: 
  -u      username
  -p      password

Example:
  %s -update -u "myuser" -p "pass1234"
`

func (fl *flArgs) initFlag() error {
	newconfig := flag.Bool("new", false, "")
	update := flag.Bool("update", false, "")
	help := flag.Bool("h", false, "")
	v := flag.Bool("v", false, "")
	fl.username = flag.String("u", "", "")
	fl.passwd = flag.String("p", "", "")
	fl.url = flag.String("url", "", "")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [ARGUMENTS]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, usage)
	}

	if *v {
		fmt.Fprintf(os.Stderr, version, os.Args[0])
		return fmt.Errorf("dislay version")
	}

	if *newconfig && *help {
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [ARGUMENTS]\n", os.Args[0])
			fmt.Fprintf(os.Stderr, usageNewCfg, os.Args[0])
		}
		flag.Usage()
		return fmt.Errorf("missing arguments")
	}

	if *update && *help {
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [ARGUMENTS]\n", os.Args[0])
			fmt.Fprintf(os.Stderr, usageUpdateCfg, os.Args[0])
		}
		flag.Usage()
		return fmt.Errorf("missing arguments")
	}

	if !*newconfig && !*update || *help || *newconfig && *update {
		flag.Usage()
		return fmt.Errorf("missing arguments")
	}

	if *update {
		if *fl.username == "" || *fl.passwd == "" {
			flag.Usage = func() {
				fmt.Fprintf(os.Stderr, "Usage: %s -update ... [ARGUMENTS]\n", os.Args[0])
				fmt.Fprintf(os.Stderr, usageUpdateCfg, os.Args[0])
			}
			flag.Usage()
			return fmt.Errorf("missing arguments")
		}

		return fl.updateCFG()

	}

	if *newconfig {
		if *fl.username == "" || *fl.passwd == "" || *fl.url == "" {
			flag.Usage = func() {
				fmt.Fprintf(os.Stderr, "Usage: %s -new ... [ARGUMENTS]\n", os.Args[0])
				fmt.Fprintf(os.Stderr, usageNewCfg, os.Args[0])
			}
			flag.Usage()
			return fmt.Errorf("missing arguments")
		}
		return fl.newCFG()
	}

	return nil

}
