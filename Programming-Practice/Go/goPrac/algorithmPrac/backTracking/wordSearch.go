package main

import "fmt"

func exist(board [][]byte,word string)bool{
	visit:=make([][]bool,len(board))
	for i,_:= range board{
		visit[i]=make([]bool,len(board[0]))
	}
	move:=[][2]int{{0,1},{0,-1},{1,0},{-1,0}}
	var dfsword func(loacte [2]int,w string) bool
	dfsword=func(locate [2]int,w string)bool{
		if len(w)==0{
			return true
		}
		visit[locate[0]][locate[1]]=true
		var flag bool
		for _,m :=range move {
			x, y := locate[0]+m[0], locate[1]+m[1]
			if x >= 0 && x < len(board) && y >= 0 &&y < len(board[0]) {
				if board[x][y] == w[0] && !visit[x][y] {
					if dfsword([2]int{x,y},w[1:]) {
						flag=true
						break
					}
				}
			}
		}
		visit[locate[0]][locate[1]]=false
		return flag

	}
	for i,row:= range board{
		for j,val := range row{
			if val==word[0]{
				if dfsword([2]int{i,j},word[1:]){
					return true
				}
			}
		}
	}
	return false
}

func main(){
	//board:=[][]byte{{'C','A','A'},{'A','A','A'},{'B','C','D'}}
	//word:="AAB"
	board:=[][]byte{{'a'}}
	word:="ab"
	fmt.Println(exist(board,word))
}
