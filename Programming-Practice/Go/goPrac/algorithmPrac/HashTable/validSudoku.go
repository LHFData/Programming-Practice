package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	row:=[10][9]int{}
	line:=[10][9]int{}
	square:=[10][9]int{}
	for i:=0;i<9;i++{
		//第i行
		for j:=0;j<9;j++{
			//第j列
			 current:=int(board[i][j])

			 if current==46{
			 	continue
			 }else{
			 	current=current-48
			 }
			if row[current][i]==1{
				return false
			}else{
				row[current][i]=1
			}
			if line[current][j]==1{
				return false
			}else{
				line[current][j]=1
			}
			if square[current][(i/3)*3+j/3]==1{
				return false
			}else{
				square[current][(i/3)*3+j/3]=1
			}

		}
	}
	return true
}

func main()  {
	input:= [][]string{
		{".","1",".","5","2",".",".",".","."},
		{".",".",".",".",".","6","4","3","."},
		{".",".",".",".",".",".",".",".","."},
		{"5",".",".",".",".",".","9",".","."},
		{".",".",".",".",".",".",".","5","."},
		{".",".",".","5",".",".",".",".","."},
		{"9",".",".",".",".","3",".",".","."},
		{".",".","6",".",".",".",".",".","."},
		{".",".",".",".",".",".",".",".","."}}
	byin:=[][]byte{}
	for i:=0;i<9;i++{
		temp:=make([]byte,9)
		for j:=0;j<9;j++{
			temp[j]=byte(rune(input[i][j][0]))
		}
		byin=append(byin,temp)
	}
	fmt.Println(isValidSudoku(byin))
}