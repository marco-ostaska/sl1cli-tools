package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

type userAdd struct {
	Organization         string   `json:"organization"`
	User                 string   `json:"user"`
	Email                string   `json:"email"`
	ContactFname         string   `json:"contact_fname"`
	ContactLname         string   `json:"contact_lname"`
	PasswdResetRequired  string   `json:"passwd_reset_required"`
	Admin                string   `json:"admin"`
	UserPolicy           string   `json:"user_policy"`
	PermissionKeys       []string `json:"permission_keys"`
	AlignedOrganizations []string `json:"aligned_organizations"`
}

const version string = "%s: v1.0.0-2020-Dec\n"
const usage string = `Create newuser on sl1

Options:
  -org                   User Organization ID
  -email                 User email  
  -name                  User full name
  -resetrequired         Password required to be changed on first login (0 or 1)
  -admin                 Admin 0 or 1
  -userpolicy            User Policy ID
  -permissionkeys        Permission Keys IDs separated by comma
  -alignedorgs           Aligned Organizations IDs separated by comma
  -h                     display this help and exit
  -v                     display version

Warning:
  Avoid using double quotes, use always single quotes on arguments

Example:
 %s -org '2' -ermail 'teste@xx.com' -name 'teste' -admin '1' -userpolicy '3' -alignedorgs '0,2,3,5'
`

func reqFlags(s ...*string) bool {
	for _, v := range s {
		if *v == "" {
			return false
		}
	}

	return true
}

func (fl *userAdd) initFlag() error {
	org := flag.String("org", "", "")
	email := flag.String("email", "", "")
	name := flag.String("name", "", "")
	// resetRequired := flag.String("resetRequired", "0", "")
	// admin := flag.String("admin", "0", "")
	userpolicy := flag.String("userpolicy", "", "")
	// alignedorgs := flag.String("alignedorgs", "", "")
	// permissionKeys := flag.String("permissionKeys", "", "")
	help := flag.Bool("h", false, "")
	v := flag.Bool("v", false, "")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [ARGUMENTS]\n", path.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, usage, path.Base(os.Args[0]))
	}

	if *v {
		fmt.Fprintf(os.Stderr, version, path.Base(os.Args[0]))
		return fmt.Errorf("dislay version")
	}

	if *help {
		flag.Usage()
		return fmt.Errorf("help")
	}

	if !reqFlags(org, email, name, userpolicy) {
		flag.Usage()
		return fmt.Errorf("missing arguments")
	}

	return nil

}
