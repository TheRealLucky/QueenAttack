package main

import (
	"errors"
	"os"
	"fmt"
	"strings"
	"strconv"
)

type Queen struct {
	x_pos int16
	y_pos int16
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

func getQueen(x_pos, y_pos int16, color string) (*Queen, error) {
	if (x_pos >= 0 && x_pos < 7) && (y_pos >= 0 && y_pos < 7) {
		return &Queen{x_pos, y_pos, color}, nil
	} else {
		return nil, errors.New("Position of the figure has to be between 0 and 7!")
	}
}

func interpret(args []string) (*Queen, *Queen, error) {
	q1_val := strings.Split(args[1], ",")
	q2_val := strings.Split(args[2], ",")

	q1_x, conv_err := strconv.Atoi(q1_val[0])
	q1_y, conv_err := strconv.Atoi(q1_val[1])
	q2_x, conv_err := strconv.Atoi(q2_val[0])
	q2_y, conv_err := strconv.Atoi(q2_val[1])

	if conv_err != nil {
		return nil, nil, conv_err
	}

	q1, err := getQueen(int16(q1_x), int16(q1_y), q1_val[2])
	q2, err := getQueen(int16(q2_x), int16(q2_y), q2_val[2])

	return q1, q2, err
}

func main() {
	q1, q2, _ := interpret(os.Args)

	fmt.Println(checkAttack(q1, q2))
}
