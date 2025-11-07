package wc_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_basic() {
	// echo "Hello World\nSecond line" | wc
	yup.MustRun(
		Wc(strings.NewReader("Hello World\nSecond line")),
	)
	// Output:
	//       2      4      24
}

