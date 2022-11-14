/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
// package leetcode

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return QuickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func QuickSelect(nums []int, left, right, k int) int {
	if left >= right {
		return nums[left]
	}
	pivot := Partition(nums, left, right)
	if pivot == k {
		return nums[pivot]
	} else if pivot < k {
		return QuickSelect(nums, pivot+1, right, k)
	} else {
		return QuickSelect(nums, left, pivot-1, k)
	}
}

func Partition(nums []int, left, right int) int {
	pivot := left
	for left < right {
		for left < right && nums[pivot] < nums[right] {
			right--
		}
		for left < right && nums[pivot] >= nums[left] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[pivot], nums[left] = nums[left], nums[pivot]
	return left
}

// @lc code=end
