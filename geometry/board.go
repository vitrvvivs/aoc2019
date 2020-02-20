package geometry

import (
	"bufio"
	"fmt"
	"os"
)

// Board represents a Box with a 2d array of chars
type Board struct {
	Box
	Data [][]rune
}

// NewBoard creates a new Board
func NewBoard(filename string) (board *Board, err error) {
	file, err := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err != nil {
		return nil, err
	}

	board = &Board{Box{0, 0}, make([][]rune, 0)}
	for scanner.Scan() {
		board.Data = append(board.Data, []rune(scanner.Text()))
	}

	board.W = len(board.Data[0])
	board.H = len(board.Data)
	for _, line := range board.Data {
		if len(line) != board.W {
			return nil, fmt.Errorf("%s is not a rectangle", filename)
		}
	}
	board.Transpose()

	return board, err
}

// Print a Board
func (board Board) Print() {
	for y := 0; y < board.H; y++ {
		for x := 0; x < board.W; x++ {
			fmt.Print(string(board.Data[x][y]))
		}
		fmt.Println()
	}
}

// PrintMapping for subclasses
func PrintMapping(otherData [][]int, mapping map[int]rune) {
	for y := 0; y < len(otherData[0]); y++ {
		for x := 0; x < len(otherData); x++ {
			fmt.Print(string(mapping[otherData[x][y]]))
		}
		fmt.Println()
	}
}

// Transpose swaps rows and columns
func (board *Board) Transpose() {
	newData := make([][]rune, board.W)

	for x := 0; x < board.W; x++ {
		newData[x] = make([]rune, len(board.Data))
		for y := 0; y < len(board.Data); y++ {
			newData[x][y] = board.Data[y][x]
		}
	}
	board.Data = newData
}
