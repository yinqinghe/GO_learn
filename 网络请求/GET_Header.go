package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var proxyUrl, _ = url.Parse("http://127.0.0.1:8080")

var transport = &http.Transport{
	Proxy: http.ProxyURL(proxyUrl), //设置代理发送请求
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true, //通过设置Transport的TLSClientConfig来忽略证书验证:
	},
}

var client = http.Client{
	Timeout:   time.Second * 10,
	Transport: transport,
}

func main() {
	url := "https://www.baidu.com"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// 处理错误
		return
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	resp, err := client.Do(req)
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

	// fmt.Println(string(body))
	fmt.Println(len(body))
}
