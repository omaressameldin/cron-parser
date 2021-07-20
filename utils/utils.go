package utils

import (
	"fmt"
	"os"
	"strings"
)

func Must(err error) {
	if err != nil {
		fmt.Println(err.Error())
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
