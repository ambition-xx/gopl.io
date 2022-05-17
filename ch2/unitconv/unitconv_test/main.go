package main

import (
	"bufio"
	"fmt"

	
	"os"
	"strconv"

	"gopl.io/ch2/unitconv"
)

func main() {
	var input []string    //定义一个slice
	if len(os.Args) > 1 { //从命令行输入
		input = os.Args[1:] //第一个为文件地址，从第一个元素开始取
	} else {
		stdInput := bufio.NewScanner(os.Stdin) //从标准输入流输入
		for stdInput.Scan() {
			input = append(input, stdInput.Text())
		}
	}
	for _, arg := range input {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid input %v\n", err)
			os.Exit(1)
		}
		f := unitconv.Fahrenheit(t)
		c := unitconv.Celsius(t)
		p := unitconv.Pound(t)
		kg := unitconv.Kilogram(t)
		fe := unitconv.Inch(t)
		m := unitconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FToC(f), c, unitconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", p, unitconv.PToK(p), kg, unitconv.KToP(kg))
		fmt.Printf("%s = %s, %s = %s\n", fe, unitconv.IToM(fe), m, unitconv.MToI(m))
	}

}
