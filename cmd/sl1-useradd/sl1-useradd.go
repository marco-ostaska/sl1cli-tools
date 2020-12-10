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

	u = userAdd{
		Organization:        "/api/organization/2",
		User:                "teste",
		Email:               "tete@ddsds.com",
		ContactFname:        "Teste",
		ContactLname:        "dsdsdsd",
		PasswdResetRequired: "1",
		Admin:               "0",
	}

	u.UserPolicy = "/api/account_policy/11"
	u.AlignedOrganizations = strings.Split("/api/organization/0,/api/organization/2,/api/organization/3,/api/organization/4", ",")

	e, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	var p apipost.APIData
	p.API = "/api/account"

	p.Payload = strings.NewReader(string(e))

	if err := p.APIPost(); err != nil {
		log.Fatal(err)
	}

	var ua sl1user.UserDetails
	if err = json.Unmarshal(p.Result, &ua); err != nil {
		log.Fatal(err)
	}

	fmt.Println((ua.User))
}
