package wc_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_lines() {
	// echo "line1\nline2\nline3" | wc -l
	gloo.MustRun(
		Wc(Lines, strings.NewReader("line1\nline2\nline3")),
	)
	// Output:
	//       3
}
