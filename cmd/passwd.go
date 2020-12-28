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
	"syscall"

	"github.com/marco-ostaska/sl1cmd/pkg/httpcalls"
	"github.com/marco-ostaska/sl1cmd/pkg/sl1user"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// passwdCmd represents the passwd command
var passwdCmd = &cobra.Command{
	Use:   "passwd [user]",
	Short: "change user password.",
	Long: `change user password.
	
Warning:
  User configured in sl1cmd vault must be administrator to execute this task.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpcalls.Insecure, _ = rootCmd.Flags().GetBool("insecure")
		if len(args) != 1 {
			fmt.Println(cmd.Short)
			if err := cmd.Usage(); err != nil {
				log.Fatalln(err)
			}
			fmt.Println("invalid number of users passed, it should be 1 user")
			return
		}
		passwd(args[0])
	},
}

func validatePasswd(s string) (string, error) {

	fmt.Println("Changing password for user", s)
	fmt.Printf("New sl1 password: ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	fmt.Printf("\nRetype new sl1 password: ")

	bytePasswordR, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	if string(bytePassword) != string(bytePasswordR) {
		return "", fmt.Errorf("Password do not mach ❌")
	}
	return string(bytePassword), nil
}

func passwd(user string) {
	var usr sl1user.UserAcct
	if err := usr.GetIDs(); err != nil {
		fmt.Println(err)
		return
	}
	i, err := usr.ID(user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "id: %s no such user\n", os.Args[1])
		os.Exit(1)
	}

	sPasswd, err := validatePasswd((usr)[i].Description)
	if err != nil {
		fmt.Printf("\n%v \n", err)
		os.Exit(2)
	}

	var p httpcalls.APIData
	p.ARGS = "/password"
	payload := fmt.Sprintf(`{"password": "%s"}`, sPasswd)

	p.Payload = strings.NewReader(payload)

	if err := p.NewPost(nil, (usr)[i].URI); err != nil {
		log.Fatal(err)
	}

	if string(p.Result) == "" {
		fmt.Println()
		fmt.Println("password updated successfully")
		return
	}
	fmt.Println()
	fmt.Println(string(p.Result))
}

func init() {
	rootCmd.AddCommand(passwdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
