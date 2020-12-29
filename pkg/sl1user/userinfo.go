/*
Copyright © 2020 Marco Ostaska

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package sl1user

import (
	"fmt"
	"path"

	"github.com/marco-ostaska/sl1cmd/pkg/httpcalls"
	"github.com/marco-ostaska/sl1cmd/pkg/sl1generics"
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
	PermissionKeys       []string      `json:"permission_keys"`
	sl1id                string
}

// AccessHooks subset of UserDetails
type AccessHooks struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

// LoadUserDetails get user details from /api/account/x
// and load it to *UserDetails
func (ud *UserDetails) LoadUserDetails(id string) error {
	ud.sl1id = id
	var api httpcalls.APIData
	api.API = "/api/account/"
	api.ARGS = id
	err := api.NewRequest(&ud)
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
	fmt.Println("sl1id                 :", ud.sl1id)
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

// GetUserDetails checks if url is reachable
func (ud *UserDetails) GetUserDetails(username string) error {
	var usr UserAcct
	if err := usr.GetIDs(); err != nil {
		return err
	}

	id, err := usr.Sl1UserID(username)
	if err != nil {
		return err
	}

	if err := ud.LoadUserDetails(id); err != nil {
		return err
	}

	return nil
}
