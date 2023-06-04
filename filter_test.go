package main

import (
	"reflect"
	"testing"
)

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
