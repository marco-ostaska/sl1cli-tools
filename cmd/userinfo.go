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

	"github.com/marco-ostaska/sl1cmd/pkg/httpcalls"
	"github.com/marco-ostaska/sl1cmd/pkg/sl1user"
	"github.com/spf13/cobra"
)

// userinfoCmd represents the userinfo command
var userinfoCmd = &cobra.Command{
	Use:   "userinfo [user]",
	Short: "print user information for the specified user.",
	Long:  `print user information for the specified user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		httpcalls.Insecure, _ = rootCmd.Flags().GetBool("insecure")
		var ud sl1user.UserDetails
		if len(args) != 1 {
			fmt.Printf("%v \n", cmd.Short)
			if err := cmd.Usage(); err != nil {
				log.Fatalln(err)
			}
			return fmt.Errorf("wrong number of user names passed")
		}
		if err := ud.GetUserDetails(args[0]); err != nil {
			fmt.Println(err)
			return nil
		}
		ud.PrintUserDetails()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(userinfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userinfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userinfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
