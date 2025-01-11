/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-02 21:19:58
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-02 23:00:02
 * @FilePath: /leetcode-practice/util/get_file_info.go
 * @Description: 工具脚本，获取当前正在运行的测试文件名称
 */
package util

import (
	"errors"
	"runtime"
	"strings"
)

func getFileInfo(skip int) (funcName, fileName string, err error) {
	pc, fileName, _, ok := runtime.Caller(skip)
	if !ok {
		return "", "", errors.New("runtime.Caller() failed")
	}
	funcName = runtime.FuncForPC(pc).Name()
	return
}

func GetFileName() (string, error) {
	_, fileName, err := getFileInfo(3)
	if err != nil {
		return "", err
	}
	caseFileName := strings.Replace(fileName, ".go", ".yaml", 1)
	return caseFileName, nil
}
