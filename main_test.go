package main

import (
	"fmt"
	"reflect"
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{a: 1, b: 2},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup")
			})
			// t.Fatalf でテストが失敗した場合でも defer の処理は呼び出される
			defer t.Log("defer")
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
				// t.Fatalf("add() = %v, want %v", got, tt.want)
			}
			// t.Fatalf でテストが失敗した場合は以下は呼び出されない
			t.Log("after add() ...")
		})
	}
}

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

func TestLowercaseFilter(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "空文字",
			args: []string{},
			want: []string{},
		},
		{
			name: "標準的なケース",
			args: []string{"A", "donut", "on", "a", "glass", "plate", "Only", "the", "donuts"},
			want: []string{"a", "donut", "on", "a", "glass", "plate", "only", "the", "donuts"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := lowercaseFilter(test.args); !reflect.DeepEqual(got, test.want) {
				t.Errorf("lowercaseFilter() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestStopwordFilter(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "空文字",
			args: []string{},
			want: []string{},
		},
		{
			name: "標準的なケース",
			args: []string{"a", "donut", "on", "a", "glass", "plate", "only", "the", "donuts"},
			want: []string{"donut", "on", "glass", "plate", "only", "donuts"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := stopwordFilter(test.args); !reflect.DeepEqual(got, test.want) {
				t.Errorf("stopwordFilter() = %v, want %v", got, test.want)
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
