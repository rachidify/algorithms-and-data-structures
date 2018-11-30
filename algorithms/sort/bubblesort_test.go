package sort_test

import (
	"reflect"
	"testing"

	"github.com/rasheedhamdawi/algorithms-and-data-structures/algorithms/sort"
)

var list, sortedList []int

func init() {
	list = []int{5, 8, 9, 6, 7, 3, 2, 1, 4}
	sortedList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func TestBubbleSort(t *testing.T) {
	sort.BubbleSort(list)

	if !reflect.DeepEqual(list, sortedList) {
		t.Errorf("expected %d to equal  %d", list, sortedList)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.BubbleSort(list)
	}
}
