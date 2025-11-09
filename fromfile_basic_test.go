package wc_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleWc_fromFile_basic() {
	// cat testdata/sample.txt | wc
	gloo.MustRun(
		Wc(gloo.File("testdata/sample.txt")),
	)
	// Output:
	//       3      8      42
}
