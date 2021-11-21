package main

func ContainsDoc(indexedTokenDocs []IndexContent, target string) bool {
	return IndexedDocIndex(indexedTokenDocs, target) >= 0
}

func IndexedDocIndex(indexedTokenDocs []IndexContent, target string) int {
	for index, doc := range indexedTokenDocs {
		if doc.Id == target {
			return index
		}
	}

	return -1
}