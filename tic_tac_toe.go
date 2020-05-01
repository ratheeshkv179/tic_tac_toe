package main

import (
	"bufio"
	"fmt"
	"os"
)

var board []string
var winningPattern [][]int
var counter int

func init() {
	counter = 0
	board = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	hori1 := []int{0, 1, 2}
	hori2 := []int{3, 4, 5}
	hori3 := []int{6, 7, 8}

	ver1 := []int{0, 3, 6}
	ver2 := []int{1, 4, 7}
	ver3 := []int{2, 5, 8}

	diag1 := []int{0, 4, 8}
	diag2 := []int{2, 4, 6}

	winningPattern = make([][]int, 0)
	winningPattern = append(winningPattern, hori1)
	winningPattern = append(winningPattern, hori2)
	winningPattern = append(winningPattern, hori3)
	winningPattern = append(winningPattern, ver1)
	winningPattern = append(winningPattern, ver2)
	winningPattern = append(winningPattern, ver3)
	winningPattern = append(winningPattern, diag1)
	winningPattern = append(winningPattern, diag2)

}

func updateBoard(index int, marker string) {
	board[index-49] = marker
}

func gameOver() bool {

	for _, pattern := range winningPattern {
		if board[pattern[0]] == board[pattern[1]] && board[pattern[1]] == board[pattern[2]] {
			if board[pattern[0]] == "X" {
				fmt.Println("Winner : Player-2")
			} else if board[pattern[0]] == "0" {
				fmt.Println("Winner : Player-1")
			}
			displayBoard()
			return true
		}
	}

	if counter == 9 {
		fmt.Println("Match Draw!")
		return true
	}
	return false
}

func displayBoard() {

	for i, val := range board {

		if (i+1)%3 == 0 {
			fmt.Printf(" %v", val)
			fmt.Println()
			if i+1 != 9 {
				fmt.Println("-----------")
			}
		} else {
			fmt.Printf(" %v |", val)
		}

	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	players := []string{"0", "X"}
	var player int = 0

	for true {
		displayBoard()
		fmt.Printf("\n[Player1: 0, Player2: X]\n\nPlayer%v : ", player+1)
		if scanner.Scan() {
			inp := scanner.Text()
			if len(inp) > 0 {
				if inp[0] >= 49 && inp[0] <= 57 {
					counter++
					updateBoard(int(inp[0]), players[player])
					if gameOver() {
						break
					}
				}
			}
		}
		if player == 0 {
			player = 1
		} else {
			player = 0
		}
		fmt.Println()
	}
}
