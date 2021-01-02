package sl1_test

import (
	"fmt"
	"log"

	"github.com/marco-ostaska/sl1cmd/pkg/sl1"
	"github.com/marco-ostaska/sl1cmd/pkg/sl1/httpcalls"
)

func ExampleBasicInfo_Load() {

	var bInfo sl1.BasicInfo
	httpcalls.Insecure = true // just for testing purposes

	if err := bInfo.Load("/api/account"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(bInfo[0])

	// Output: {/api/account/1 em7admin 1}

}

func ExampleBasicInfo_LookupIdxByDesc() {
	var bInfo sl1.BasicInfo
	httpcalls.Insecure = true // to accept invalid certificate

	if err := bInfo.Load("/api/account"); err != nil {
		log.Fatalln(err)
	}

	i, err := bInfo.LookupIdxByDesc("em7admin")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(i)

	// Output: 0
}

func ExampleBasicInfo_LookupIDbyDesc() {
	var bInfo sl1.BasicInfo
	httpcalls.Insecure = true // to accept invalid certificate

	if err := bInfo.Load("/api/account"); err != nil {
		fmt.Println(err)
	}

	sl1id, err := bInfo.LookupIDbyDesc("em7admin")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(sl1id)

	// Output: 1
}

func ExampleBasicInfo_LookupDescByURI() {
	var bInfo sl1.BasicInfo
	httpcalls.Insecure = true // to accept invalid certificate

	if err := bInfo.Load("/api/account"); err != nil {
		log.Fatalln(err)
	}

	desc, err := bInfo.LookupDescByURI("/api/account/1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(desc)
	// Output:
	// em7admin
}
