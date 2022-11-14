package main

import (
	"container/heap"
	"math"
	"sort"
)

// data struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.two sum
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, exist := m[target-nums[i]]; exist {
			return []int{i, m[target-nums[i]]}
		}
		m[nums[i]] = i
	}
	return []int{-1, -1}
}

// 2.two numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		node := &ListNode{Val: sum % 10}
		carry = sum / 10
		curr.Next = node
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

// 3.Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	lenOfLongestSub := 1
	set := make(map[byte]struct{})
	left, right := 0, 0
	for left < len(s) {
		if left > 0 {
			delete(set, s[left-1])
		}
		for right < len(s) {
			if _, exist := set[s[right]]; exist {
				break
			}
			set[s[right]] = struct{}{}
			right++
		}
		lenOfLongestSub = max(lenOfLongestSub, right-left)
		left++
	}
	return lenOfLongestSub
}

// 4.Median of Two Sorted Arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	k1, k2 := (len(nums1)+len(nums2)+1)/2, (len(nums1)+len(nums2)+2)/2
	return float64(findKthSmall(nums1, 0, nums2, 0, k1)+findKthSmall(nums1, 0, nums2, 0, k2)) / 2.
}

func findKthSmall(nums1 []int, i int, nums2 []int, j int, k int) int {
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
		return findKthSmall(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return findKthSmall(nums1, i, nums2, j+k/2, k-k/2)
	}
}

// 5. Longest Palindromic Substring
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	lenOfLongestPalindrome, left, right := 0, 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := subLongestPalindromic(s, i, i)
		left2, right2 := subLongestPalindromic(s, i, i+1)
		len1, len2 := right1-left1+1, right2-left2+1
		if len1 > len2 && len1 > lenOfLongestPalindrome {
			lenOfLongestPalindrome = len1
			left = left1
			right = right1
		} else if len2 > len1 && len2 > lenOfLongestPalindrome {
			lenOfLongestPalindrome = len2
			left = left2
			right = right2
		}
	}
	return s[left : right+1]
}

func subLongestPalindromic(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j - 1
}

// 10. Regular Expression Matching
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if matches(s, i, p, j-1) {
					dp[i][j] = dp[i-1][j] || dp[i][j]
				}
			} else if matches(s, i, p, j) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}

func matches(s string, i int, p string, j int) bool {
	if i == 0 {
		return false
	}
	if p[j-1] == '.' {
		return true
	}
	return s[i-1] == p[j-1]
}

// 11. Container With Most Water
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	maxWaterArea, left, right := 0, 0, len(height)-1
	for left < right {
		tempWaterArea := (right - left) * min(height[left], height[right])
		maxWaterArea = max(tempWaterArea, maxWaterArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxWaterArea
}

// 15. 3Sum
func threeSum(nums []int) [][]int {
	threeSumSet := make([][]int, 0)
	if len(nums) < 3 {
		return threeSumSet
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				threeSumSet = append(threeSumSet, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return threeSumSet
}

// 17. Letter Combinations of a Phone Number
//var phoneNumberMap = map[byte]string{
//	'2': "abc",
//	'3': "def",
//	'4': "ghi",
//	'5': "jkl",
//	'6': "mno",
//	'7': "pqrs",
//	'8': "tuv",
//	'9': "wxyz",
//}
//
//func letterCombinations(digits string) []string {
//	ans := make([]string, 0)
//	if len(digits) == 0 {
//		return ans
//	}
//	ans = backtrack(digits, 0, "", ans)
//	return ans
//}
//
//func backtrack(digits string, curr int, s string, ans []string) []string {
//	if curr == len(digits) {
//		ans = append(ans, s)
//		return ans
//	}
//	chars := phoneNumberMap[digits[curr]]
//	for _, c := range chars {
//		s += string(c)
//		ans = backtrack(digits, curr+1, s, ans)
//		s = s[:len(s)-1]
//	}
//	return ans
//}

// 19. Remove Nth Node From End of List
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	pre := slow
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	pre.Next = pre.Next.Next
	return head
}

// 20. Valid Parentheses
//func isValid(s string) bool {
//	n := len(s)
//	if n%2 == 1 {
//		return false
//	}
//	pairs := map[byte]byte{
//		')': '(',
//		']': '[',
//		'}': '{',
//	}
//	stack := []byte{}
//	for i := 0; i < n; i++ {
//		if _, exist := pairs[s[i]]; exist {
//			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
//				return false
//			}
//			stack = stack[:len(stack)-1]
//		} else {
//			stack = append(stack, s[i])
//		}
//	}
//	return len(stack) == 0
//}

// 21. Merge Two Sorted Lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	n1, n2, curr := list1, list2, dummy
	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			curr.Next = n1
			curr = n1
			n1 = n1.Next
		} else {
			curr.Next = n2
			curr = n2
			n2 = n2.Next
		}
	}
	if n1 != nil {
		curr.Next = n1
	}
	if n2 != nil {
		curr.Next = n2
	}
	return dummy.Next
}

