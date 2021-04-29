package main



import "fmt"

func main(){
	T:=0
	fmt.Scan(&T)
	GameNum:=0
	fmt.Scan(&GameNum)
	var time [][]int
	var score [][]int
	for t:=0;t<T;t++{
		var ti []int
		for i:=0;i<GameNum;i++{
			temp:=0
			fmt.Scan(&temp)
			ti=append(ti,temp)
		}
		time=append(time,ti)
		var sc []int
		for i:=0;i<GameNum;i++{
			temp:=0
			fmt.Scan(&temp)
			sc=append(sc,temp)
		}
		score=append(score,sc)
	}
	m:=make(map[int]int)

	for t:=0;t<T;t++{
		count:=0
		for i:=0;i<GameNum;i++{
			temp,ok:=m[time[t][i]]
			if ok{
				if temp<score[t][i]{
					count=count+m[time[t][i]]
					m[time[t][i]]=score[t][i]
				}else{
					count=count+score[t][i]
				}
			}else{
				m[time[t][i]]=score[t][i]
			}
		}
		fmt.Println(count)
	}
}
