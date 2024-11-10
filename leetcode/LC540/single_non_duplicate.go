package LC540

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/10 下午9:04
 * @FilePath: single_non_duplicate
 * @Description:
 */

// 取巧的做法，使用异或运算
func singleNonDuplicate(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

// 通用的高效解法，充分利用有序的已知条件
/**
1. 已知数组中其他元素都出现2次，仅一个元素出现了1次，现在需要找出这个出现1次的元素，该元素的左侧必定有偶数个数字
2. 可以推论，该元素的下标一定是偶数，我们可以枚举偶数下标，判断nums[2k]和nums[2k+1]是否想等，如果不相等，那么nums[2k]就是答案
3. 充分利用数组有序的条件，对于目标元素左侧的数字，nums[2k] == nums[2k+1]，对于目标元素右侧的元素，nums[2k]<nums[2k+1]
4. 根据3得出原数组本身存在二段性，可以考虑使用二分的思路来进行计算，二分偶数下标，如果nums[2k] == nums[2k+1]，继续二分后半段，nums[2k] < nums[2k+1]，继续二分前半段
*/
func singleNonDuplicate2(nums []int) int {
	// 二分偶数位置，所以high要比元素数量少一个
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		//确保mid是偶数
		mid -= mid & 1
		if nums[mid] == nums[mid+1] {
			//如果是相等的，对后半段进行二分
			low = mid + 2
		} else {
			high = mid
		}
	}
	return nums[low]
}
