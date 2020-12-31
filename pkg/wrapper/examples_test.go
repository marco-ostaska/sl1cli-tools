package wrapper_test

import (
	"fmt"

	"github.com/marco-ostaska/sl1cmd/pkg/wrapper"
)

func ExampleEpochToUnix() {

	unixTime, err := wrapper.EpochToUnix("1609427816")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(unixTime)

	// Output:
	// 2020-12-31 12:16:56 -0300 -03
}
