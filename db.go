package main

import (
	"fmt"
	"github.com/cornelk/hashmap"
	"github.com/gofrs/uuid"
)

type IndexedDoc struct {
	Id string
	Content string
}

type IndexContent struct {
	Token string
	Occurrences int
}

var docs *hashmap.HashMap
var index *hashmap.HashMap

func init() {
	index = &hashmap.HashMap{}
	docs = &hashmap.HashMap{}
}

func IndexDocument(input string, id string) IndexedDoc {
	if id == "" {
		id = uuid.Must(uuid.NewV4()).String()
	}
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

func UpdateDocument(id string, newContent string) error {
	oldContent, ok := docs.Get(id)

	if !ok {
		return fmt.Errorf("document with id %s does not exist", id)
	}

	oldTokens := Tokenize(oldContent.(string))

	for _, token := range oldTokens {
		indexedToken, ok := index.Get(token)
		if !ok {
			return fmt.Errorf("unable to update document with id %s. token does %s not exist", id, indexedToken)
		}

		index.Set(token, RemoveToken(indexedToken.([]IndexContent), token))
	}

	docs.Del(id)
	IndexDocument(newContent, id)

	return nil
}

func Search(input string) []string {
	tokens := Tokenize(input)
	indexedResults := make([]IndexContent, 0)

	for _, token := range tokens {
		indexedTokenDocs, _ := index.Get(token)

		for _, indexedTokenDoc := range indexedTokenDocs.([]IndexContent) {
			indexedDocIdx := IndexedDocIndex(indexedResults, indexedTokenDoc.Token)
			if indexedDocIdx >= 0 {
				oldContent := indexedResults[indexedDocIdx]
				indexedResults[indexedDocIdx] = IndexContent{ Token: oldContent.Token, Occurrences: oldContent.Occurrences + 1}
			} else {
				indexedResults = append(indexedResults, indexedTokenDoc)
			}
		}
	}

	var indexedDocs = make([]string, 0)

	for _, result := range indexedResults {
		document, _ := docs.Get(result.Token)

		indexedDocs = append(indexedDocs, document.(string))
	}

	return indexedDocs
}