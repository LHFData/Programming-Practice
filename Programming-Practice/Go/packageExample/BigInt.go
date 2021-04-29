package main

import (
	"fmt"
	"math/big"

)
func main(){
	t:=big.NewInt(1)
	s:=big.NewInt(2)
	ss:=[]*big.Int{t,s}
	for i:=0;i<len(ss);i++{
		fmt.Println(ss[i])
	}
	max:=func(a *big.Int,b *big.Int)*big.Int{
		if a.Cmp(b)==1{
			return a
		}
		return b
	}
	fmt.Println(max(t,s))
	fmt.Println(s.Mul(t,s))
	fmt.Println(s.Sub(t,s))
}


