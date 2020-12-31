package httpcalls

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsReachable(t *testing.T) {
	tt := []struct {
		name     string
		url      string
		expected string
		insecure bool
	}{
		{"Invalid Url", "http://www.g0gle.vv", "no such host", false},
		{"Valid URL", "http://www.google.com", "url ok", false},
		{"Bad Certificate", "https://192.168.4.30", "bad certificate", false},
		{"No route", "https://10.7.4.31", "no route", false},
		{"Insecure", "https://192.168.4.30", "insecure", true},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s", tc.name), func(t *testing.T) {
			Insecure = tc.insecure
			err := isReachable(tc.url)
			if err != nil {
				switch {
				case strings.Contains(err.Error(), "no such host"):
					return
				case strings.Contains(err.Error(), "cannot validate certificate"):
					return
				case strings.Contains(err.Error(), "no route to host"):
					return
				case strings.Contains(err.Error(), "Timeout exceeded while awaiting"):
					return
				case strings.Contains(err.Error(), "certificate relies"):
					return
				default:
					t.Errorf("%v: expected %v, got %v", tc.name, tc.expected, err)
				}
			}
		})

	}

}