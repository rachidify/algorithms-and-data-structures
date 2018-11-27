package search_test

import (
	"sort"
	"testing"

	"github.com/rasheedhamdawi/algorithms-and-data-structures/algorithms/search"
)

var list []int

func init() {
	list = []int{4, 59, 3, 93, 3, 85, 2, 8, 97}
	sort.Ints(list)
}

func TestBinarySearch(t *testing.T) {
	if result := search.BinarySearch(list, 8); result != 4 {
		t.Error("Expected 4, got", result)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search.BinarySearch(list, 8)
	}
}
