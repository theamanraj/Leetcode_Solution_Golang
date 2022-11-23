package main

import "sort"

func frequencySort(s string) string {
	freq := make([]int, 255)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		if freq[b[i]] == freq[b[j]] {
			return b[i] > b[j]
		}
		return freq[b[i]] > freq[b[j]]
	})
	return string(b)
}