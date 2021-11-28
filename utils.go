package main

func ContainsDoc(indexedTokenDocs []IndexContent, target string) bool {
	return IndexedDocIndex(indexedTokenDocs, target) >= 0
}

func IndexedDocIndex(indexedTokenDocs []IndexContent, target string) int {
	for index, doc := range indexedTokenDocs {
		if doc.Token == target {
			return index
		}
	}

	return -1
}

func RemoveToken(index []IndexContent, token string) []IndexContent {
	newIndex := make([]IndexContent, 0)

	for _, indexedToken := range index {
		if indexedToken.Token != token {
			newIndex = append(newIndex, indexedToken)
		}
	}

	return newIndex
}

func IsEqualSliceOfStrings(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}