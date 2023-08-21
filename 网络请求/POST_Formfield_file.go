package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
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
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	// 打开要上传的文件
	file, err := os.Open("D:\\youxi\\0x7eTools\\Tools\\WebShell\\webshell_Manage\\Behinder\\Behinder_v4.0.6\\server\\mike.jsp")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// 添加表单字段
	_ = writer.WriteField("field1", "value1")
	_ = writer.WriteField("field2", "value2")

	// 添加文件
	fileWriter, err := writer.CreateFormFile("fileField", "filename.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = writer.Close()

	request, err := http.NewRequest("POST", "https://example.com", &buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)
}
