package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var board = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
var turn = 0
var player string

func main() {
	//var box int
	var win bool
	//var placed bool
	go runServer()

	//i := 0
	//for ; i < 9; i++ {
	for turn < 9 && !win {
		time.Sleep(1 * time.Second)
		fmt.Println("Enter the box you want to play")
		//fmt.Scanln(&box)
		//placed = placeCounter(box)
		//for !placed {
		//	placed = placeCounter(box)
		//}
		//fmt.Println(board)
		win = checkWin()
		if win {
			fmt.Println(player, " has won!")
		}
	}
	if turn == 9 && !win {
		fmt.Println("draw")
	}
}

func placeCounter(box int) {
	//var placed bool
	fmt.Println(board)

	player = playPiece(box)
	drawBoard()
	//if placed {
	//	return true
	//} else {
	//	return false
	//}

}

func playPiece(box int) string {
	var player = ""
	if board[int(box/3)][box%3] != " " {
		return player
	}
	if turn%2 == 0 {
		player = "x"
		board[int(box/3)][box%3] = player
	} else {
		player = "o"
		board[int(box/3)][box%3] = player
	}
	turn++
	fmt.Println(board)
	return player
}

func checkWin() bool {
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

func drawBoard() {
	for x := 0; x < 3; x++ {
		var currentPrint string
		for i := 0; i < 3; i++ {
			currentPrint += board[x][i]
		}
		fmt.Println(currentPrint)
	}
}

func runServer() {
	http.HandleFunc("/state", getState)
	http.HandleFunc("/play", playMove)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(board)
}

//http://localhost/play?place=2
// http://blah/add?param1=somevalue&param2=something

func playMove(w http.ResponseWriter, r *http.Request) {

	targetPlace := r.URL.Query()["place"]
	fmt.Println(targetPlace)
	box, err := strconv.Atoi(targetPlace[0])
	if err != nil {
		fmt.Println(err)
	}
	placeCounter(box)
}
