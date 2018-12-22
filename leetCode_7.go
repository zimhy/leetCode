package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if float64(x) >= math.Pow(2, float64(31)) || float64(x) < -math.Pow(2, float64(31)) {
		return 0
	}
	num := int64(x)
	isPostive := x > 0
	if !isPostive {
		num = -num
	}
	splited := make([]int64, 0)
	for {
		if num > 0 {
			splited = append(splited, (num % 10))
			num = num / 10
		} else {
			break
		}
	}
	num = 0
	for i := 0; i < len(splited); i++ {
		si := splited[i]
		posi := math.Pow10(len(splited) - i - 1)
		tmpnum := si * int64(posi)

		num += tmpnum
	}
	if !isPostive {
		num = -num
	}
	if float64(num) >= math.Pow(2, float64(31)) || float64(num) < -math.Pow(2, float64(31)) {
		return 0
	}
	return int(num)
}

func main() {
	//fmt.Println(math.MaxInt32,math.MinInt32)
	fmt.Println(reverse(1534236469))
	//1534236469
	//2147483412
}
