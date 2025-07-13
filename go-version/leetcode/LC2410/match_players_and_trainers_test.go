package LC2410

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/13 下午9:20
 * @FilePath: match_players_and_trainers_test
 * @Description: 对数组进行排序之后遍历比较即可
 */

type TestData map[string]struct {
	Players   []int `yaml:"players"`
	Trainers  []int `yaml:"trainers"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestMatchPlayersAndTrainers(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := matchPlayersAndTrainers(caseData.Players, caseData.Trainers); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
