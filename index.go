package main

import "strings"

type index map[string][]int

func (idx index) add(docs []document) {
	for _, doc := range docs {
		for _, word := range strings.Fields(doc.Text) {
			word = strings.ToLower(word)
			ids := idx[word]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[word] = append(ids, doc.ID)
		}
	}
}
