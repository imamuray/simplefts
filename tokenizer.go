package main

import (
	"strings"
	"unicode"
)

// 単語境界でテキストを分割する
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// テキストをトークンに分割
func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	// ステミング入れるならここ
	return tokens
}
