
package main

import (
"fmt"
)


func main(){
	n,k:=0,0
	fmt.Scan(&n)
	fmt.Scan(&k)
	var relation [][3]int
	for i:=0;i<n-1;i++{
		var tmp [3]int
		fmt.Scan(&tmp[0])
		fmt.Scan(&tmp[1])
		fmt.Scan(&tmp[2])
		relation=append(relation,tmp)
	}
	var res int
	for i:=0;i<n-1;i++{
		if relation[i][2]<k{
			res++
		}
	}
}
