package code

// 全排列问题
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	visited := make(map[int]bool)

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(nums) == len(path) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for _, v := range nums {
			if visited[v] {
				continue
			}

			path = append(path, v)
			visited[v] = true
			dfs(path)
			path = path[:len(path) - 1]
			visited[v] = false
		}
	}

	dfs([]int{})
	return res
}