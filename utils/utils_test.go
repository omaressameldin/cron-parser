package utils

import "testing"

type EqualizeStringsSizesTest struct {
	input  []string
	output []string
}

func createEqualizeStringsSizesTestTable() []EqualizeStringsSizesTest {
	return []EqualizeStringsSizesTest{
		{
			input: []string{"test", "t", "", "many tests"},
			output: []string{
				"test      ",
				"t         ",
				"          ",
				"many tests",
			},
		},
	}
}

func TestEqualizeStringsSizes(t *testing.T) {
	for _, row := range createEqualizeStringsSizesTestTable() {
		res := EqualizeStringsSizes(row.input)
		for i, word := range res {
			if word != row.output[i] {
				t.Errorf(
					"Incorrect output expected (%s), got (%s)",
					row.output[i],
					word,
				)
			}
		}
	}
}
