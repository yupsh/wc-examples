package wc_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/wc"
)

func ExampleWc_fromFile_lines() {
	// cat testdata/sample.txt | wc -l
	yup.MustRun(
		Wc(Lines, yup.File("testdata/sample.txt")),
	)
	// Output:
	//       3
}

