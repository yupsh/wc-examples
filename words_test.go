package wc_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_words() {
	// echo "one two three four" | wc -w
	yup.MustRun(
		Wc(Words, strings.NewReader("one two three four")),
	)
	// Output:
	//       4
}

