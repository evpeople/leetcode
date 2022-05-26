package f26

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(Lremain int, rRemain int, path string)
	dfs = func(Lremain, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
		}
		if Lremain > 0 {
			dfs(Lremain-1, rRemain, path+"(")
		}
		if Lremain < rRemain {
			dfs(Lremain, rRemain-1, path+")")
		}
	}
	dfs(n, n, "")
	return res
}
