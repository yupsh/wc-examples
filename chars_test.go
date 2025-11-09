package wc_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_chars() {
	// echo "Hello" | wc -m
	gloo.MustRun(
		Wc(Chars, strings.NewReader("Hello")),
	)
	// Output:
	//       5
}