// 22. Generate Parentheses
// func generateParenthesis(n int) []string {
// 	if n == 0 {
// 		return []string{}
// 	}
// 	return backtrack(n, n, "", []string{})
// }

// func backtrack(left, right int, s string, ans []string) []string {
// 	if left == 0 && right == 0 {
// 		return append(ans, s)
// 	}
// 	if left == right {
// 		ans = backtrack(left-1, right, s+"(", ans)
// 	} else {
// 		if left > 0 {
// 			ans = backtrack(left-1, right, s+"(", ans)
// 		}
// 		ans = backtrack(left, right-1, s+")", ans)
// 	}
// 	return ans
// }

// 23. Merge k Sorted Lists
//func mergeKLists(lists []*ListNode) *ListNode {
//	if len(lists) == 0 {
//		return nil
//	}
//	return mergeLists(lists, 0, len(lists)-1)
//}
//
//func mergeLists(lists []*ListNode, left, right int) *ListNode {
//	if left >= right {
//		return lists[left]
//	}
//	mid := left + (right-left)/2
//	lefts, rights := mergeLists(lists, left, mid), mergeLists(lists, mid+1, right)
//	return merge(lefts, rights)
//}
//
//func merge(left, right *ListNode) *ListNode {
//	dummy := &ListNode{}
//	curr := dummy
//	for left != nil && right != nil {
//		if left.Val < right.Val {
//			curr.Next = left
//			left = left.Next
//		} else {
//			curr.Next = right
//			right = right.Next
//		}
//		curr = curr.Next
//	}
//	if left != nil {
//		curr.Next = left
//	}
//	if right != nil {
//		curr.Next = right
//	}
//	return dummy.Next
//}

type IntHeap []int

func (pq IntHeap) Len() int {
	return len(pq)
}

func (pq IntHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq IntHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *IntHeap) Push(obj interface{}) {
	*pq = append(*pq, obj.(int))
}

func (pq *IntHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	*pq = old[:n-1]
	return ret
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	minHeap := &IntHeap{}
	heap.Init(minHeap)
	for _, list := range lists {
		for list != nil {
			heap.Push(minHeap, list.Val)
			list = list.Next
		}
	}
	dummy := &ListNode{}
	curr := dummy
	for minHeap.Len() != 0 {
		val := heap.Pop(minHeap)
		node := &ListNode{Val: val.(int)}
		curr.Next = node
		curr = curr.Next
	}
	return dummy.Next
}

// 31. Next Permutation
func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}

	reverse := func(i int) {
		j := len(nums) - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(i + 1)
}

// 32. Longest Valid Parentheses
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] += 2
				if i-2 >= 0 {
					dp[i] += dp[i-2]
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] += 2 + dp[i-1]
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
		}
		ans = max(dp[i], ans)
	}

	return ans
}

// 33. Search in Rotated Sorted Array
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 34. Find First and Last Position of Element in Sorted Array
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	ansL, ansR := -1, -1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[right] != target {
		return []int{-1, -1}
	}
	ansL = right

	left, right = 0, len(nums)-1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	ansR = left
	return []int{ansL, ansR}
}

// 39. Combination Sum
//func combinationSum(candidates []int, target int) [][]int {
//	if len(candidates) == 0 {
//		return [][]int{}
//	}
//	return backtrack(candidates, 0, target, 0, []int{}, [][]int{})
//}
//
//func backtrack(candidates []int, curr, target, sum int, list []int, ans [][]int) [][]int {
//	if curr >= len(candidates) {
//		if sum == target {
//			ans = append(ans, append([]int{}, list...))
//		}
//		return ans
//	}
//
//	ans = backtrack(candidates, curr+1, target, sum, list, ans)
//
//	if sum+candidates[curr] <= target {
//		list = append(list, candidates[curr])
//		ans = backtrack(candidates, curr, target, sum+candidates[curr], list, ans)
//		list = list[:len(list)-1]
//	}
//	return ans
//}

