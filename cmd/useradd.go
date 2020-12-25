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
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/marco-ostaska/sl1cmd/pkg/httpcalls"
	"github.com/marco-ostaska/sl1cmd/pkg/sl1user"
	"github.com/spf13/cobra"
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

// useraddCmd represents the useradd command
var useraddCmd = &cobra.Command{
	Use:   "useradd",
	Short: "create a new user.",
	Long: `create newuser on sl1

Warning:
  User configured in sl1cmd vault must be administrator to execute this task.`,
	Example: `sl1cmd useradd --user 'myuser' -o '2' --email 'teste@xx.com' --name 'My Name' --admin 1 --userpolicy 3 -g 0 -g 2 -g 3`,
	Run: func(cmd *cobra.Command, args []string) {
		var u userAdd
		u.parseFlags(cmd)

		j, err := json.Marshal(u)
		if err != nil {
			log.Fatal(err)
		}

		var p httpcalls.APIData

		p.Payload = strings.NewReader(string(j))

		fmt.Println(p.Payload)

		if err := p.NewPost(&j, "/api/account"); err != nil {
			log.Fatal(err)
		}

		var ua sl1user.UserDetails
		if err = json.Unmarshal(p.Result, &ua); err != nil {
			log.Fatal("abc:", err)
		}

		if ua.User == u.User {
			fmt.Println(u.User, "created successful")
			return
		}
		fmt.Println(string(p.Result))

	},
}

func parseStringFlag(cmd *cobra.Command, s string) string {
	sFlag, err := cmd.Flags().GetString(s)
	if err != nil {
		log.Fatalln(err)
	}

	return sFlag
}

func parseStringArrFlag(cmd *cobra.Command, s string) []string {
	sFlag, err := cmd.Flags().GetStringArray(s)
	if err != nil {
		log.Fatalln(err)
	}

	return sFlag
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

func (ua *userAdd) parseFlags(cmd *cobra.Command) {

	ua.User = parseStringFlag(cmd, "user")
	ua.Email = parseStringFlag(cmd, "email")
	ua.Organization = parseStringFlag(cmd, "organization")
	ua.UserPolicy = parseStringFlag(cmd, "userpolicy")
	ua.PasswdResetRequired = parseStringFlag(cmd, "resetrequired")
	ua.Admin = parseStringFlag(cmd, "admin")
	ua.PermissionKeys = parseStringArrFlag(cmd, "permissionkeys")
	ua.AlignedOrganizations = parseStringArrFlag(cmd, "alignedorgs")

	ua.nameFMT(parseStringFlag(cmd, "name"))

}

func init() {
	rootCmd.AddCommand(useraddCmd)
	useraddCmd.Flags().StringP("user", "u", "", "username")

	useraddCmd.Flags().StringP("email", "e", "", "email")
	useraddCmd.Flags().StringP("name", "n", "", "user full name")
	useraddCmd.Flags().StringP("organization", "o", "", "organization id")
	useraddCmd.Flags().StringP("userpolicy", "p", "", "User Policy id")

	useraddCmd.MarkFlagRequired("user")
	useraddCmd.MarkFlagRequired("email")
	useraddCmd.MarkFlagRequired("name")
	useraddCmd.MarkFlagRequired("organization")
	useraddCmd.MarkFlagRequired("userpolicy")

	useraddCmd.Flags().StringP("resetrequired", "r", "1", "Password required to be changed on first login")
	useraddCmd.Flags().StringP("admin", "a", "1", "should be admin?")
	useraddCmd.Flags().StringArrayP("permissionkeys", "k", nil, "Permission keys IDs")
	useraddCmd.Flags().StringArrayP("alignedorgs", "g", nil, "Aligned Organizations IDs")

}
