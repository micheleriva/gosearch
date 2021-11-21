package tokenizer

import (
	"log"
	"regexp"
	"strings"
)

func CountTokens(tokens []string) map[string]int {
	dict := make(map[string]int)

	for _, token := range tokens {
		dict[token] = dict[token] + 1
	}

	return dict
}

func Tokenize(input string) []string {
	lowerInput := strings.ToLower(input)
	return strings.Fields(filterSpecialCharacters(lowerInput))
}

func filterSpecialCharacters(input string) string {
	re, err := regexp.Compile(`[^\w|\d|\s]`)
	if err != nil {
		log.Fatal(err)
	}

	return re.ReplaceAllString(input, "")
}