package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.example.com/api"
	data := strings.NewReader(`{"name": "John", "age": 30}`)

	resp, err := http.Post(url, "application/json", data)
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

	fmt.Println(string(body))
}
