package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0{
		return 1
	}
	if n == 1{
		return x
	}
	if n<0{
		x = 1/x
		n = -n
	}
	ret:=1.0
	for n>=1{
		if n & 1 == 1{
			ret *= x
			n--
		}else{
			x *= x
			n = n >> 1
		}
	}
	return ret
}

func main()  {
	fmt.Println(myPow(0.00001,
	2147483647))
}
