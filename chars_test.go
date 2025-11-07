package wc_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_chars() {
	// echo "Hello" | wc -m
	yup.MustRun(
		Wc(Chars, strings.NewReader("Hello")),
	)
	// Output:
	//       5
}

