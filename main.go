package main

import "fmt"

var board = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}

func main() {
	var box int
	var end bool
	i := 0
	for ; i < 9; i++ {
		fmt.Println("Enter the box you want to play")
		fmt.Scanln(&box)
		end = placeCounter(&board, box, i)
		if end {
			break
		}
	}
	if i == 9 && !end {
		fmt.Println("draw")
	}
}

func placeCounter(board *[3][3]string, box int, turn int) bool {
	var player string
	if turn%2 == 0 {
		player = "x"
		board[int(box/3)][box%3] = player
	} else {
		player = "o"
		board[int(box/3)][box%3] = player
	}
	drawBoard(board)
	win := checkWin(board)
	if win {
		fmt.Println(player + " has won")
		return true
	} else {
		return false
	}
}

func checkWin(board *[3][3]string) bool {
	for x := 0; x < 3; x++ {
		if board[x][0] != " " && board[x][0] == board[x][1] && board[x][1] == board[x][2] {
			return true
		}
	}
	for y := 0; y < 3; y++ {
		if board[0][y] != " " && board[0][y] == board[1][y] && board[1][y] == board[2][y] {
			return true
		}
	}
	if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	if board[0][2] != " " && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}
	return false
}

func drawBoard(board *[3][3]string) {
	for x := 0; x < 3; x++ {
		var currentPrint string
		for i := 0; i < 3; i++ {
			currentPrint += board[x][i]
		}
		fmt.Println(currentPrint)
	}
}
