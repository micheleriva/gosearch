package main

import (
	"reflect"
	"testing"
)

func TestDBMethods(t *testing.T) {
	doc1 := "hello, world"
	doc2 := "hello, world! This is Golang."
	doc3 := "hello, Golang!"

	idx1 := IndexDocument(doc1, "")
	IndexDocument(doc2, "")
	IndexDocument(doc3, "")

	result1 := Search("hello world")
	result2 := Search("golang")
	result3 := Search("this is")

	if !reflect.DeepEqual(result1, []string{doc1, doc2, doc3}) {
		t.Fatalf("retrieved documents in wrong order: %s", result1)
	}

	if !reflect.DeepEqual(result2, []string{doc2, doc3}) {
		t.Fatalf("retrieved documents in wrong order: %s", result2)
	}

	if !reflect.DeepEqual(result3, []string{doc2}) {
		t.Fatalf("retrieved documents in wrong order: %s", result3)
	}

	err := UpdateDocument(idx1.Id, "hello, haskell")
	if err != nil {
		t.Fatalf("unable to update doc %s", idx1.Id)
	}

	result4 := Search("haskell")
	Search("world")
	Search("hello world")

	if !reflect.DeepEqual(result4, []string{"hello, haskell"}) {
		t.Fatalf("retrieved documents in wrong order: %s", result4)
	}
}
