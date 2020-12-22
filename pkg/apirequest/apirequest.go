// December 2020
// v1.0.0

// Package apirequest makes http request calls on sl1api
package apirequest

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/marco-ostaska/sl1cli-tools/pkg/apicryptcfg"
)

// APIData an abstraction to API
type APIData struct {
	API    string // API section as : /api/account
	ARGS   string // any parameter need to be sent to api
	Result []byte // result from call
}

func isReachable(url string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	timeout := time.Duration(15 * time.Second)
	c := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	_, err := c.Get(url)

	return err
}

// apiRequest make the http calls
func (a *APIData) apiRequest() error {

	var uCFG apicryptcfg.UserInfo
	if err := uCFG.ReadCryptFile(); err != nil {
		return err
	}

	if err := isReachable(uCFG.URL); err != nil {
		return fmt.Errorf("%s is unreachable", uCFG.URL)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
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

// NewRequest make new request to sl1 API
func (a *APIData) NewRequest(v interface{}, as ...string) error {
	a.API = as[0]
	if len(as) > 1 {
		a.ARGS = as[1]
	}

	if err := a.apiRequest(); err != nil {
		return err
	}

	return json.Unmarshal(a.Result, &v)

}
