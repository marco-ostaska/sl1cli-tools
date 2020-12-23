// Package httpcalls makes http request calls on sl1api
package httpcalls

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/marco-ostaska/sl1cli-tools/pkg/cryptcfg"
)

// isReachable checks if url is reachable
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

// APIData an abstraction to API
type APIData struct {
	API     string          // API section as : /api/account
	ARGS    string          // any parameter need to be sent to api
	Payload *strings.Reader //payload for posting
	Result  []byte          // result from call
}

// httpcalls make the http calls
func (a *APIData) httpcalls(method string) error {

	var uCFG cryptcfg.UserInfo
	if err := uCFG.ReadCryptFile(); err != nil {
		return err
	}

	if err := isReachable(uCFG.URL); err != nil {
		return fmt.Errorf("%s is unreachable", uCFG.URL)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url := uCFG.URL + a.API + a.ARGS + "?hide_filterinfo=1"

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

func (a *APIData) makeCall(method string, v interface{}) error {
	if err := a.httpcalls("GET"); err != nil {
		return err
	}

	return json.Unmarshal(a.Result, &v)

}

// NewRequest make new call to sl1 API
func (a *APIData) NewRequest(v interface{}, as ...string) error {
	a.API = as[0]
	if len(as) > 1 {
		a.ARGS = as[1]
	}

	return a.makeCall("GET", v)

}

// NewPost make new post to sl1 API
func (a *APIData) NewPost(v interface{}, as ...string) error {
	a.API = as[0]
	if len(as) == 2 {
		a.ARGS = as[1]
	}

	if len(as) > 2 {
		a.Payload = strings.NewReader(as[2])
	}

	return a.makeCall("POST", v)

}
