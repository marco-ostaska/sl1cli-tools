package httpcalls_test

import (
	"fmt"
	"log"

	"github.com/marco-ostaska/sl1cmd/pkg/sl1/httpcalls"
)

func ExampleAPIData_NewRequest() {
	httpcalls.Insecure = true // used for invalid certificates

	var desc []struct {
		URI         string `json:"URI"`
		Description string `json:"description"`
	}

	var a httpcalls.APIData
	a.API = "/api/account"
	if err := a.NewRequest(&desc); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(desc[0])

	// Output:
	// {/api/account/1 em7admin}

}
