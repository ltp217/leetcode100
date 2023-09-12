package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// 2023.4.26
// https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	// TODO
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// https://leetcode.cn/problems/moving-stones-until-consecutive/
func numMovesStones(a int, b int, c int) []int {
	x := min(min(a, b), c)
	z := max(max(a, b), c)
	y := a + b + c - x - z
	res := []int{2, z - x - 2}
	if z-y == 1 && y-x == 1 {
		res[0] = 0
	} else if z-y <= 2 || y-x <= 2 {
		res[0] = 1
	}
	return res
}

// https://leetcode.cn/problems/time-needed-to-inform-all-employees/
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := make(map[int][]int)
	for i, m := range manager {
		g[m] = append(g[m], i)
	}
	var dfs func(int) int
	dfs = func(cur int) (res int) {
		for _, neighbor := range g[cur] {
			res1 := dfs(neighbor)
			if res1 > res {
				res = res1
			}
		}
		return informTime[cur] + res
	}
	return dfs(headID)
}

// https://leetcode.cn/problems/binary-string-with-substrings-representing-1-to-n/
func queryString(S string, N int) bool {
	for i := 1; i <= N; i++ {
		if !strings.Contains(S, strconv.FormatInt(int64(i), 2)) {
			return false
		}
	}
	return true
}

// https://leetcode.cn/problems/distant-barcodes/
func rearrangeBarcodes(barcodes []int) []int {
	mmap := make(map[int]int)
	for _, v := range barcodes {
		mmap[v]++
	}
	var res []int
	for len(mmap) > 0 {
		var max, maxK int
		for k, v := range mmap {
			if v > max {
				max = v
				maxK = k
			}
		}
		delete(mmap, maxK)
		for i := 0; i < max; i++ {
			res = append(res, maxK)
		}
	}
	return res
}

// https://leetcode.cn/problems/flip-columns-for-maximum-number-of-equal-rows/
func maxEqualRowsAfterFlips(matrix [][]int) int {
	mmap := make(map[string]int)
	for _, row := range matrix {
		var s []int
		for _, v := range row {
			s = append(s, v^row[0])
		}
		var str strings.Builder
		for _, v := range s {
			str.WriteString(strconv.Itoa(v))
		}
		mmap[str.String()]++
	}
	var res int
	for _, v := range mmap {
		if v > res {
			res = v
		}
	}
	return res
}

// https://leetcode.cn/problems/statistics-from-a-large-sample/
// 输入：count = [0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
// 输出：[1.00000,3.00000,2.37500,2.50000,3.00000]
func sampleStats(count []int) []float64 {
	n := len(count)
	total := 0
	for i := 0; i < n; i++ {
		total += count[i]
	}
	mean := 0.0
	median := 0.0
	minimum := 256
	maxnum := 0
	mode := 0

	left := (total + 1) / 2
	right := (total + 2) / 2
	cnt := 0
	maxfreq := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum += int(count[i]) * int(i)
		if count[i] > maxfreq {
			maxfreq = count[i]
			mode = i
		}
		if count[i] > 0 {
			if minimum == 256 {
				minimum = i
			}
			maxnum = i
		}
		if cnt < right && cnt+count[i] >= right {
			median += float64(i)
		}
		if cnt < left && cnt+count[i] >= left {
			median += float64(i)
		}
		cnt += count[i]
	}
	mean = float64(sum) / float64(total)
	median = median / 2.0
	return []float64{float64(minimum), float64(maxnum), mean, median, float64(mode)}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/delete-nodes-and-return-forest/
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	mmap := make(map[int]bool)
	for _, v := range to_delete {
		mmap[v] = true
	}
	var res []*TreeNode
	var dfs func(*TreeNode, bool) *TreeNode
	dfs = func(node *TreeNode, isRoot bool) *TreeNode {
		if node == nil {
			return nil
		}
		deleted := false
		if mmap[node.Val] {
			deleted = true
		}
		if isRoot && !deleted {
			res = append(res, node)
		}
		node.Left = dfs(node.Left, deleted)
		node.Right = dfs(node.Right, deleted)
		if deleted {
			return nil
		}
		return node
	}
	dfs(root, true)
	return res
}

