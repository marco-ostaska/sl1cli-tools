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
	"strings"

	"github.com/marco-ostaska/sl1cmd/pkg/sl1/vault"
	"github.com/marco-ostaska/sl1cmd/pkg/wrapper"
	"github.com/spf13/cobra"
)

var vCredential vault.Credential

// vaultCmd represents the vault command
var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "create or update login information vault for api.",
	Long:  `create or update login information vault for api.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v\n", cmd.Long)
		if err := cmd.Usage(); err != nil {
			log.Fatalln(err)
		}
	},
}

var newCmd = &cobra.Command{
	Use:           "new",
	Short:         "create new vault.",
	Long:          `create new vault.`,
	SilenceErrors: true,
	Example: `
  Unix Based OS: (use single quotes)
      sl1cmd vault new -u 'myuser' -p 'pass1234' --url 'https://sl1api.com'
  Windows: (use double quotes)
      sl1cmd vault new -u "myuser" -p "pass1234" --url "https://sl1api.com"
`,
	RunE: newVault,
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing vault.",
	Long:  `update an existing vault.`,
	Example: `  
  Unix based OS:  (use single quotes)
      sl1cmd update -u 'myuser' -p 'pass1234'
  Windows: (use double quotes)
      sl1cmd update -u "myuser" -p "pass1234"`,
	RunE: updateVault,
}

var deleteCmd = &cobra.Command{
	Use:           "delete",
	Short:         "delete an existing vault.",
	Long:          `delete an existing vault.`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE:          deleteVault,
}

func addCommandUpdateCmd() error {
	vaultCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("user", "u", "", "username")
	updateCmd.Flags().StringP("password", "p", "", "password")

	err := updateCmd.MarkFlagRequired("user")
	err1 := updateCmd.MarkFlagRequired("password")

	return wrapper.ReturnError(err, err1)

}

func addCommandNewCmd() error {
	vaultCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("user", "u", "", "username")
	newCmd.Flags().StringP("password", "p", "", "password")
	newCmd.Flags().String("url", "", "API URI")

	err := newCmd.MarkFlagRequired("user")
	err1 := newCmd.MarkFlagRequired("password")
	err2 := newCmd.MarkFlagRequired("url")

	if re := wrapper.ReturnError(err, err1, err2); re != nil {
		return re
	}

	return nil

}

func newVault(cmd *cobra.Command, args []string) error {
	user, err := cmd.Flags().GetString("user")
	passwd, err1 := cmd.Flags().GetString("password")
	uri, err2 := cmd.Flags().GetString("url")

	if re := wrapper.ReturnError(err, err1, err2); re != nil {
		return re
	}

	if err = vCredential.SetInfo(user, passwd, uri); err != nil {
		return err
	}

	fmt.Println("Vault configured ✔")
	return nil
}

func updateVault(cmd *cobra.Command, args []string) error {
	err := vCredential.ReadFile()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return fmt.Errorf("No credentials found, please try create a new credential vault first ❌")
		}
		return err
	}
	user, err := cmd.Flags().GetString("user")
	passwd, err1 := cmd.Flags().GetString("password")
	if re := wrapper.ReturnError(err, err1); re != nil {
		return err
	}

	fmt.Println("Updating credentials for", vCredential.URL, "user", vCredential.UserAPI)

	if err = vCredential.SetInfo(user, passwd, vCredential.URL); err != nil {
		return err
	}
	fmt.Println("Vault configured ✔")
	return nil

}

func deleteVault(cmd *cobra.Command, args []string) error {
	if err := vCredential.UserInfo(); err != nil {
		return fmt.Errorf("%s ❌", err)
	}

	if err := os.Remove(vCredential.File); err != nil {
		return fmt.Errorf("%s ❌", err.Error())
	}

	fmt.Println("Vault deleted ✔")
	return nil

}

func init() {
	rootCmd.AddCommand(vaultCmd)
	vaultCmd.AddCommand(deleteCmd)

	err := addCommandNewCmd()
	err1 := addCommandUpdateCmd()

	if re := wrapper.ReturnError(err, err1); re != nil {
		log.Fatalln(err)
	}

}
