package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file_path := "D:\\C#\\Go-code\\leanr_GO\\文件操作\\file\\url.txt"
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.SplitAfter(string(data), "\n")

	//for i, line := range lines {
	//	fmt.Printf("%d: %s\n", i+1, line)
	//}
	// 忽略前 9 行
	start := 9
	for i, line := range lines[start:15] { //切片的作用
		// 重新计算行号
		lineNo := i + start + 1
		fmt.Printf("%d: %s\n", lineNo, line)
	}
	//读取文件所有内容到 data
	//利用 strings.SplitAfter() 按换行分割成 lines 数组
	//使用 lines[start:] 忽略前 start=9 行
	//为lines[start:]重新计算从 start+1 开始的行号
	//遍历 lines[start:] 每行,打印带行号的内容
}
