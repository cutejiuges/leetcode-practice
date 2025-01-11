package lc1216

import (
	datastruct "cutejiuge/leetcode-practice/data_struct"
	"fmt"
	"testing"
)

func TestFindElements(t *testing.T) {
	nodes := []string{"1", "2", "2", "3", "4", "4", "3"}
	root := datastruct.Construct(nodes)
	fmt.Println(datastruct.PreOrder(root))
	fmt.Println(datastruct.MiddleOrder(root))
	fmt.Println(datastruct.PostOrder(root))
}
