package main

import (
	"fmt"
	"os"
)

func ReadFile(path string) {
	//file, err := os.Open(path)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0776)
	//os.O_RDWR|os.O_CREATE  表示以读写方式打开文件，如果文件不存在，则创建这个文件

	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 1024)
	fmt.Println("以下是文件内容：")
	for {
		len, _ := file.Read(buf)
		if len == 0 {
			break
		}
		fmt.Println(string(buf))
	}
	file.Close()
}

// ReadAt从指定的位置（相对于文件开始位置）读取len(b)字节数据并写入byte数组b。
// 它返回读取的字节数和可能遇到的任何错误。当n＜len(b)时，本方法总是会返回错误；如果是因为到达文件结尾，返回值err会是io.EOF。

func ReadFile2(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 1024)
	fmt.Println("以下是文件内容：")
	_, _ = file.ReadAt(buf, 9) //从第10个字符读
	fmt.Println(string(buf))
	_ = file.Close()
}

func main() {
	file_path := "D:\\C#\\Go-code\\leanr_GO\\文件操作\\file\\url.txt"

	//ReadFile(file_path)
	ReadFile2(file_path)
}
