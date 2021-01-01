package sl1_test

import (
	"strings"
	"testing"

	"github.com/marco-ostaska/sl1cmd/pkg/sl1"
	"github.com/marco-ostaska/sl1cmd/pkg/sl1/httpcalls"
)

func TestBasicInfo_Load(t *testing.T) {
	tt := []struct {
		name string
		api  string
	}{
		{"/api/account", "/api/account"},
		{"invalid api", "/api/accountsss"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var bInfo sl1.BasicInfo
			httpcalls.Insecure = true // just for testing purposes

			if err := bInfo.Load(tc.api); err != nil {
				if strings.Contains(err.Error(), "json: cannot unmarshal object into Go value of type sl1.BasicInfo") {
					return
				}
				t.Errorf("%s", err)
			}
		})
	}
}

func TestBasicInfo_SearchByURI(t *testing.T) {
	tt := []struct {
		name string
		api  string
	}{
		{"/api/account/1", "/api/account/1"},
		{"invalid uri", "/api/accountsss"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var bInfo sl1.BasicInfo
			httpcalls.Insecure = true // just for testing purposes
			bInfo.Load("/api/account/")

			if _, err := bInfo.SearchByURI(tc.api); err != nil {
				if strings.Contains(err.Error(), "no such uri") {
					return
				}
				t.Errorf("%s", err)
			}
		})
	}
}

func TestBasicInfo_IndexPosition(t *testing.T) {
	tt := []struct {
		name string
		args string
	}{
		{"/api/account/1", "em7admin"},
		{"Invalid Position", "blank"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var bInfo sl1.BasicInfo
			httpcalls.Insecure = true // just for testing purposes
			bInfo.Load("/api/account/")
			if _, err := bInfo.IndexPosition(tc.args); err != nil {
				if strings.Contains(err.Error(), "no such") {
					return
				}
				t.Errorf("%s", err)
			}
		})
	}
}

func TestBasicInfo_Sl1ID(t *testing.T) {
	tt := []struct {
		name string
		args string
	}{
		{"/api/account/1", "em7admin"},
		{"Invalid Position", "blank"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var bInfo sl1.BasicInfo
			httpcalls.Insecure = true // just for testing purposes
			bInfo.Load("/api/account/")
			if _, err := bInfo.Sl1ID(tc.args); err != nil {
				if strings.Contains(err.Error(), "no such") {
					return
				}
				t.Errorf("%s", err)
			}
		})
	}
}
