package lc225

import (
	"cutejiuge/leetcode-practice/util"
	"strconv"
	"testing"
)

type TestData map[string]struct {
	Operate   []string `yaml:"Operate"`
	Nums      [][]int  `yaml:"Nums"`
	ExpectRes []string `yaml:"ExpectRes"`
}

func TestMyStack(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			var stack MyStack
			for i, op := range caseData.Operate {
				switch op {
				case "MyStack":
					stack = Constructor()
				case "push":
					for _, num := range caseData.Nums[i] {
						stack.Push(num)
					}
				case "top":
					x := stack.Top()
					if strconv.Itoa(x) != caseData.ExpectRes[i] {
						t.Errorf("%s, top = %v, ExpectTop = %v", caseName, x, caseData.ExpectRes[i])
					}
				case "empty":
					e := stack.Empty()
					if strconv.FormatBool(e) != caseData.ExpectRes[i] {
						t.Errorf("%s, empty = %v, ExpectEmpty = %v", caseName, e, caseData.ExpectRes[i])
					}
				}
			}
		})
	}
}
