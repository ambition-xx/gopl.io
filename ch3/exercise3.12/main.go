package main

import (
	"fmt"
)

func main() {
	a, b := "aaabbb", "bbbaaa"
	fmt.Println(isAnagram(a, b))
}

func isAnagram(a, b string) bool{
	mp := make(map[rune]int)
	for _, v := range a {
		mp[v]++
	}
	for _, v := range b{
		mp[v]--
	}
	for _, v := range mp{
		if v!=0{
			return false
		}
	}
	return true
}
