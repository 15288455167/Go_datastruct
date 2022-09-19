package main

import "fmt"

func mostFrequentEven(nums []int)  {
	ret := make(map[int]int)
	for _, v := range nums{
		if v % 2 != 0 {
			continue
		}
		ret[v]++
	}

	res := 0
	for i := range ret {
		if ok := i; ok {

		}
	}

	fmt.Println(res)
}



func main() {
	nums := []int{0,1,2,2,4,4,1}
	mostFrequentEven(nums)
}
