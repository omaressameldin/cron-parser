package utils

import (
	"reflect"
	"testing"
)

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

type createArrFromTest struct {
	start             int
	end               int
	step              int
	output            []int
	shouldReturnError bool
}

func createArrFromTestTable() []createArrFromTest {
	return []createArrFromTest{
		{
			output:            []int{2, 3, 4, 5},
			start:             2,
			end:               5,
			step:              1,
			shouldReturnError: false,
		},
		{
			output:            []int{-2, 0, 2},
			start:             -2,
			end:               2,
			step:              2,
			shouldReturnError: false,
		},
		{
			output:            []int{2},
			start:             2,
			end:               4,
			step:              5,
			shouldReturnError: false,
		},
		{
			output:            nil,
			start:             2,
			end:               2,
			step:              0,
			shouldReturnError: true,
		},
		{
			output:            nil,
			start:             2,
			end:               1,
			shouldReturnError: true,
		},
	}
}

func TestCreateArrFrom(t *testing.T) {
	for _, row := range createArrFromTestTable() {
		res, err := CreateArrFrom(row.start, row.end, row.step)
		if row.shouldReturnError && (err == nil || res != nil) {
			t.Errorf(
				"expected: (<nil>,error), got: (%v,%v)",
				res,
				err,
			)
		} else if !row.shouldReturnError && (err != nil || !reflect.DeepEqual(res, row.output)) {
			t.Errorf(
				"expected: (%v,<nil>), got: (%v,%v)",
				row.output,
				res,
				err,
			)
		}
	}
}

type convertIntArrToStringTest struct {
	arr       []int
	separator string
	expected  string
}

func createConvertIntArrToStringTestTable() []convertIntArrToStringTest {
	return []convertIntArrToStringTest{
		{
			arr:       []int{2, 3, 4, 5},
			separator: "-",
			expected:  "2-3-4-5",
		},
		{
			arr:       []int{2, 3, 4, 5},
			separator: ",",
			expected:  "2,3,4,5",
		},
		{
			arr:       []int{2},
			separator: "-",
			expected:  "2",
		},
		{
			arr:       []int{},
			separator: "",
			expected:  "",
		},
	}
}

func TestConvertIntArrToString(t *testing.T) {
	for _, row := range createConvertIntArrToStringTestTable() {
		res := ConvertIntArrToString(row.arr, row.separator)
		if res != row.expected {
			t.Errorf("Incorrect output expected (%s), got (%s)", row.expected, res)
		}
	}
}

type convertMapToSortedArrayTest struct {
	input    map[int]bool
	expected []int
}

func createConvertMapToSortedArrayTestTable() []convertMapToSortedArrayTest {
	return []convertMapToSortedArrayTest{
		{
			input: map[int]bool{
				10: true,
				2:  true,
				4:  true,
				7:  true,
				9:  true,
			},
			expected: []int{2, 4, 7, 9, 10},
		},
		{
			input: map[int]bool{
				10: true,
			},
			expected: []int{10},
		},
		{
			input: map[int]bool{
				10:  true,
				-10: true,
				0:   true,
				5:   true,
			},
			expected: []int{-10, 0, 5, 10},
		},
		{
			input:    map[int]bool{},
			expected: []int{},
		},
	}
}

func TestConvertMapToSortedArray(t *testing.T) {
	for _, row := range createConvertMapToSortedArrayTestTable() {
		res := ConvertMapToSortedArr(row.input)
		if !reflect.DeepEqual(row.expected, res) {
			t.Errorf("expected [%v], got [%v]", row.expected, res)
		}
	}
}
