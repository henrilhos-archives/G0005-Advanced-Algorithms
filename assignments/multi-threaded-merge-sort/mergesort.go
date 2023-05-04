package mergesort

import (
	"sync"
)

func merge(nums []int, middle int) {
	helper := make([]int, len(nums))
	copy(helper, nums)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(nums) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			nums[current] = helper[helperLeft]
			helperLeft++
		} else {
			nums[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		nums[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}

func parallelMergeSort(nums []int) {
	len := len(nums)

	if len > 1 {
		middle := len / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			parallelMergeSort(nums[:middle])
		}()

		go func() {
			defer wg.Done()
			parallelMergeSort(nums[middle:])
		}()

		wg.Wait()
		merge(nums, middle)
	}
}
