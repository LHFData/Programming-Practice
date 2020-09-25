package main

import "fmt"

func reverse(x int) int {
	var r = x
	var single = []int{}
	for i := 0; r != 0; i++ {
		single = append(single, r%10)
		fmt.Print(single)
		r = r / 10
	}
	var output int = 0
	var start bool = false
	for _, v := range single {
		if v != 0 && !start {
			start = true
			output = v + 10*output
			continue
		}
		if start {
			output = v + 10*output
		}
	}
	if (output > 4294967295) || (output < (-4294967296)) {
		return 0
	}
	return output
}
