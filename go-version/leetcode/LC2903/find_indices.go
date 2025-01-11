package LC2903

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/19 下午11:31
 * @FilePath: find_indices
 * @Description: 双指针
 */

func findIndices(nums []int, indexDifference, valueDifference int) []int {
	maxIndex, minIndex := 0, 0
	for j := indexDifference; j < len(nums); j++ {
		i := j - indexDifference
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		} else if nums[i] < nums[minIndex] {
			minIndex = i
		}

		if nums[maxIndex]-nums[j] >= valueDifference {
			return []int{maxIndex, j}
		}
		if nums[j]-nums[minIndex] >= valueDifference {
			return []int{minIndex, j}
		}
	}
	return []int{-1, -1}
}
