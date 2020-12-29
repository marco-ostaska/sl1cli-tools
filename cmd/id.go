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
	"github.com/marco-ostaska/sl1cmd/pkg/sl1generics"
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
		var err error
		httpcalls.Insecure, err = rootCmd.Flags().GetBool("insecure")
		if err != nil {
			log.Fatalln(err)
		}
		var bInfo sl1generics.BasicInfo
		if err := bInfo.Load("/api/account"); err != nil {
			log.Fatalln(err)
		}
		a := bInfo.ListBasic(args, cmd.Name(), "no such user")

		for _, v := range a {
			fmt.Println(v)
		}

	},
}

func init() {
	rootCmd.AddCommand(idCmd)
}
