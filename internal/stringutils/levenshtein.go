package stringutils

import (
	"container/heap"
)

func FindTopMatches(words []string, word string, num int) []string {
	pq := make(PriorityQueue, 0, len(words))

	for i, w := range words {
		score := LevenshteinDistance(w, word)
		pq.Push(&Item{value: w, priority: score, index: i})
	}
	heap.Init(&pq)

	matches := make([]string, 0, num)
	for pq.Len() > 0 && len(matches) < num {
		item := heap.Pop(&pq).(*Item)
		matches = append(matches, item.value)
	}

	return matches
}

func LevenshteinDistance(str1 string, str2 string) int {
	// Create a 2D slice to store the distances.
	distances := make([][]int, len(str1)+1)
	for i := range distances {
		distances[i] = make([]int, len(str2)+1)
	}

	// Initialize the first row and column.
	for i := range distances {
		distances[i][0] = i
	}
	for i := range distances[0] {
		distances[0][i] = i
	}

	// Populate the matrix
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				distances[i][j] = distances[i-1][j-1]
			} else {
				distances[i][j] = 1 + min(distances[i-1][j], distances[i][j-1], distances[i-1][j-1])
			}
		}
	}
	return distances[len(str1)][len(str2)]
}