// https://leetcode.cn/problems/minimum-cost-tree-from-leaf-values/
func mctFromLeafValues(arr []int) int {
	n := len(arr)
	var res int
	var stack []int
	stack = append(stack, math.MaxInt32)
	for i := 0; i < n; i++ {
		for stack[len(stack)-1] <= arr[i] {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res += tmp * min(stack[len(stack)-1], arr[i])
		}
		stack = append(stack, arr[i])
	}
	for len(stack) > 2 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res += tmp * stack[len(stack)-1]
	}
	return res
}

// https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	left, right := 0, price[len(price)-1]-price[0]
	for left < right {
		mid := (left + right + 1) / 2
		if check(price, k, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func check(price []int, k int, tastiness int) bool {
	prev := int(math.Inf(-1)) >> 1
	cnt := 0
	for _, p := range price {
		if p-prev >= tastiness {
			cnt++
			prev = p
		}
	}
	return cnt >= k
}

// https://leetcode.cn/problems/count-vowel-strings-in-ranges/
func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	prefixSums := make([]int, n+1)
	for i := 0; i < n; i++ {
		value := 0
		if checkStringHasPrefixAndSuffixVowel(words[i]) {
			value = 1
		}
		prefixSums[i+1] = prefixSums[i] + value
	}
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		start := queries[i][0]
		end := queries[i][1]
		ans[i] = prefixSums[end+1] - prefixSums[start]
	}
	return ans
}

func checkStringHasPrefixAndSuffixVowel(w string) bool {
	if len(w) == 0 {
		return false
	}
	mmap := make(map[byte]bool)
	mmap['a'] = true
	mmap['e'] = true
	mmap['i'] = true
	mmap['o'] = true
	mmap['u'] = true
	return mmap[w[0]] && mmap[w[len(w)-1]]
}

// https://leetcode.cn/problems/apply-operations-to-an-array/
func applyOperations(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
	}
	res := make([]int, n)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			res[j] = nums[i]
			j++
		}
	}
	return res
}

// https://leetcode.cn/problems/number-of-unequal-triplets-in-array/
func unequalTriplets(nums []int) int {
	total := 0
	mmap := make(map[int]int)
	for _, num := range nums {
		mmap[num]++
	}
	if len(mmap) <= 2 {
		return 0
	}
	n := len(nums)
	t := 0
	for _, v := range mmap {
		total, t = total+t*v*(n-t-v), t+v
	}
	return total
}

// https://leetcode.cn/problems/number-of-times-binary-string-is-prefix-aligned/
// 1375
func numTimesAllBlue(flips []int) int {
	res, right := 0, 0
	for i, flip := range flips {
		//right = max(right, flip)
		if flip > right {
			right = flip
		}
		if right == i+1 {
			res++
		}
	}
	return res
}

// https://leetcode.cn/problems/can-make-palindrome-from-substring/
//
//	1177
func canMakePaliQueries(s string, queries [][]int) []bool {
	sum := make([][26]int, len(s)+1)
	for i, c := range s {
		sum[i+1] = sum[i]
		sum[i+1][c-'a']++
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		left, right, k, m := q[0], q[1], q[2], 0
		for j := 0; j < 26; j++ {
			m += (sum[right+1][j] - sum[left][j]) % 2
		}
		res[i] = m/2 <= k
	}
	return res
}

// https://leetcode.cn/problems/reconstruct-a-2-row-binary-matrix/
func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	res := make([][]int, 2)
	res[0] = make([]int, n)
	res[1] = make([]int, n)
	for i := 0; i < n; i++ {
		if colsum[i] == 2 {
			res[0][i] = 1
			res[1][i] = 1
			upper--
			lower--
		}
	}
	for i := 0; i < n; i++ {
		if colsum[i] == 1 {
			if upper > 0 {
				res[0][i] = 1
				upper--
			} else {
				res[1][i] = 1
				lower--
			}
		}
	}
	if upper != 0 || lower != 0 {
		return [][]int{}
	}
	return res
}

// https://leetcode.cn/problems/maximum-alternating-subsequence-sum/
func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	dp := make([][2]int64, n)
	dp[0][0] = int64(nums[0])
	for i := 1; i < n; i++ {
		dp[i][0] = max2(dp[i-1][0], dp[i-1][1]+int64(nums[i]))
		dp[i][1] = max2(dp[i-1][1], dp[i-1][0]-int64(nums[i]))
	}
	return dp[n-1][0]
}

