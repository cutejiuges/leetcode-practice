package LC661

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/18 下午10:56
 * @FilePath: image_smoother_test
 * @Description:
 */

type TestData map[string]struct {
	Img       [][]int `yaml:"img"`
	ExpectRes [][]int `yaml:"expect_res"`
}

func TestImageSmoother(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := imageSmoother(caseData.Img)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
