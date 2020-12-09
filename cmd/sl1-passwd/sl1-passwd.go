package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/marco-ostaska/sl1cli-tools/pkg/apipost.go"
	"github.com/marco-ostaska/sl1cli-tools/pkg/sl1user"
	"golang.org/x/crypto/ssh/terminal"
)

func passwd(s string) (string, error) {

	fmt.Println("Changing password for user", s)
	fmt.Printf("New sl1 password: ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	fmt.Printf("\nRetype new sl1 password: ")

	bytePasswordR, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	if string(bytePassword) != string(bytePasswordR) {
		return "", fmt.Errorf("Password do not mach")
	}
	return string(bytePassword), nil
}

func main() {
	if chkArgs() {
		if chkArgs() {
			var usr sl1user.UserAcct
			if err := usr.GetIDs(); err != nil {
				fmt.Println(err)
				return
			}
			i, err := usr.ID(os.Args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "id: %s no such user\n", os.Args[1])
				os.Exit(1)
			}

			sPasswd, err := passwd((usr)[i].Description)
			if err != nil {
				fmt.Printf("\n%v \n", err)
				os.Exit(2)
			}

			var p apipost.APIData
			p.API = (usr)[i].URI
			p.ARGS = "/password"
			payload := fmt.Sprintf(`{"password": "%s"}`, sPasswd)

			p.Payload = strings.NewReader(payload)

			if err := p.APIPost(); err != nil {
				log.Fatal(err)
			}

			if string(p.Result) == "" {
				fmt.Println("password updated successfully")
				return
			}
			fmt.Println()
			fmt.Println(string(p.Result))

		}
	}

}
