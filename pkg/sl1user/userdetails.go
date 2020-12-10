package sl1user

import (
	"fmt"
	"path"

	"github.com/marco-ostaska/sl1cli-tools/pkg/apirequest"
	"github.com/marco-ostaska/sl1cli-tools/pkg/sl1generics"
)

// UserDetails is an abstraction of /api/account/x
type UserDetails struct {
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

// AccessHooks subset of UserDetails
type AccessHooks struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

// LoadUserDetails get user details from /api/account/x
// and load it to *UserDetails
func (ud *UserDetails) LoadUserDetails(id string) error {
	var api apirequest.APIData
	err := api.NewRequest(&ud, "/api/account/", id)
	if err != nil {
		return err
	}
	return nil
}

// PrintUserDetails prints user details from /api/account/x
func (ud *UserDetails) PrintUserDetails() {
	var ua UserAcct
	if err := ua.GetIDs(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("User                  :", ud.User)
	fmt.Println("First Name            :", ud.ContactFname)
	fmt.Println("Last Name             :", ud.ContactLname)
	fmt.Println("Email                 :", ud.Email)
	fmt.Println("Organization          :", ud.Organization)
	fmt.Println("PasswdExpiration      :", ud.PasswdExpiration)
	t, err := sl1generics.EpochToUnix(ud.PasswdSetDate)
	if err == nil {
		fmt.Println("PasswdSetDate         :", t)
	}
	fmt.Println("PasswdPrevCount       :", ud.PasswdPrevCount)
	fmt.Println("PasswdResetRequired   :", ud.PasswdResetRequired)
	fmt.Println("PasswdStrength        :", ud.PasswdStrength)
	fmt.Println("LoginState            :", ud.LoginState)
	fmt.Println("RestrictIP            :", ud.RestrictIP)
	fmt.Println("Admin                 :", ud.Admin)
	t, err = sl1generics.EpochToUnix(ud.CreateDate)
	if err == nil {
		fmt.Println("CreateDate            :", t)
	}

	userID, err := ua.SearchByURI(ud.CreatedBy)
	if err == nil {
		fmt.Printf("CreatedBy             : sl1id=%s(%s)\n", path.Base(ud.CreatedBy), userID)
	}
	t, err = sl1generics.EpochToUnix(ud.EditDate)
	if err == nil {
		fmt.Println("EditDate              :", t)
	}
	userID, err = ua.SearchByURI(ud.UpdatedBy)
	if err == nil {
		fmt.Printf("UpdatedBy             : sl1id=%s(%s)\n", path.Base(ud.UpdatedBy), userID)
	}

	fmt.Println("UserPolicy            :", ud.UserPolicy)
	fmt.Println("AlignedOrganizations  :", ud.AlignedOrganizations)
	fmt.Println("AccessHooks           :", ud.AccessHooks)

}
