package main

import (
	"fmt"
	"math/big"
)
func cuttingRopeII(n int) int {
	dp:=make([]*big.Int,n+1)
	dp[0]=big.NewInt(0)
	dp[1]=big.NewInt(1)
	dp[2]=big.NewInt(1)
	modNum:=big.NewInt(1000000007)
	max:=func(a *big.Int,b *big.Int)*big.Int{
		if a.Cmp(b)==1{
			return a
		}
		return b
	}
	for i:=3;i<=n;i++{
		for j:=2;j<=i;j++{
			bigI:=big.NewInt(int64(i))
			bigJ:=big.NewInt(int64(j))
			temp:=max(bigJ.Mul(bigJ,bigI.Sub(bigI,bigJ)),bigJ.Mul(bigJ,dp[i-j]))
			dp[i]=max(dp[i],temp)
		}
	}
	return int(dp[n].Mod(dp[n],modNum).Int64())
}
func main (){
	fmt.Println(cuttingRopeII(10))
}
