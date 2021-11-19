package main

import (
	"fmt"
	"github.com/cornelk/hashmap"
	"github.com/gofrs/uuid"
)

type Tuple struct {
	id, occurrencies interface{}
}

var index *hashmap.HashMap
var docs *hashmap.HashMap

func init() {
	index = &hashmap.HashMap{}
	docs = &hashmap.HashMap{}
}

func IndexDocument(input string) {
	id := uuid.Must(uuid.NewV4()).String()
	tokenized := Tokenize(input)
	countedTokens := CountTokens(tokenized)

	for k, v := range countedTokens {
		values, exists := index.Get(k)

		var newValues []Tuple

		if exists {
			for _, val := range values.([]Tuple) {
				newValues = append(newValues, val)
			}
		}

		newValues = append(newValues, Tuple{id, v})
		index.Set(k, newValues)
	}

	docs.Set(id, input)
}

func Search(input string) []interface{} {
	tokens := Tokenize(input)
	indexedTokens := make([]interface{}, 0)

	for _, token := range tokens {
		indexToken, _ := index.Get(token)
		indexedTokens = append(indexedTokens, indexToken)
	}

	return indexedTokens
}

func main() {
	IndexDocument("My name is Michele. Michele is my name.")
	IndexDocument("Hello everyone! I'm Michele.")
	IndexDocument("So here I won't way my name. Sorry.")
	IndexDocument("I'm indexing this. Hello world.")

	searchResults := Search("hello world")

	for _, wordResult := range searchResults {
		for _, doc := range wordResult.([]Tuple) {
			document, _ := docs.Get(doc.id)

			fmt.Println(document)
		}
	}
}
