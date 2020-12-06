package main

import (
	"fmt"
	"strings"

	"github.com/marco-ostaska/sl1cli-tools/internal/pkg/apicryptcfg"
)

var usr apicryptcfg.UserInfo

func (fl *flArgs) newCFG() error {
	return usr.SetInfo(*fl.username, *fl.passwd, *fl.url)
}

func (fl *flArgs) updateCFG() error {
	err := usr.ReadCryptFile()
	if err != nil {
		eStr := fmt.Sprintf("%v", err)
		if strings.Contains(eStr, "no such file or directory") {
			fmt.Println("No credentials found, please try create a new credential vault first")
			return err
		}
		fmt.Println(err)
		return err
	}

	fmt.Println("Updating credentials for", usr.URL, "user", usr.UserAPI)
	return usr.SetInfo(*fl.username, *fl.passwd, usr.URL)
}

func main() {
	var fl flArgs
	err := fl.initFlag()
	if err != nil {
		return
	}
}

// var u apicryptcfg.UserInfo
// err := u.SetInfo("ibmwpool", "K$V&ik3Hz@kM", "https://ibmwhirlpool.sciencelogic.net")
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println("CryptPass:", u.CryptP)
// fmt.Println("B64", u.B64)
