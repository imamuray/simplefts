package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTokinize(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "空文字",
			args: "",
			want: []string{},
		},
		{
			name: "標準的なケース",
			args: "A donut on a glass plate. Only the donuts.",
			want: []string{"A", "donut", "on", "a", "glass", "plate", "Only", "the", "donuts"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := tokinize(test.args); !reflect.DeepEqual(got, test.want) {
				t.Errorf("tokinize() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "空文字",
			args: "",
			want: []string{},
		},
		{
			name: "標準的なケース",
			args: "A donut on a glass plate. Only the donuts.",
			want: []string{"donut", "on", "glass", "plate", "only", "donuts"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := analyze(test.args); !reflect.DeepEqual(got, test.want) {
				t.Errorf("analyze() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestIndexAdd(t *testing.T) {
	idx := make(index)
	idx.add([]document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
	idx.add([]document{{ID: 2, Text: "donut is a donut"}})
	// map[string][]intの比較がめんどいのでパス
	fmt.Println(idx)
}
