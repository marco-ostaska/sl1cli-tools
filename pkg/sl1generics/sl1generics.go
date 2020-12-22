// Package sl1generics have the generic routines
//to be used throughout the sl1tools
package sl1generics

import (
	"crypto/tls"
	"net/http"
	"strconv"
	"time"
)

// EpochToUnix converts epoch time to Unix time
func EpochToUnix(epoch string) (time.Time, error) {
	var t time.Time
	i, err := strconv.ParseInt(epoch, 10, 64)
	if err != nil {
		return t, err
	}
	t = time.Unix(i, 0)
	return t, nil
}

// IsReachable checks if url is reachable
func IsReachable(url string) error {
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
