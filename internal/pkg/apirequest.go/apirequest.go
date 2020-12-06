package apirequest

import (
	"io/ioutil"
	"net/http"

	"github.com/marco-ostaska/sl1cli-tools/internal/pkg/apicryptcfg"
)

// APIData an abstraction to API
type APIData struct {
	API    string // API section as : /api/account
	ARGS   string // any parameter need to be sent to api
	Result []byte // result from call
}

// APIRequest wrapper to make http calls
func (a *APIData) APIRequest() error {

	var uCFG apicryptcfg.UserInfo
	if err := uCFG.ReadCryptFile(); err != nil {
		return err
	}

	url := uCFG.URL + a.API + a.ARGS + "?hide_filterinfo=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+uCFG.B64)

	res, err := client.Do(req)
	defer func() {
		cerr := res.Body.Close()
		if err == nil {
			err = cerr
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	a.Result = body
	return nil
}
