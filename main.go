package main

import (
	"errors"
	_ "os"
	_ "fmt"
)

type Queen struct {
	x_pos int
	y_pos int
	color string
}

var field [7][7]int

func checkAttack(q1, q2 Queen) bool {
	//TODO
	return false
}

func getQueen(x_pos, y_pos int, color string) (*Queen, error) {
	if (x_pos > 0 && x_pos < 8) && (y_pos > 0 && y_pos < 8) {
		return &Queen{x_pos, y_pos, color}, nil
	} else {
		return nil, errors.New("Position of the figure has to be between 0 and 7!")
	}
}

func main() {
}