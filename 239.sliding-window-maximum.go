/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	heap := &Heap{}
	for i := 0; i < k; i++ {
		heap.Push([]int{nums[i], i})
	}
	ans := []int{}
	ans = append(ans, heap.Peek()[0])
	for i := k; i < len(nums); i++ {
		heap.Push([]int{nums[i], i})
		for !heap.Empty() && heap.Peek()[1] <= i-k {
			heap.Pop()
		}
		ans = append(ans, heap.Peek()[0])
	}
	return ans
}

type Heap [][]int

func (h *Heap) Push(x []int) {
	*h = append(*h, x)
	h.UpAdjust()
}

func (h *Heap) Pop() []int {
	temp := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.DownAdjust()
	return temp
}

func (h *Heap) Empty() bool {
	return len(*h) == 0
}

func (h *Heap) Peek() []int {
	return (*h)[0]
}

func (h *Heap) UpAdjust() {
	curr := len(*h) - 1
	parent := 0
	for curr != 0 {
		parent = (curr - 1) / 2
		if h.Compare((*h)[parent], (*h)[curr]) {
			break
		} else {
			(*h)[parent], (*h)[curr] = (*h)[curr], (*h)[parent]
		}
		curr = parent
	}
}

func (h *Heap) DownAdjust() {
	curr := 0
	leftChild := 2*curr + 1
	for leftChild < len(*h) {
		rightChild := leftChild + 1
		if rightChild < len(*h) && h.Compare((*h)[rightChild], (*h)[leftChild]) {
			leftChild = rightChild
		}
		if h.Compare((*h)[curr], (*h)[leftChild]) {
			break
		} else {
			(*h)[curr], (*h)[leftChild] = (*h)[leftChild], (*h)[curr]
		}
		curr = leftChild
		leftChild = 2*curr + 1
	}
}

func (h *Heap) Compare(a, b []int) bool {
	if a[0] == b[0] {
		return a[1] > b[1]
	} else {
		return a[0] > b[0]
	}
}

// @lc code=end

