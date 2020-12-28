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

// Package cmd is default pkg used for cobra
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "sl1cmd",
	Short:   "sl1cmd is a command line interface to interact with ScienceLogic Monitoring tool API.",
	Long:    `sl1cmd is a command line interface to interact with ScienceLogic Monitoring tool API.`,
	Version: "Unreleased Dec/2020",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//doc.GenMarkdownTree(rootCmd, "./docs/cmd")

}

func init() {
	rootCmd.PersistentFlags().Bool("insecure", false, "accept invalid certificates.")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("version", "v", false, "output version information and exit")
}
