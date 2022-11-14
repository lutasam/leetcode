/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start

func canFinish(numCourses int, prerequisites [][]int) bool {
	courseGraph := map[int]*Queue{}
	indegrees := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		courseGraph[i] = &Queue{}
	}
	for i := 0; i < len(prerequisites); i++ {
		courseGraph[prerequisites[i][1]].Push(prerequisites[i][0])
		indegrees[prerequisites[i][0]]++
	}
	courseList := &Queue{}
	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			courseList.Push(i)
		}
	}
	count := 0
	for !courseList.Empty() {
		course := courseList.Pop()
		count++
		size := courseGraph[course].Size()
		for i := 0; i < size; i++ {
			c := courseGraph[course].Pop()
			indegrees[c]--
			if indegrees[c] == 0 {
				courseList.Push(c)
			}
		}
	}
	return count == numCourses
}

type Queue []int

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}

func (q *Queue) Pop() int {
	temp := (*q)[0]
	*q = (*q)[1:]
	return temp
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func (q *Queue) Size() int {
	return len(*q)
}

// @lc code=end
