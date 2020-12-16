package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	m := make(map[byte]bool)
	for i := 0;i < len(board);i++{
		for j:=0;j < len(board);j++{
			if board[i][j] == '.'{
				continue
			}
			if _,ok := m[board[i][j]];ok{
				return false
			}else{
				m[board[i][j]] = true
			}
		}
		m = make(map[byte]bool)
	}
	for i := 0;i < len(board);i++{
		for j:=0;j < len(board);j++{
			if board[j][i] == '.'{
				continue
			}
			if _,ok := m[board[j][i]];ok{
				return false
			}else{
				m[board[j][i]] = true
			}
		}
		m = make(map[byte]bool)
	}
	res := true
	for i := 0;i <= 6;i += 3{
		for j := 0;j <= 6;j += 3{
			res = res && regionIsValid(i,i+2,j,j+2,&board)
			if !res{
				return res
			}
		}
	}
	return res
}

func regionIsValid(left,right,top,down int,board *[][]byte) bool{
	m := make(map[byte]bool)
	for i := left;i <= right;i++{
		for j := top;j <= down;j++{
			if (*board)[i][j] == '.'{
				continue
			}
			if _,ok := m[(*board)[i][j]];ok{
				return false
			}else{
				m[(*board)[i][j]] = true
			}
		}
	}
	return true
}

func main(){
	a := [][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(a))
}