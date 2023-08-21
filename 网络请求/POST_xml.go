package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
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
	ip_ := "https://api.example.com/submit" // Replace with the actual API endpoint

	xmlData := `
	<Request>
		<Name>John Doe</Name>
		<Email>john.doe@example.com</Email>
	</Request>
	`

	body := bytes.NewBuffer([]byte(xmlData))

	// Create the HTTP request
	request, err := http.NewRequest("POST", ip_, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the content type header
	request.Header.Set("Content-Type", "application/xml")
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
