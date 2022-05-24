//一次循环版
package main

import (
	"fmt"
	"reflect"
)

func main() {
	string := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(reflect.TypeOf(string[:])) //检查变量类型
	fmt.Println(string)

	fmt.Println(rotate(string, 2))

}

func rotate(s []int, n int) []int {
	for i := 0; i < n; i++ {
		s = append(s, s[0])
		s = s[1:]
		fmt.Println(s)
	}
	return s
}
