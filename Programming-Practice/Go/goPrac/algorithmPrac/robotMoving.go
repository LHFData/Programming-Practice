package main

import (
	"container/list"
	"fmt"
)
func legalSquare(i int,j int,k int)bool{
	isum:=0
	for i%10>10{
		isum=isum+i%10
		i=(i-i%10)/10
	}
	isum=isum+i
	jsum:=0
	for j%10>10{
		jsum=jsum+j%10
		j=(j-j%10)/10
	}
	jsum=jsum+j
	if jsum+isum>k{
		return false
	}
	return true
}
func movingCount(m int, n int, k int) int {
	visit:=make([][]bool,m)
	for i,_:=range visit{
		visit[i]=make([]bool,n)
	}
	bfs:=func(v [][]bool) int{
		var count int
		queue:=list.New()
		queue.PushBack([2]int{0,0})
		move:=[][2]int{{1,0},{-1,0},{0,1},{0,-1}}
		for queue.Len()>0{
			count++
			temp:=queue.Remove(queue.Front()).([2]int)
			for _,val:=range move{
				x,y:=temp[0]+val[0],temp[1]+val[1]
				if x>=0&&x<len(visit)&&y>=0&&y<len(visit[0]){
					if visit[x][y]{
						queue.PushBack([2]int{x,y})
					}
				}
			}
		}
		return count
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if legalSquare(i,j,k){
				visit[i][j]=true
			}
		}
	}
	return bfs(visit)
}

func main(){
	fmt.Println(movingCount(5,4,0))
}
