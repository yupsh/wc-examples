package wc_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_fromFile_lines() {
	// cat testdata/sample.txt | wc -l
	gloo.MustRun(
		Wc(Lines, gloo.File("testdata/sample.txt")),
	)
	// Output:
	//       3
}
