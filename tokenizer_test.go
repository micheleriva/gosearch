package main

import (
	"reflect"
	"testing"
)

type TokenizeTest struct {
	test []string
	expect []string
}

type CountTokensTest struct {
	test map[string]int
	expect map[string]int
}

func TestTokenize(t *testing.T) {

	cases := []TokenizeTest{
		{
			test:   Tokenize("hello, world!"),
			expect: []string{"hello", "world"},
		},
		{
			test:   Tokenize("Lorem ipsum. Dolor? Sit amet!"),
			expect: []string{"lorem", "ipsum", "dolor", "sit", "amet"},
		},
		{
			test:   Tokenize(""),
			expect: []string{},
		},
	}

	for i, testCase := range cases {
		if !reflect.DeepEqual(testCase.test, testCase.expect) {
			t.Fatalf(`%c case doesn't match'`, i)
		}
	}
}

func TestCountTokens(t *testing.T) {
	cases := []CountTokensTest{
		{
			test:   CountTokens([]string{"hello", "world"}),
			expect: map[string]int{"world": 1, "hello": 1},
		},
		{
			test:   CountTokens([]string{"this", "is", "duplicated", "duplicated", "is"}),
			expect: map[string]int{"duplicated": 2, "is": 2, "this": 1},
		},
	}

	for i, testCase := range cases {
		if !reflect.DeepEqual(testCase.test, testCase.expect) {
			t.Fatalf(`%c case doesn't match'`, i)
		}
	}
}