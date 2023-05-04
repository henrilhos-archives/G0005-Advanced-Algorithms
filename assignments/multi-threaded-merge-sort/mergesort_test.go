package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const size = 1000000

func TestMergeSort(t *testing.T) {
	nums := []int{5, 8, 9, 5, 0, 10, 1, 6}
	parallelMergeSort(nums)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, nums)

	nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	parallelMergeSort(nums)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, nums)
}

func BenchmarkParallelMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := random(size)
		b.StartTimer()
		parallelMergeSort(s)
		b.StopTimer()
	}
}
