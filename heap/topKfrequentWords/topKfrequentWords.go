// #692
// https://leetcode.com/problems/top-k-frequent-words/description/
package leetcode

import (
	"container/heap"
	"maps"
	"slices"
)

type Pair struct {
	str  string
	freq int
}

type Heap []Pair

func (h Heap) Len() int      { return len(h) }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)   { *h = append(*h, x.(Pair)) }

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h Heap) Less(i, j int) bool {
	// Prioritize lexicographical order for strings with equal frequencies.
	if h[i].freq == h[j].freq {
		return h[i].str < h[j].str
	}
	// Otherwise, prioritize strings with higher frequencies.
	return h[i].freq > h[j].freq
}

func topKFrequent(words []string, k int) []string {
	// Create a hash map that counts the frequency of each string.
	freqs := make(map[string]Pair)
	for _, word := range words {
		if pair, ok := freqs[word]; ok {
			pair.freq++
			freqs[word] = pair
		} else {
			freqs[word] = Pair{str: word, freq: 1}
		}
	}
	// Collect all `Pairs' in slice.
	pairs := Heap(slices.Collect(maps.Values(freqs)))
	// Heapify on all string - frequency pairs.
	heap.Init(&pairs)
	// Pop the most frequent string off the heap 'k' times and return
	// these 'k' most frequent strings.
	result := make([]string, 0, k)
	for range k {
		result = append(result, heap.Pop(&pairs).(Pair).str)
		// result = append(result, pairs.Pop().(Pair).str)
	}
	return result
}
