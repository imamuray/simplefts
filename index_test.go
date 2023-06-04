package main

import (
	"reflect"
	"testing"
)

func TestIndexAdd(t *testing.T) {
	idx := make(index)
	idx.add([]document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
	idx.add([]document{{ID: 2, Text: "donut is a donut"}})

	want := map[string][]int{
		"donut": {1, 2}, "donuts": {1}, "glass": {1}, "is": {2}, "on": {1}, "only": {1},
	}

	for key, value := range want {
		t.Run("indexへの追加テスト", func(t *testing.T) {
			if got, ok := idx[key]; !ok {
				t.Errorf("idx[%v] dosen't exist.", key)
			} else {
				if !reflect.DeepEqual(got, value) {
					t.Errorf("idx[%v] = %v, want %v", key, got, value)
				}
			}
		})
	}

}
