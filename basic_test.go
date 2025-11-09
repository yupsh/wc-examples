package wc_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_basic() {
	// echo "Hello World\nSecond line" | wc
	gloo.MustRun(
		Wc(strings.NewReader("Hello World\nSecond line")),
	)
	// Output:
	//       2      4      24
}