// 42. Trapping Rain Water
func trap(height []int) int {
	lefts, rights := make([]int, len(height)), make([]int, len(height))
	lefts[0] = height[0]
	for i := 1; i < len(lefts); i++ {
		lefts[i] = max(lefts[i-1], height[i])
	}
	rights[len(height)-1] = height[len(height)-1]
	for i := len(rights) - 2; i >= 0; i-- {
		rights[i] = max(rights[i+1], height[i])
	}
	maxWaterVolumn := 0
	for i := 0; i < len(height); i++ {
		maxWaterVolumn += min(lefts[i], rights[i]) - height[i]
	}
	return maxWaterVolumn
}

// 46. Permutations
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	return backtrack(nums, 0, []int{}, [][]int{}, make([]bool, len(nums)))
}

func backtrack(nums []int, curr int, list []int, ans [][]int, vis []bool) [][]int {
	if curr >= len(nums) {
		return append(ans, append([]int{}, list...))
	}
	for i := 0; i < len(nums); i++ {
		if !vis[i] {
			vis[i] = true
			list = append(list, nums[i])
			ans = backtrack(nums, curr+1, list, ans, vis)
			list = list[:len(list)-1]
			vis[i] = false
		}
	}
	return ans
}

// 48. Rotate Image
func rotate(matrix [][]int) {
	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row/2; i++ {
		for j := 0; j < col; j++ {
			matrix[i][j], matrix[row-i-1][j] = matrix[row-i-1][j], matrix[i][j]
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 49. Group Anagrams
func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	m := make(map[string][]string, 0)
	for _, s := range strs {
		temp := []byte(s)
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		tempStr := string(temp)
		if _, exist := m[tempStr]; !exist {
			m[tempStr] = []string{}
		}
		m[tempStr] = append(m[tempStr], s)
	}
	for _, list := range m {
		ans = append(ans, list)
	}
	return ans
}

// 53. Maximum Subarray
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

// 55. Jump Game
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxJumpDistance := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxJumpDistance < i {
			return false
		}
		maxJumpDistance = max(maxJumpDistance, i+nums[i])
	}
	return true
}

// 56. Merge Intervals
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})
	ans := [][]int{intervals[0]}
	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if last[1] >= intervals[i][0] {
			last[1] = max(last[1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
			last = intervals[i]
		}
	}
	return ans
}

// 62. Unique Paths
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 64. Minimum Path Sum
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// 70. Climbing Stairs
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

// 72. Edit Distance
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			w1Del := dp[i-1][j] + 1
			w2Del := dp[i][j-1] + 1
			wChange := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				wChange++
			}
			dp[i][j] = min(w1Del, min(w2Del, wChange))
		}
	}
	return dp[m][n]
}

// 75. Sort Colors
func sortColors(nums []int) {
	left, right := -1, len(nums)-1
	for i := 0; i <= right; i++ {
		if nums[i] == 0 {
			left++
			nums[left], nums[i] = nums[i], nums[left]
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			i--
			right--
		}
	}
}

// 76. Minimum Window Substring
func minWindow(s string, t string) string {
	sm, tm := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		tm[t[i]]++
	}
	left, right := 0, 0
	tempMin := 0x3f3f3f3f
	l, r := -1, 0
	for right < len(s) {
		if _, exist := tm[s[right]]; exist {
			sm[s[right]]++
		}
		for left <= right && isValid(sm, tm) {
			if right-left+1 < tempMin {
				tempMin = right - left + 1
				l = left
				r = right
			}
			if _, exist := tm[s[left]]; exist {
				sm[s[left]]--
			}
			left++
		}
		right++
	}
	if l == -1 {
		return ""
	} else {
		return s[l : r+1]
	}
}

func isValid(sm, tm map[byte]int) bool {
	for key, val := range tm {
		if sVal, exist := sm[key]; !exist || sVal < val {
			return false
		}
	}
	return true
}

// 84. Largest Rectangle in Histogram
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	oldVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return oldVal
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func largestRectangleArea(heights []int) int {
	stack := Stack{}
	nums := append([]int{0}, heights...)
	nums = append(nums, 0)
	ans := 0
	for i := 0; i < len(nums); i++ {
		for !stack.IsEmpty() && nums[i] < nums[stack.Peek()] {
			ans = max(ans, nums[stack.Pop()]*(i-stack.Peek()-1))
		}
		stack.Push(i)
	}
	return ans
}



func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
