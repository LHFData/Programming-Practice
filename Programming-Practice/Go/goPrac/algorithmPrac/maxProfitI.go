package main

import "fmt"

func maxProfit(prices []int) int {
	//最低点是可以更新的，因为有先后顺序，所以即使前面有更低的价格，由于在前面的时间无法获得更大收益
	//在遇到更低的价格时，倘若其能创造更大收益，则不会被替换，反之则应该由更低的价格替换
	//假设前面出现了次小价格，前面出现的是次大高峰，在最低和最高出现之前，其收益也是次大
	//前面出现了次小价格和最大高峰，在最低出现之时，其被替换为最低，但收益是否被替换取决于最小能否带来更大收益
	//能，则收益被替换，不能被替换意味着最小价格不产生影响
	minPrice:=prices[0]
	profit:=0

	for i:=0;i<len(prices);i++{
		if prices[i]<minPrice{
			minPrice=prices[i]

		}
		if(prices[i]-minPrice>profit){
			profit=prices[i]-minPrice
		}

	}
	return profit
}
func main()  {
	p:=[]int{3,1,4,5,2,6}
	fmt.Println(maxProfit(p))
}