package main

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
