package main

import (
	"fmt"
	"github.com/cornelk/hashmap"
	"github.com/gofrs/uuid"
	"github.com/micheleriva/lepus/tokenizer"
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
	tokenized := tokenizer.Tokenize(input)
	countedTokens := tokenizer.CountTokens(tokenized)

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
	tokens := tokenizer.Tokenize(input)
	indexedResults := make([]Tuple, 0)

	for _, token := range tokens {
		indexedTokenDocs, _ := index.Get(token)

		for _, indexedTokenDoc := range indexedTokenDocs.([]Tuple) {
			indexedDocIdx := IndexedDocIndex(indexedResults, indexedTokenDoc.id.(string))
			if indexedDocIdx >= 0 {
				oldContent := indexedResults[indexedDocIdx]
				indexedResults[indexedDocIdx] = Tuple{ id: oldContent.id, occurrencies: oldContent.occurrencies.(int) + 1}
			} else {
				indexedResults = append(indexedResults, indexedTokenDoc)
			}
		}
	}

	var indexedDocs = make([]interface{}, 0)

	for _, result := range indexedResults {
		document, _ := docs.Get(result.id)
		indexedDocs = append(indexedDocs, document)
	}

	return indexedDocs
}

func main() {
	IndexDocument("My name is Michele. Michele is my name.")
	IndexDocument("Hello everyone! I'm Michele.")
	IndexDocument("So here I won't way my name. Sorry.")
	IndexDocument("I'm indexing this. Hello world.")

	searchResult := Search("hello world")

	for _, result := range searchResult {
		fmt.Println(result)
	}
}
