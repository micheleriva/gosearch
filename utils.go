package main

func ContainsDoc(indexedTokenDocs []Tuple, target string) bool {
	return IndexedDocIndex(indexedTokenDocs, target) >= 0
}

func IndexedDocIndex(indexedTokenDocs []Tuple, target string) int {
	for index, doc := range indexedTokenDocs {
		if doc.id == target {
			return index
		}
	}

	return -1
}