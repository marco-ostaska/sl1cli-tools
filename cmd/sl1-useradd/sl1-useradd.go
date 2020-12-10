package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/marco-ostaska/sl1cli-tools/pkg/apipost"
	"github.com/marco-ostaska/sl1cli-tools/pkg/sl1user"
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

func main() {

	var u userAdd

	if err := u.initFlag(); err != nil {
		log.Fatal(err)
	}

	j, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	var p apipost.APIData
	p.API = "/api/account"

	p.Payload = strings.NewReader(string(j))

	if err := p.APIPost(); err != nil {
		log.Fatal(err)
	}

	var ua sl1user.UserDetails
	if err = json.Unmarshal(p.Result, &ua); err != nil {
		log.Fatal(err)
	}

	if ua.User == u.User {
		fmt.Println(u.User, "created successful")
		return
	}
	fmt.Println(string(p.Result))
}
