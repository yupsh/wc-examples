package wc_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_lines() {
	// echo "line1\nline2\nline3" | wc -l
	yup.MustRun(
		Wc(Lines, strings.NewReader("line1\nline2\nline3")),
	)
	// Output:
	//       3
}

