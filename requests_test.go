package requests_test

import (
	"fmt"
	"github.com/r0fls/requests"
)

func ExampleGet() {
	v := requests.Get("https://www.reddit.com/r/programming")
	fmt.Println(v.Status)
	// Output: 200 OK
}
