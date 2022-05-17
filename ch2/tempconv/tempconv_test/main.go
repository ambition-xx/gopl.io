package main

import (
	"fmt"

	tempconv "gopl.io/ch2/tempconv"
)

func main() {
	fmt.Println("AbsoluteZeroC:", tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Println("FreezingC:", tempconv.CToK(tempconv.FreezingC))
	fmt.Println("BoilingC:", tempconv.CToK(tempconv.BoilingC))
}
