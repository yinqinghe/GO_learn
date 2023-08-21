package main

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	args := os.Args //GO 获取命令行输入的内容

	// Check if any arguments were provided
	if len(args) > 1 {
		// Print the arguments
		fmt.Println("Command-line arguments:")
		for i, arg := range args[1:] {
			fmt.Printf("%d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("No command-line arguments provided.")
	}

	// Create command

	// cmd := exec.Command("sh", "-c", cmdStr)   //在LInux执行一个命令字符串
	// cmd := exec.Command("whoami")
	// cmd := exec.Command("ipconfig", "/all")

	cmdStr := "ipconfig /all"

	cmd := exec.Command("cmd", "/C", cmdStr) //在windows里执行一个命令字符串

	// Execute command and capture output
	out, err := cmd.CombinedOutput()
	// output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//可以解决命令执行结果乱码问题
	output, _ := simplifiedchinese.GB18030.NewDecoder().String(string(out)) //转换字符集，解决UTF-8乱码

	// Print output
	fmt.Println(string(output))
}
