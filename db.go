package main

import (
	"github.com/cornelk/hashmap"
	"github.com/gofrs/uuid"
)

type IndexedDoc struct {
	Id string
	Content string
}

type IndexContent struct {
	Id string
	Occurrences int
}

var docs *hashmap.HashMap
var index *hashmap.HashMap

func init() {
	index = &hashmap.HashMap{}
	docs = &hashmap.HashMap{}
}

func IndexDocument(input string) IndexedDoc {
	id := uuid.Must(uuid.NewV4()).String()
	tokenized := Tokenize(input)
	countedTokens := CountTokens(tokenized)

	for k, v := range countedTokens {
		values, exists := index.Get(k)

		var newValues []IndexContent

		if exists {
			for _, val := range values.([]IndexContent) {
				newValues = append(newValues, val)
			}
		}

		newValues = append(newValues, IndexContent{id, v})
		index.Set(k, newValues)
	}

	docs.Set(id, input)

	return IndexedDoc{
		Id: id,
		Content: input,
	}
}

func Search(input string) []interface{} {
	tokens := Tokenize(input)
	indexedResults := make([]IndexContent, 0)

	for _, token := range tokens {
		indexedTokenDocs, _ := index.Get(token)

		for _, indexedTokenDoc := range indexedTokenDocs.([]IndexContent) {
			indexedDocIdx := IndexedDocIndex(indexedResults, indexedTokenDoc.Id)
			if indexedDocIdx >= 0 {
				oldContent := indexedResults[indexedDocIdx]
				indexedResults[indexedDocIdx] = IndexContent{ Id: oldContent.Id, Occurrences: oldContent.Occurrences + 1}
			} else {
				indexedResults = append(indexedResults, indexedTokenDoc)
			}
		}
	}

	var indexedDocs = make([]interface{}, 0)

	for _, result := range indexedResults {
		document, _ := docs.Get(result.Id)
		indexedDocs = append(indexedDocs, document)
	}

	return indexedDocs
}