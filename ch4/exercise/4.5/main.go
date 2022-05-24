// 练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
package main

import (
	"fmt"
)

func main() {
	strs := []string{"da", "da","d", "d", "d", "f","a","d"}
	fmt.Println(strs)
	unique(strs)
	// fmt.Println(strs)
}
func unique(strs []string) {
	for i := 0; i < len(strs); i++ {
		if i+1 < len(strs) && strs[i] == strs[i+1] {
			copy(strs[i:], strs[i+1:])
			strs = strs[:len(strs)-1]
			i--
		}

	}
	fmt.Println(strs)
}
