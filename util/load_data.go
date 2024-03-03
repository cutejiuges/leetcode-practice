/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-02 21:28:14
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-03 14:44:18
 * @FilePath: /leetcode-practice/util/load_data.go
 * @Description: 在此通过数据文件，将整个yaml测试数据读入进来
 */
package util

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func DataProvide(val interface{}) {
	fileName, err := GetFileName()
	if err != nil {
		panic("获取运行数据失败")
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("获取case数据失败")
		panic(err)
	}
	err = yaml.Unmarshal(data, val)
	if err != nil {
		fmt.Println("数据加载失败")
		panic(err)
	}
}
