package utils

import "testing"

type equalizeStringsSizesTest struct {
	input  []string
	output []string
}

func createEqualizeStringsSizesTestTable() []equalizeStringsSizesTest {
	return []equalizeStringsSizesTest{
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

type validateMinLengthTest struct {
	arr               []string
	minLength         int
	shouldReturnError bool
}

func createValidateMinLengthTestTable() []validateMinLengthTest {
	return []validateMinLengthTest{
		{
			arr:               []string{"test1", "test2", "test3"},
			minLength:         3,
			shouldReturnError: false,
		},
		{
			arr:               []string{"test1", "test2", "test3"},
			minLength:         2,
			shouldReturnError: false,
		},
		{
			arr:               []string{"test1", "test2", "test3"},
			minLength:         4,
			shouldReturnError: true,
		},
		{
			arr:               []string{},
			minLength:         0,
			shouldReturnError: false,
		},
		{
			arr:               []string{},
			minLength:         1,
			shouldReturnError: true,
		},
		{
			arr:               nil,
			minLength:         0,
			shouldReturnError: true,
		},
	}
}

func TestValidateMinLength(t *testing.T) {
	for _, row := range createValidateMinLengthTestTable() {
		res := ValidateMinLength(row.minLength, row.arr)
		if row.shouldReturnError && res == nil {
			t.Errorf(
				"error was expected for [%v] with min length %d but got <nil>",
				row.arr,
				row.minLength,
			)
		} else if !row.shouldReturnError && res != nil {
			t.Errorf(
				"error was not expected for [%v] with min length %d but got: %v",
				row.arr,
				row.minLength,
				res,
			)
		}
	}
}
