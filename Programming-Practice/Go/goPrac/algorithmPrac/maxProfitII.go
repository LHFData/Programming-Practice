package main

import "fmt"

func maxProfitII(prices []int) int {
	//动态规划
	//每天状态按持股与否，划分为持有与没持有
	//持有是因为前面就有，或者是前一天才买入
	//未持有是因为本来就没持有，或者是今天刚卖出
	//需要注意的是当天持有的情况，若是当天买入，则意味着是前一天未持有的价值减去今天的股价
	//或许需要具象一个不存在的无限账户，才方便考虑两种结果中后面的情况，买了就是当天的钱我给了，收盘的时候就是当天我按这个价钱卖出去了
	//dp:=[][]int{}

	//贪心法的角度就是开上帝视角，每一天只要和第二天是正的，那我就都要，可以理解为最终结果是无数个隔天交易的集合，我只吃第二天涨的区间
	//因为如果你跌了始终要升的，跨过第二天毫无意义，我只要注意不产生高买低卖我一定会有最大的收益，这个前提建立在交易次数不限制
	max:=func(a int,b int)int{
		if a>b{
			return a
		}else{
			return b
		}
	}
	for i:=0;i<len(prices);i++{
		//dp[i]=make([]int,len(prices))
	}
	//dp[0][0]=0
	//dp[0][1]=0
	preSold:=0
	preBuy:=-prices[0]
	for i:=1;i<len(prices);i++{
		preSold=max(preSold,preBuy+prices[i])
		preBuy=max(preBuy,preSold-prices[i])
		//dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
		//dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])

	}
	return preSold
}

func main()  {
	p:=[]int{7,1,5,3,6,4}
	fmt.Println(maxProfitII(p))

}