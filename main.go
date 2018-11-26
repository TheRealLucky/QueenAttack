package main

import (
	"errors"
	_ "os"
	"fmt"
)

type Queen struct {
	x_pos int
	y_pos int
	color string
}

var field [7][7]int

func checkAttack(q1, q2 *Queen) bool {
	if(q1.x_pos == q2.x_pos || q1.y_pos == q2.y_pos) {
		return true
	} else {
		x_pos := q1.x_pos
		y_pos := q1.y_pos 
		for {
			if (x_pos >= 0 && x_pos <= q1.x_pos) && (y_pos >= 0 && y_pos <= q1.y_pos) {
				x_pos -= 1
				y_pos -= 1
				if(x_pos == q2.x_pos && y_pos == q2.y_pos) {
					return true
				}
			} else {
				x_pos := q1.x_pos
				y_pos := q1.y_pos
				for{ 
					if (x_pos < 7 && x_pos >= q1.x_pos) && (y_pos < 7 && y_pos >= q1.y_pos) {
						x_pos += 1
						y_pos += 1
						if(x_pos == q2.x_pos && y_pos == q2.y_pos) {
							return true
						}
					} else {
						return false
					}
				}
			}
		}
	}
	return false
}

func getQueen(x_pos, y_pos int, color string) (*Queen, error) {
	if (x_pos >= 0 && x_pos < 7) && (y_pos >= 0 && y_pos < 7) {
		return &Queen{x_pos, y_pos, color}, nil
	} else {
		return nil, errors.New("Position of the figure has to be between 0 and 7!")
	}
}

func main() {
	q1, err := getQueen(1, 2, "black")
	q2, err := getQueen(0, 1, "white")

	if err != nil{
		panic(err)
	}

	fmt.Println(checkAttack(q1, q2))
}