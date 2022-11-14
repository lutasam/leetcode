/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */

// @lc code=start
package leetcode

func sortArray(nums []int) []int {
	return HeapSort(nums)
}

func HeapSort(nums []int) []int {
	for i := (len(nums) - 1) / 2; i >= 0; i-- {
		AdjustMaxHeap(nums, i, len(nums))
	}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		AdjustMaxHeap(nums, 0, i)
	}
	return nums
}

func AdjustMaxHeap(nums []int, root, len int) {
	val := nums[root]
	leftChild := 2*root + 1
	for leftChild < len {
		rightChild := leftChild + 1
		if rightChild < len && nums[leftChild] < nums[rightChild] {
			leftChild = rightChild
		}
		if val > nums[leftChild] {
			break
		}
		nums[root] = nums[leftChild]
		root = leftChild
		leftChild = 2*root + 1
	}
	nums[root] = val
}

// @lc code=end
