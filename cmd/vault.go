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
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/marco-ostaska/sl1cmd/pkg/cryptcfg"
	"github.com/spf13/cobra"
)

var usr cryptcfg.UserInfo

// vaultCmd represents the vault command
var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "create or update login information vault for api.",
	Long:  `create or update login information vault for api.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v\n", cmd.Short)
		if err := cmd.Usage(); err != nil {
			log.Fatalln(err)
		}
	},
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create new vault.",
	Long: `create new vault.

Information:
  You may experience login issues using double quotes, use single quotes instead.
  `,
	Example: `sl1cmd new -u 'myuser' -p 'pass1234' --url 'https://sl1api.com'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := cmd.Flags().GetString("user")
		passwd, err1 := cmd.Flags().GetString("password")
		uri, err2 := cmd.Flags().GetString("url")

		if err != nil || err1 != nil || err2 != nil {
			return err
		}

		if err = usr.SetInfo(user, passwd, uri); err != nil {
			return err
		}

		fmt.Println("Vault configured ✔")
		return nil
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing vault.",
	Long: `update an existing vault.

Information:
  You may experience login issues using double quotes, use single quotes instead.
  `,
	Example: `sl1cmd update -u 'myuser' -p 'pass1234'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := usr.ReadCryptFile()
		if err != nil {
			eStr := fmt.Sprintf("%v", err)
			if strings.Contains(eStr, "no such file or directory") {
				fmt.Println("No credentials found, please try create a new credential vault first ❌")
				return err
			}
			fmt.Println(err)
			return err
		}

		fmt.Println("Updating credentials for", usr.URL, "user", usr.UserAPI)
		user, err := cmd.Flags().GetString("user")
		passwd, err1 := cmd.Flags().GetString("password")
		if err != nil || err1 != nil {
			return err
		}

		if err = usr.SetInfo(user, passwd, usr.URL); err != nil {
			return err
		}
		fmt.Println("Vault configured ✔")
		return nil

	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an existing vault.",
	Long:  `delete an existing vault.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		usr, err := user.Current()
		if err != nil {
			return err
		}

		if err = os.Remove(usr.HomeDir + "/.local/sl1api/" + "sl1api.cfg"); err != nil {
			return err
		}

		fmt.Println("Vault deleted ✔")
		return nil

	},
}

func addFlag() {
	vaultCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("user", "u", "", "username")
	newCmd.Flags().StringP("password", "p", "", "password")
	newCmd.Flags().String("url", "", "API URI")

	if err := newCmd.MarkFlagRequired("user"); err != nil {
		log.Fatalln(err)
	}
	if err := newCmd.MarkFlagRequired("password"); err != nil {
		log.Fatalln(err)
	}
	if err := newCmd.MarkFlagRequired("url"); err != nil {
		log.Fatalln(err)
	}

}

func updateFlag() {
	vaultCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("user", "u", "", "username")
	updateCmd.Flags().StringP("password", "p", "", "password")

	if err := updateCmd.MarkFlagRequired("user"); err != nil {
		log.Fatalln(err)
	}

	if err := updateCmd.MarkFlagRequired("password"); err != nil {
		log.Fatalln(err)
	}

}

func init() {
	rootCmd.AddCommand(vaultCmd)
	addFlag()
	updateFlag()
	vaultCmd.AddCommand(deleteCmd)
}
