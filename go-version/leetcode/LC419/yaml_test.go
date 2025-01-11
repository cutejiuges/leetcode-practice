package LC419

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/11 上午12:32
 * @FilePath: yaml_test
 * @Description:
 */
func TestPrintYaml(t *testing.T) {
	data := [][]byte{
		{'X', '.'},
		{'.', 'X'},
	}
	ydata, err := yaml.Marshal(data)
	fmt.Println(err)
	if err == nil {
		fmt.Println(string(ydata))
	}
}
