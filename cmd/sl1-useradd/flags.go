package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

const version string = "%s: v1.0.0-2020-Dec\n"
const usage string = `Create newuser on sl1

Mandatory Options:
  -user                  User name
  -email                 User email 
  -name                  User full name
  -org                   User Organization ID
  -userpolicy            User Policy ID

Options:
  -resetrequired         Password required to be changed on first login (0 or 1) (Default: 1)
  -admin                 Admin 0 or 1                                            (Default: 1)
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

func (ua *userAdd) nameFMT(name string) {
	a := strings.Split(name, " ")
	switch {
	case len(a) == 2:
		ua.ContactFname = a[0]
		ua.ContactLname = a[1]
	case len(a) > 2:
		ua.ContactFname = strings.Join(a[:2], " ")
		ua.ContactLname = strings.Join(a[2:], " ")
	default:
		ua.ContactFname = name
	}

}

func buildSlice(s *string, api string) []string {
	if reqFlags(s) {
		arr := strings.Split(*s, ",")

		for i := 0; i < len(arr); i++ {
			arr[i] = api + arr[i]
		}

		return arr
	}
	var arr []string
	return arr

}

func (ua *userAdd) initFlag() error {
	org := flag.String("org", "", "")
	user := flag.String("user", "", "")
	email := flag.String("email", "", "")
	name := flag.String("name", "", "")
	resetRequired := flag.String("resetRequired", "0", "")
	admin := flag.String("admin", "0", "")
	userpolicy := flag.String("userpolicy", "", "")
	alignedorgs := flag.String("alignedorgs", "", "")
	permissionKeys := flag.String("permissionKeys", "", "")
	help := flag.Bool("h", false, "")
	v := flag.Bool("v", false, "")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]...\n", path.Base(os.Args[0]))
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

	if !reqFlags(org, email, name, userpolicy, user) {
		flag.Usage()
		return fmt.Errorf("missing arguments")
	}

	ua.Organization = *org
	ua.User = *user
	ua.Email = *email
	ua.nameFMT(*name)
	ua.PasswdResetRequired = *resetRequired
	ua.Admin = *admin
	ua.UserPolicy = "/api/account_policy/" + *userpolicy
	ua.AlignedOrganizations = buildSlice(alignedorgs, "/api/organization/")
	ua.PermissionKeys = buildSlice(permissionKeys, "/api/permission_key/")
	return nil

}
