package main

import "fmt"

func main() {
	ch := make(chan [2]string)

	go func() {
		ch <- [2]string{"a", "b"}
	}()

	arr := <-ch
	fmt.Println(arr[0]) // a
	fmt.Println(arr[1]) // b
}
