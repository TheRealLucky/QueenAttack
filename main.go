package main

import (
	"errors"
	"os"
	"fmt"
	"strings"
	"strconv"
)

// Purpose of this program:
//
// Two queens(w = white, b=black) are placed on this board.
// The two queens can attack each other if they are in the
// same column, row, or diagonal to each other.

/* For example with parameters: 1,2,b 3,4,w
0 1 2 3 4 5 6 7 x|y

x x x x x x x x   0
x x x x x x x x   1
x b x x x x x x   2
x x x x x x x x   3
x x x w x x x x   4
x x x x x x x x   5
x x x x x x x x   6
x x x x x x x x   7

In this example b and w can attack each other because they are
placed diagonal.
*/

// Represents the Queen 
type Queen struct {
	x_pos int16
	y_pos int16
	color string
}

// Represents the 8x8 chess board
var field [7][7]int

// Takes two queens as parameters and checks if the queens are able
// to attack each other. This is only the case if the are in the same
// row, column or diagonal to each other.
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

// Takes the x and y position, as well as the color of the queen as
// parameters and returns a Queen
func getQueen(x_pos, y_pos int16, color string) (*Queen, error) {
	if (x_pos >= 0 && x_pos < 7) && (y_pos >= 0 && y_pos < 7) {
		return &Queen{x_pos, y_pos, color}, nil
	} else {
		return nil, errors.New("Position of the figure has to be between 0 and 7!")
	}
}

// Interprets the commandline arguments. 
// The help function dispalays the format, how the arguments
// should look like.
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

// Displays a help message how the commandline arguments
// must be formatted.
func help() {
	fmt.Println("Usage:")
	fmt.Println("queenAttack <x_pos>,<y_pos>,<color> <x_pos>,<y_pos>,<color>")
	fmt.Println()
	fmt.Println("Where the first 3 arguments stand for the position and color of")
	fmt.Println("the first queen and the second 3 for the second queen.")
}

func main() {
	q1, q2, _ := interpret(os.Args)

	fmt.Println(checkAttack(q1, q2))
}
