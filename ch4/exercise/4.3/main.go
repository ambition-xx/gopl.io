//数组指针版
package main

import (
	"fmt"
	"reflect"
)

func main()  {
	string := [...]int{1,2,3,4,5,6,7}
	fmt.Println(reflect.TypeOf(string))//检查变量类型
	reverse(&string)
	fmt.Println(string)

}

func reverse(s *[7]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
