package vault_test

import (
	"fmt"

	"github.com/marco-ostaska/sl1cmd/pkg/sl1/vault"
)

func ExampleCredential_SetInfo() {
	var vCred vault.Credential

	if err := vCred.SetInfo("myUser", "myPass@#$%^&*", "https://sl1api/"); err != nil {
		fmt.Println(err)
	}

}

func ExampleCredential_ReadFile() {
	var vCred vault.Credential
	if err := vCred.ReadFile(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("User:", vCred.UserAPI)
	fmt.Println("Pass:", vCred.DcryptP)
	// Output:
	// User: myUser
	// Pass: myPass@#$%^&*

}
