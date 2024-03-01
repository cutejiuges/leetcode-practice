package lc2369

import "testing"

func TestValidPartition(t *testing.T) {
	testData := map[string][]int{
		"case1": {4, 4, 4, 5, 6},
		"case2": {1, 1, 1, 2},
	}
	expectRes := map[string]bool{
		"case1": true,
		"case2": false,
	}
	for key, val := range testData {
		caseName := key
		caseData := val
		t.Run(caseName, func(t *testing.T) {
			res := validPartition(caseData)
			if res != expectRes[caseName] {
				t.Error(caseData, "结果不符合预期")
			}
		})
	}
}
