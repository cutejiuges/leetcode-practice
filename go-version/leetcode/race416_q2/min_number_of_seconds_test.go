package race416_q2

import (
	"fmt"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/22 上午11:35
 * @FilePath: min_number_of_seconds_test
 * @Description:
 */

func TestMinNumberOfSeconds(t *testing.T) {
	fmt.Println(minNumberOfSeconds(4, []int{2, 1, 1}))
	fmt.Println(minNumberOfSeconds(10, []int{3, 2, 2, 4}))
}
