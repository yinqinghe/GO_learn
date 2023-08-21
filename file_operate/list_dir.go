package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ListDir(dirPth string) error { //列出目录下的文件和文件夹
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return err
	}

	for _, fi := range dir {
		if fi.IsDir() { //忽略目录
			fmt.Println("目录：" + fi.Name())
		} else {
			fmt.Println("文件" + fi.Name())
		}
	}
	return nil
}
func createDir(path string, dirName string) { //创建文件夹
	dirPath := path + "\\" + dirName

	err := os.Mkdir(dirPath, 0777)
	if err != nil {
		fmt.Println(err)
	}
	os.Chmod(dirPath, 0777)
	fmt.Println("Create Dir =>" + path + dirName)
}

func main() {
	//ListDir("C:\\users")
	createDir("D:\\C#\\Go-code\\leanr_GO", "test")
}
