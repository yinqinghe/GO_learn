package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://www.baidu.com",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=baidu"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("http://www.baidu.com",
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.baidu.com", strings.NewReader("name=baidu"))
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=baidu")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

func main() {
	httpGet()
	httpPost()
	httpPostForm()
	httpDo()
}
