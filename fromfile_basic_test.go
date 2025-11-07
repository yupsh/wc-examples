package wc_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleWc_fromFile_basic() {
	// cat testdata/sample.txt | wc
	yup.MustRun(
		Wc(yup.File("testdata/sample.txt")),
	)
	// Output:
	//       3      8      42
}

