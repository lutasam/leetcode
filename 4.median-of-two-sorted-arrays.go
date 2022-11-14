/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	k1, k2 := (len1+len2+1)/2, (len1+len2+2)/2
	return float64(findKthSmallest(nums1, 0, nums2, 0, k1)+findKthSmallest(nums1, 0, nums2, 0, k2)) / 2.
}

func findKthSmallest(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}
	mid1, mid2 := math.MaxInt, math.MaxInt
	if i+k/2-1 < len(nums1) {
		mid1 = nums1[i+k/2-1]
	}
	if j+k/2-1 < len(nums2) {
		mid2 = nums2[j+k/2-1]
	}
	if mid1 < mid2 {
		return findKthSmallest(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return findKthSmallest(nums1, i, nums2, j+k/2, k-k/2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

