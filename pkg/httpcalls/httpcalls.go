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

// Package httpcalls makes http request calls on sl1api
package httpcalls

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/marco-ostaska/sl1cmd/pkg/cryptcfg"
)

// Insecure variable used to validate certificate
var Insecure bool

// isReachable checks if url is reachable
func isReachable(url string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: Insecure},
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
	API     string    // API section as : /api/account
	ARGS    string    // any parameter need to be sent to api
	Payload io.Reader //payload for posting
	Result  []byte    // result from call
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
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: Insecure}
	url := uCFG.URL + a.API + a.ARGS + "?hide_filterinfo=1"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, a.Payload)

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
	err := a.httpcalls(method)
	if err != nil {
		return err
	}

	if method == "GET" {
		return json.Unmarshal(a.Result, &v)
	}

	return err

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

// DeleteRequest make delete call to sl1 API
func (a *APIData) DeleteRequest() error {
	return a.httpcalls("DELETE")

}
