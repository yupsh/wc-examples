package wc_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_words() {
	// echo "one two three four" | wc -w
	gloo.MustRun(
		Wc(Words, strings.NewReader("one two three four")),
	)
	// Output:
	//       4
}
