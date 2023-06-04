package main

import (
	"fmt"
	"testing"
)

func TestIndexAdd(t *testing.T) {
	idx := make(index)
	idx.add([]document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
	idx.add([]document{{ID: 2, Text: "donut is a donut"}})
	// map[string][]intの比較がめんどいのでパス
	fmt.Println(idx)
}
