package LC1895

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/6 上午1:02
 * @FilePath: find_sum_pairs
 * @Description: 使用hash实现该能力
 */

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	mp    map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	mp := make(map[int]int)
	for _, num := range nums2 {
		mp[num]++
	}
	return FindSumPairs{nums1: nums1, nums2: nums2, mp: mp}
}

func (this *FindSumPairs) Add(index int, val int) {
	old := this.nums2[index]
	this.mp[old]--
	this.mp[old+val]++
	this.nums2[index] += val
}

func (this *FindSumPairs) Count(tot int) int {
	var ans int
	for _, num := range this.nums1 {
		if t, ok := this.mp[tot-num]; ok {
			ans += t
		}
	}
	return ans
}
