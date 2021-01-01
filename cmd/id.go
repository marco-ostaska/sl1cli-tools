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
	"github.com/marco-ostaska/sl1cmd/pkg/sl1"
	"github.com/marco-ostaska/sl1cmd/pkg/sl1/httpcalls"
	"github.com/spf13/cobra"
)

// idCmd represents the id command
var idCmd = &cobra.Command{
	Use:           "id [user]...",
	Short:         "print users sl1 id for the specified user.",
	SilenceErrors: true,
	Long: `print users sl1 id for the specified user.
or (when USER omitted) prints a list of all users sl1 ids.
.`,
	RunE: id,
}

func id(cmd *cobra.Command, args []string) error {
	httpcalls.Insecure, _ = rootCmd.Flags().GetBool("insecure")
	var u sl1.BasicInfo
	if err := u.Load(sl1.AccountAPI); err != nil {
		return err
	}
	u.Println(args)
	return nil

}

func init() {
	rootCmd.AddCommand(idCmd)
}
