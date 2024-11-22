package main

import (
	"fmt"
	"strings"
)

type index map[string][]int

func (idx index) add(docs []document) {
	for _, doc := range docs {
		//change to a tokenizer function later
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

func (idx index) search(query string) []int {
	var result []int
	for _, word := range strings.Fields(query) {
		word = strings.ToLower(word)
		fmt.Println("word:", word)
		ids := idx[word]
		if ids == nil {
			return nil
		}
		if result == nil {
			result = ids
		} else {
			result = intersection(result, ids)
		}
	}
	return result
}
