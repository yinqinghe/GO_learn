package main

import "fmt"

func one() (name string, age int) {
	name = "哦发呢顺丰"
	age = 18
	return
}

func main() {
	names, ages := one()
	fmt.Printf("%s %d", names, ages)
}