func max2(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// https://leetcode.cn/problems/alternating-digit-sum/
func alternateDigitSum(n int) int {
	total := 0
	var s1, s2 int
	for n > 0 {
		total++
		if total%2 == 1 {
			s1 += n % 10
		} else {
			s2 += n % 10
		}
		n /= 10
	}
	if total%2 == 1 {
		return s1 - s2
	}
	return s2 - s1
}

// https://leetcode.cn/problems/minimum-falling-path-sum/
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[0][i] = matrix[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i-1][j] + matrix[i][j]
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+matrix[i][j])
			}
			if j < n-1 {
				dp[i][j] = min(dp[i][j], dp[i-1][j+1]+matrix[i][j])
			}
		}
	}
	res := dp[n-1][0]
	for _, v := range dp[n-1] {
		res = min(res, v)
	}
	return res
}

// https://leetcode.cn/problems/number-of-ways-to-buy-pens-and-pencils/?envType=daily-question&envId=2023-09-01
// 给你一个整数 total ，表示你拥有的总钱数。同时给你两个整数 cost1 和 cost2 ，分别表示一支钢笔和一支铅笔的价格。你可以花费你部分或者全部的钱，去买任意数目的两种笔。
// 请你返回购买钢笔和铅笔的 不同方案数目 。
func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	if cost1 < cost2 {
		return waysToBuyPensPencils(total, cost2, cost1)
	}
	var res, cnt int
	for cnt*cost1 <= total {
		res += (total-cnt*cost1)/cost2 + 1
		cnt++
	}
	return int64(res)
}

// https://leetcode.cn/problems/maximum-enemy-forts-that-can-be-captured/?envType=daily-question&envId=2023-09-02
// 1-自己控制的城堡 0-敌人  -1-空地
func captureForts(forts []int) int {
	res, pre := 0, -1
	for i, v := range forts {
		if v == 1 || v == -1 {
			if pre >= 0 && forts[pre] != v {
				if i-pre-1 > res {
					res = i - pre - 1
				}
			}
			pre = i
		}
	}
	return res
}

// https://leetcode.cn/problems/eliminate-maximum-number-of-monsters/?envType=daily-question&envId=2023-09-03
func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = (dist[i]-1)/speed[i] + 1
	}
	sort.Ints(t)
	for i := 0; i < n; i++ {
		if t[i] <= i {
			return i
		}
	}
	return n
}

// https://leetcode.cn/problems/form-smallest-number-from-two-digit-arrays/?envType=daily-question&envId=2023-09-05
func minNumber(nums1 []int, nums2 []int) int {
	mmap := make(map[int]bool)
	min1, min2 := nums1[0], nums2[0]
	for _, v := range nums1 {
		mmap[v] = true
		if v < min1 {
			min1 = v
		}
	}
	var same []int
	for _, v := range nums2 {
		if mmap[v] {
			same = append(same, v)
		}
		if v < min2 {
			min2 = v
		}
	}
	if len(same) > 0 {
		sort.Ints(same)
		return same[0]
	}
	if min1 < min2 {
		return min1*10 + min2
	}
	return min2*10 + min1
}

// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/?envType=daily-question&envId=2023-09-06
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, res := lcaDeepestLeavesHelper(root, 0)
	return res
}

func lcaDeepestLeavesHelper(root *TreeNode, depth int) (int, *TreeNode) {
	if root == nil {
		return depth, nil
	}
	leftDepth, leftNode := lcaDeepestLeavesHelper(root.Left, depth+1)
	rightDepth, rightNode := lcaDeepestLeavesHelper(root.Right, depth+1)
	if leftDepth == rightDepth {
		return leftDepth, root
	}
	if leftDepth > rightDepth {
		return leftDepth, leftNode
	}
	return rightDepth, rightNode
}

