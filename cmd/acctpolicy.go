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
	"github.com/marco-ostaska/sl1cmd/pkg/sl1generics"
	"github.com/spf13/cobra"
)

// accountpolicy represents the id command
var acctpolicy = &cobra.Command{
	Use:   "acctpolicy",
	Short: "print users sl1 account policies.",
	Long: `print users sl1 id for the specified user.
or (when USER omitted) prints a list of all users sl1 ids.
.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpcalls.Insecure, _ = rootCmd.Flags().GetBool("insecure")
		var bGen sl1generics.Basic
		if err := bGen.GetIDs("account_policy"); err != nil {
			log.Fatalln(err)
		}
		bGen.PrintBasic(args)

	},
}

func init() {
	rootCmd.AddCommand(acctpolicy)
}
