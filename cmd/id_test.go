package cmd

import (
	"testing"
)

func TestId(t *testing.T) {
	tt := []struct {
		name string
		args []string
	}{
		{"Listall", []string{"id", "--insecure"}},
		{"User", []string{"id", "em7admin", "--insecure"}},
		{"No such account", []string{"id", "em7admin", "teste1234", "em7admin", "--insecure"}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rootCmd.SetArgs(tc.args)
			rootCmd.SilenceErrors = true
			rootCmd.SilenceUsage = true
			if err := rootCmd.Execute(); err != nil {
				t.Errorf(err.Error())
			}
		})
	}

}