// https://leetcode.cn/problems/minimum-time-to-repair-cars/?envType=daily-question&envId=2023-09-07
// 给你一个整数数组 ranks ，表示一些机械工的 能力值 。ranksi 是第 i 位机械工的能力值。能力值为 r 的机械工可以在 r * n2 分钟内修好 n 辆车。
// 同时给你一个整数 cars ，表示总共需要修理的汽车数目。
// 请你返回修理所有汽车 最少 需要多少时间。
// 注意：所有机械工可以同时修理汽车。
func repairCars(ranks []int, cars int) int64 {
	l, r := 1, ranks[0]*cars*cars
	var check = func(m int) bool {
		cnt := 0
		for _, x := range ranks {
			cnt += int(math.Sqrt(float64(m / x)))
		}
		return cnt >= cars
	}
	for l < r {
		m := (l + r) >> 1
		if check(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return int64(l)
}

// https://leetcode.cn/problems/course-schedule/?envType=daily-question&envId=2023-09-09
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(int)
	)
	dfs = func(i int) {
		visited[i] = 1
		for _, v := range edges[i] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[i] = 2
		result = append(result, i)
	}
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// https://leetcode.cn/problems/camelcase-matching/solutions/2224532/tuo-feng-shi-pi-pei-by-leetcode-solution-pwq7/?envType=daily-question&envId=2023-09-10
func camelMatch(queries []string, pattern string) []bool {
	n := len(queries)
	res := make([]bool, n)
	for i := 0; i < n; i++ {
		res[i] = true
		p := 0
		for _, c := range queries[i] {
			if p < len(pattern) && pattern[p] == byte(c) {
				p++
			} else if unicode.IsUpper(c) {
				res[i] = false
				break
			}
		}
		if p < len(pattern) {
			res[i] = false
		}
	}
	return res
}

// https://leetcode.cn/problems/course-schedule-iii/?envType=daily-question&envId=2023-09-11
func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := &Heap{}
	total := 0 // 优先队列中所有课程的总时间
	for _, course := range courses {
		if t := course[0]; total+t <= course[1] {
			total += t
			heap.Push(h, t)
		} else if h.Len() > 0 && t < h.IntSlice[0] {
			total += t - h.IntSlice[0]
			h.IntSlice[0] = t
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// https://leetcode.cn/problems/course-schedule-iv/?envType=daily-question&envId=2023-09-12
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	g := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	isPre := make([][]bool, numCourses)
	for i, _ := range isPre {
		isPre[i] = make([]bool, numCourses)
		g[i] = []int{}
	}
	for _, p := range prerequisites {
		indegree[p[1]]++
		g[p[0]] = append(g[p[0]], p[1])
	}
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, ne := range g[cur] {
			isPre[cur][ne] = true
			for i := 0; i < numCourses; i++ {
				isPre[i][ne] = isPre[i][ne] || isPre[i][cur]
			}
			indegree[ne]--
			if indegree[ne] == 0 {
				q = append(q, ne)
			}
		}
	}
	var res []bool
	for _, query := range queries {
		res = append(res, isPre[query[0]][query[1]])
	}
	return res
}

func main() {
	fmt.Println(captureForts([]int{1, 0, 0, -1, 0, 0, 0, 0, 1}))
	//maxEqualRowsAfterFlips([][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}})
	//words := []string{"bzmxvzjxfddcuznspdcbwiojiqf", "mwguoaskvramwgiweogzulcinycosovozppl", "uigevazgbrddbcsvrvnngfrvkhmqszjicpieahs", "uivcdsboxnraqpokjzaayedf", "yalc", "bbhlbmpskgxmxosft", "vigplemkoni", "krdrlctodtmprpxwditvcps", "gqjwokkskrb", "bslxxpabivbvzkozzvdaykaatzrpe", "qwhzcwkchluwdnqjwhabroyyxbtsrsxqjnfpadi", "siqbezhkohmgbenbkikcxmvz", "ddmaireeouzcvffkcohxus", "kjzguljbwsxlrd", "gqzuqcljvcpmoqlnrxvzqwoyas", "vadguvpsubcwbfbaviedr", "nxnorutztxfnpvmukpwuraen", "imgvujjeygsiymdxp", "rdzkpk", "cuap", "qcojjumwp", "pyqzshwykhtyzdwzakjejqyxbganow", "cvxuskhcloxykcu", "ul", "axzscbjajazvbxffrydajapweci"}
	//queries := [][]int{{4, 4}, {6, 17}, {10, 17}, {9, 18}, {17, 22}, {5, 23}, {2, 5}, {17, 21}, {5, 17}, {4, 8}, {7, 17}, {16, 19}, {7, 12}, {9, 20}, {13, 23}, {1, 5}, {19, 19}}
	//fmt.Println(vowelStrings(words, queries))
}
