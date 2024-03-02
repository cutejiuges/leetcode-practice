package lc2368

import "testing"

type TestData struct {
	N          int
	Edges      [][]int
	Restricted []int
	ExpectRes  int
}

func TestReachableNodes(t *testing.T) {
	testData := map[string]TestData{
		"case1": {
			N:          7,
			Edges:      [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}},
			Restricted: []int{4, 5},
			ExpectRes:  4,
		},
		"case2": {
			N:          7,
			Edges:      [][]int{{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5}},
			Restricted: []int{4, 2, 1},
			ExpectRes:  3,
		},
	}

	for key, val := range testData {
		caseName := key
		caseData := val
		t.Run(caseName, func(t *testing.T) {
			res := reachableNodes(caseData.N, caseData.Edges, caseData.Restricted)
			if res != caseData.ExpectRes {
				t.Errorf("%s, res = %d, ExpectRes = %d", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
