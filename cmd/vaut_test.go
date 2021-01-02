package cmd

import (
	"testing"
)

func TestNewVault(t *testing.T) {
	tt := []struct {
		name     string
		args     []string
		expected string
	}{
		{"Missing args 1", []string{"vault", "new"}, `required flag(s) "password", "url", "user" not set`},
		{"Missing args 2", []string{"vault", "new", "-u", "em7admin"}, `required flag(s) "password", "url" not set`},
		{"Missing args 2", []string{"vault", "new", "-u", "em7admin", "-p", "em7admin"}, `required flag(s) "url" not set`},
		{"Ok", []string{"vault", "new", "-u", "em7admin", "-p", "em7admin!@#$%&&*", "--url", "https://sl1.lab"}, ""},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rootCmd.SetArgs(tc.args)
			rootCmd.SilenceErrors = true
			rootCmd.SilenceUsage = true
			if err := rootCmd.Execute(); err != nil {
				if err.Error() == tc.expected {
					t.Skip(err)
					return
				}
				t.Errorf(err.Error())
			}
		})
	}

	t.Run("Check credentials", func(t *testing.T) {
		user := "em7admin"
		passwd := "em7admin!@#$%&&*"
		uri := "https://sl1.lab"
		if err := vCredential.ReadFile(); err != nil {
			t.Errorf(err.Error())
		}

		switch {
		case vCredential.UserAPI != user:
			t.Errorf("got %v, expected %v", vCredential.UserAPI, user)
		case vCredential.DcryptP != passwd:
			t.Errorf("got %v, expected %v", vCredential.DcryptP, passwd)
		case vCredential.URL != uri:
			t.Errorf("got %v, expected %v", vCredential.URL, uri)
		}

	})

}

func TestUpdateVault(t *testing.T) {
	tt := []struct {
		name     string
		args     []string
		expected string
	}{
		{"Missing args 1", []string{"vault", "update"}, `required flag(s) "password", "user" not set`},
		{"Missing args 2", []string{"vault", "update", "-u", "em8admin"}, `required flag(s) "password" not set`},
		{"OK", []string{"vault", "update", "-u", "em8admin", "-p", "em7admin"}, `required flag(s) "url" not set`},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rootCmd.SetArgs(tc.args)
			rootCmd.SilenceErrors = true
			rootCmd.SilenceUsage = true
			if err := rootCmd.Execute(); err != nil {
				if err.Error() == tc.expected {
					t.Skip(err)
					return
				}
				t.Errorf(err.Error())
			}
		})
	}

	t.Run("Check credentials", func(t *testing.T) {
		user := "em8admin"
		passwd := "em7admin"
		uri := "https://sl1.lab"
		if err := vCredential.ReadFile(); err != nil {
			t.Errorf(err.Error())
		}

		switch {
		case vCredential.UserAPI != user:
			t.Errorf("got %v, expected %v", vCredential.UserAPI, user)
		case vCredential.DcryptP != passwd:
			t.Errorf("got %v, expected %v", vCredential.DcryptP, passwd)
		case vCredential.URL != uri:
			t.Errorf("got %v, expected %v", vCredential.URL, uri)
		}

	})
}

func TestDeleteVault(t *testing.T) {

	rootCmd.SetArgs([]string{"vault", "delete"})
	if err := rootCmd.Execute(); err != nil {
		t.Errorf(err.Error())
	}

}
