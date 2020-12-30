package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/marco-ostaska/sl1cmd/pkg/wrapper"
)

func TestNewVault(t *testing.T) {
	user := "em7admin"
	passwd := "em7admin"
	uri := "https://sl1.lab"
	if err := vCredential.SetInfo(user, passwd, uri); err != nil {
		t.Errorf(err.Error())
	}

	if err := vCredential.ReadFile(); err != nil {
		t.Errorf(err.Error())
	}

	switch {
	case vCredential.UserAPI != user:
		t.Errorf("got %v, expected %v", vCredential.UserAPI, user)
	case vCredential.UserAPI != user:
		t.Errorf("got %v, expected %v", vCredential.DcryptP, passwd)
	case vCredential.URL != uri:
		t.Errorf("got %v, expected %v", vCredential.URL, uri)
	}

}

func TestUpdateVault(t *testing.T) {
	err := vCredential.ReadFile()
	if err != nil {
		t.Errorf(err.Error())

	}

	user := "em8admin"
	passwd := "em8admin"

	if err = vCredential.SetInfo(user, passwd, vCredential.URL); err != nil {
		t.Errorf(err.Error())
	}

	if err := vCredential.ReadFile(); err != nil {
		t.Errorf(err.Error())
	}

	switch {
	case vCredential.UserAPI != user:
		t.Errorf("got %v, expected %v", vCredential.UserAPI, user)
	case vCredential.UserAPI != user:
		t.Errorf("got %v, expected %v", vCredential.DcryptP, passwd)
	}
}

func TestDeleteVault(t *testing.T) {

	err := vCredential.UserInfo()
	err1 := os.Remove(vCredential.File)

	if re := wrapper.ReturnError(err, err1); re != nil {
		switch {
		case strings.Contains(err.Error(), vCredential.File+": no such file or directory"):
			return
		default:
			t.Errorf(err.Error())
		}

	}

}
