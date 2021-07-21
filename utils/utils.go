package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Must(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func EqualizeStringsSizes(words []string) []string {
	longestString := getLongest(words)

	return setSizes(words, len(longestString))
}

func getLongest(words []string) string {
	longestSoFar := ""
	for _, word := range words {
		if len(word) > len(longestSoFar) {
			longestSoFar = word
		}
	}

	return longestSoFar
}

func setSizes(words []string, newSize int) []string {
	newList := make([]string, 0, len(words))

	for _, word := range words {
		var sb strings.Builder
		spaces := newSize - len(word)
		sb.WriteString(word)
		for i := spaces; i > 0; i-- {
			sb.WriteRune(' ')
		}
		newList = append(newList, sb.String())
	}

	return newList
}

func ValidateMinLength(minSize int, arr []string) error {
	if arr == nil || len(arr) < minSize {
		return fmt.Errorf(
			"array is smaller than expected needed: %d, got %d",
			minSize,
			len(arr),
		)
	}

	return nil
}

func CreateArrFrom(start, end int) ([]int, error) {
	if end < start {
		return nil, fmt.Errorf("end can't be smaller than start, start: %d, end: %d", start, end)
	}
	arr := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		arr = append(arr, i)
	}

	return arr, nil
}

func ConvertIntArrToString(arr []int, separator string) string {
	strArr := make([]string, len(arr))
	for i, item := range arr {
		strArr[i] = strconv.Itoa(item)
	}

	return strings.Join(strArr, separator)
}

func GetArrValues(arr []int, step int) []int {
	transformedArr := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i += step {
		transformedArr = append(transformedArr, arr[i])
	}

	return transformedArr
}
