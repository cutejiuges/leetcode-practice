package lc1287

func findSpecialInteger(arr []int) int {
	n := len(arr)
	num, cnt := arr[0], 0
	for i := 0; i < n; i++ {
		if num == arr[i] {
			cnt++
			if cnt*4 > n {
				return num
			}
		} else {
			num = arr[i]
			cnt = 1
		}
	}
	return -1
}
