package main

import (
	_ "os"
	_ "fmt"
)

type Queen struct {
	x_pos int
	y_pos int
	color string
}

var field [8][8]int

func checkAttack(q1, q2 Queen) bool {
	//TODO
	return false
}

func main() {
}