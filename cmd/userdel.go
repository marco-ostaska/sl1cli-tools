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
	"path"

	"github.com/marco-ostaska/sl1cmd/pkg/httpcalls"
	"github.com/marco-ostaska/sl1cmd/pkg/sl1user"
	"github.com/spf13/cobra"
)

// userdelCmd represents the userdel command
var userdelCmd = &cobra.Command{
	Use:   "userdel [user]",
	Short: "delete user account.",
	Long: `Delete the specified USER.

Warning:
  User configured in sl1cmd vault must be administrator to execute this task.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) != 1 {
			return fmt.Errorf("wrong number of users passed, it should be 1 user")
		}

		var usr sl1user.UserAcct
		if err := usr.GetIDs(); err != nil {
			return err
		}
		i, err := usr.ID(args[0])
		if err != nil {
			fmt.Printf("id: %s no such user\n", path.Base(args[0]))
			return nil
		}

		fmt.Println("deleting:", (usr)[i].Description)

		var p httpcalls.APIData
		p.API = (usr)[i].URI

		if err := p.DeleteRequest(); err != nil {
			return err
		}

		if string(p.Result) == "" {
			fmt.Println((usr)[i].Description, "deleted successfully")
			return nil
		}
		fmt.Println()
		fmt.Println(string(p.Result))
		return nil

	},
}

func init() {
	rootCmd.AddCommand(userdelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userdelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userdelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
