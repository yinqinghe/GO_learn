// package main

// import "fmt"

// func main() {
// 	var s1 = []int{}

// 	s1 = append(s1, 500) //切片添加数据
// 	s1 = append(s1, 00)
// 	s1 = append(s1, 55)
// 	fmt.Printf("s1: %v\n", s1)

// 	var s2 = []int{1, 2, 3, 4, 5}
// 	s2 = append(s2[:2], s2[3:]...) //删除索引为2的数据
// 	//a = append(s2[:index], s2[index+1:]...)
// 	fmt.Printf("s2: %v\n", s2)

// 	//copy  是一块数据复制为另一块数据   赋值是共用一块数据的地址，如果通过一个变量修改了数据，
// 	//那么另一个数据访问数据时，是已经被修改过的数据
// 	var s3 = make([]int, len(s2))
// 	copy(s3, s2)
// 	fmt.Printf("s3: %v\n", s3)
// 	for _, v := range v {

// 	}

// }
