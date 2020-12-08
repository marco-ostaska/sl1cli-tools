package main

import (
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/marco-ostaska/sl1cli-tools/pkg/apirequest.go"
)

const version string = "v1.0.0 dec-2020"

type sl1UserDetails struct {
	Organization         string        `json:"organization"`
	User                 string        `json:"user"`
	Email                string        `json:"email"`
	PasswdExpiration     string        `json:"passwd_expiration"`
	PasswdSetDate        string        `json:"passwd_set_date"`
	PasswdPrevCount      string        `json:"passwd_prev_count"`
	PasswdResetRequired  string        `json:"passwd_reset_required"`
	PasswdStrength       string        `json:"passwd_strength"`
	LoginState           string        `json:"login_state"`
	RestrictIP           string        `json:"restrict_ip"`
	Admin                string        `json:"admin"`
	Active               string        `json:"active"`
	CreateDate           string        `json:"create_date"`
	CreatedBy            string        `json:"created_by"`
	EditDate             string        `json:"edit_date"`
	UpdatedBy            string        `json:"updated_by"`
	UserPolicy           string        `json:"user_policy"`
	Country              string        `json:"country"`
	ContactFname         string        `json:"contact_fname"`
	ContactLname         string        `json:"contact_lname"`
	AlignedOrganizations []string      `json:"aligned_organizations"`
	AlignedTicketQueues  []interface{} `json:"aligned_ticket_queues"`
	AccessHooks          AccessHooks   `json:"access_hooks"`
}
type AccessHooks struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

type userStruct interface {
	getUsers()
}

func (dat *sl1UserLst) findUser(uri string) (string, error) {
	for _, u := range *dat {
		if u.URI == uri {
			return filepath.Base(u.URI) + u.Description, nil
		}
	}
	return "", fmt.Errorf("user not found")
}

// need to create a func to avoid copy and paste
func getDetails(id string) error {
	var api apirequest.APIData
	api.API = "/api/account/" + id

	if err := api.APIRequest(); err != nil {
		return err
	}

	var dat sl1UserDetails

	if err := json.Unmarshal(api.Result, &dat); err != nil {
		return err
	}

	fmt.Println("User                :", dat.User)
	fmt.Println("Email               :", dat.Email)
	fmt.Println("Organization        :", path.Base(dat.Organization))
	fmt.Println("PasswdExpiration    :", path.Base(dat.PasswdExpiration))

	//cb, err := dat.findUser()
	fmt.Println("CreatedBy           :", path.Base(dat.CreatedBy))

	i, _ := strconv.ParseInt(dat.PasswdSetDate, 10, 64)
	t := time.Unix(i, 0)
	fmt.Println("PasswdSetDate       :", t)

	// PasswdPrevCount     : 5
	// PasswdResetRequired : 0
	// PasswdStrength      : 0
	// LoginState          : 1
	// RestrictIP          :
	// Admin               : 0
	// Active              : 0
	// CreateDate          : 1603145206
	// CreatedBy           : /api/account/4
	// EditDate            : 1607101184
	// UpdatedBy           : /api/account/2
	// UserPolicy          : /api/account_policy/11
	// AllOrgs             : 1
	// ContactFname        : Marco Testetetete
	// ContactLname        : Aurelio Najar Ostaka
	// AlignedOrganizations: [/api/organization/0 /api/organization/2 /api/organization/3 /api/organization/4]

	return nil
}

func main() {

	if chkArgs() {

		//listIDs()
		var su sl1UserLst
		su.getUsers()
		cb, err := su.findUser("/api/account/4")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cb)
	}
}
