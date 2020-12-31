package sl1

import (
	"fmt"

	"github.com/marco-ostaska/sl1cmd/pkg/sl1/httpcalls"
)

func ExampleBasicInfo_Load() {
	var bInfo BasicInfo
	httpcalls.Insecure = true // just for testing purposes

	if err := bInfo.Load("/api/account"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(bInfo[0].Description, bInfo[0].URI)

	// Output: em7admin /api/account/1

}

func ExampleBasicInfo_Sl1ID() {
	var bInfo BasicInfo
	httpcalls.Insecure = true // just for testing purposes

	if err := bInfo.Load("/api/account"); err != nil {
		fmt.Println(err)
	}

	sl1id, err := bInfo.Sl1ID("em7admin")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sl1id)

	// Output: 1
}
