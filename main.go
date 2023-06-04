package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// 単純な文字列一致で検索
func search(docs []document, term string) []document {
	var r []document
	for _, doc := range docs {
		if strings.Contains(doc.Text, term) {
			r = append(r, doc)
		}
	}
	return r
}

// 正規表現で検索
// term が完全一致で検索できるようになるが時間がかかる
func searchRegexp(docs []document, term string) []document {
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)
	var r []document
	for _, doc := range docs {
		if re.MatchString(doc.Text) {
			r = append(r, doc)
		}
	}
	return r
}

// 単語境界でテキストを分割する
func tokinize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// テキストをトークンに分割
func analyze(text string) []string {
	tokens := tokinize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	// ステミング入れるならここ
	return tokens
}

// 転置インデックス
type index map[string][]int

// 転置インデックスに要素を追加
func (idx index) add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// 同じIDを2回数えないようにするため
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

// 2つのスライスの共通部分のスライスを返す
// それぞれのスライスの要素はソートされていることを前提とする
func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// 転置インデックスから要素を検索
func (idx index) search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %v filename term\n", os.Args[0])
		os.Exit(0)
	}

	term := os.Args[2]

	fmt.Println("load start.")
	startTime := time.Now()
	docs, err := loadDocuments(os.Args[1])
	if err != nil {
		fmt.Println("loadDocuments failed.")
		os.Exit(1)
	}
	fmt.Printf("load finised: %s\n", time.Since(startTime))
	fmt.Println("load finised.")

	fmt.Printf("term is %s.\n", term)
	fmt.Println("search start.")
	startTime = time.Now()
	// searchResult := search(docs, term)
	searchResult := searchRegexp(docs, term)
	fmt.Printf("search finised: %s\n", time.Since(startTime))
	fmt.Printf("hit: %d\n", len(searchResult))
	if len(searchResult) > 0 {
		fmt.Printf("first data: %v\n", searchResult[0])
	}

	idx := make(index)
	idx.add(docs)
	fmt.Println("index search start.")
	startTime = time.Now()
	indexSearchResult := idx.search(term)
	fmt.Printf("index search finised: %s\n", time.Since(startTime))
	fmt.Printf("hit: %d\n", len(indexSearchResult))
	if len(indexSearchResult) > 0 {
		// docs の ID とインデックスが揃っているからこの取り出し方にしている
		fmt.Printf("first data: %v\n", docs[indexSearchResult[0]])
	}
}
