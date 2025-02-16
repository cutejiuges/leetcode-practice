package lc1299

func replaceElements(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	mx := -1
	for i := n - 1; i >= 0; i-- {
		ans[i] = mx
		mx = max(mx, arr[i])
	}
	return ans
}
