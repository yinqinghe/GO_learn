package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	routineCountTotal = 5 //限制线程数
)

func main() {
	file_path := "D:\\C#\\Go-code\\leanr_GO\\文件操作\\file\\url.txt"
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.SplitAfter(string(data), "\r\n")

	client = &http.Client{}
	beg := time.Now()
	wg := &sync.WaitGroup{}
	tasks := make(chan [2]string)
	results := make(chan string)
	//receiver接受响应并处理的函数块, 也可以单独写在一个函数
	//go func() {
	//	for result := range results {
	//		if result == "" {
	//			close(results)
	//		} else {
	//			fmt.Println("result:", result)
	//		}
	//	}
	//}()
	for i := 0; i < routineCountTotal; i++ {
		wg.Add(1)
		go worker(wg, tasks, results)
	}
	start := 25
	end := len(lines)
	//分发任务
	for i, task := range lines[start:end] {
		lineNo := i + start + 1

		tasks <- [2]string{task, strconv.Itoa(lineNo)}
	}
	tasks <- [2]string{} //worker结束标志
	wg.Wait()            //同步结束
	results <- ""        // result结束标志
	fmt.Printf("time consumed: %fs", time.Now().Sub(beg).Seconds())
}

func worker(group *sync.WaitGroup, tasks chan [2]string, result chan string) {
	for task := range tasks {
		if len(task) == 0 { //如果task为空  则关闭taskts
			close(tasks)
		} else {
			NumberQueryRequest(task)
			//result <- string(respBody)

			//if err != nil {
			//	fmt.Printf("error occurred in NumberQueryRequest: %s\n", task)
			//	result <- err.Error()
			//} else {
			//	result <- string(respBody)
			//}
		}
	}
	group.Done()
}

var client = &http.Client{
	Timeout: time.Second * 10,
}

func NumberQueryRequest(t [2]string) {
	//arr := <-t
	ip_ := t[0]
	//url := ip_ + "/general/mytable/intel_view/video_file.php?MEDIA_DIR=../../../inc/&MEDIA_NAME=oa_config.php"
	url := ip_
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// 解码gbk
	decoder := simplifiedchinese.GBK.NewDecoder()
	result, _ := decoder.Bytes(body)

	//fmt.Println(string(result))
	//fmt.Println(string(body))
	if resp.StatusCode == 200 {
		println("响应码为：", resp.StatusCode)
	}
	fmt.Println(resp.StatusCode, len(body))
	if strings.Contains(string((result)), "北京创信达科技有限公司") {
		println("存在sb")
		println(t[1], ip_)
		//println("存在baidu")
	}
}
