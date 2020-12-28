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
	"log"

	"github.com/marco-ostaska/sl1cmd/pkg/httpcalls"
	"github.com/marco-ostaska/sl1cmd/pkg/sl1user"
	"github.com/spf13/cobra"
)

// idCmd represents the id command
var idCmd = &cobra.Command{
	Use:   "id [user]...",
	Short: "print users sl1 id for the specified user.",
	Long: `print users sl1 id for the specified user.
or (when USER omitted) prints a list of all users sl1 ids.
.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpcalls.Insecure, _ = rootCmd.Flags().GetBool("insecure")
		var usr sl1user.UserAcct
		if err := usr.GetIDs(); err != nil {
			log.Fatalln(err)
		}
		usr.PrintID(args)

	},
}

func init() {
	rootCmd.AddCommand(idCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// idCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// idCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
