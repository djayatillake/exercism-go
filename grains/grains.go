/*package grains calculates the number of grains of wheat on a chessboard given
that the number on each square doubles; in total and on a given square*/
package grains

import "errors"

// Square returns the number of grains on a square
func Square(square int) (uint64, error) {
	if square <= 0 || square > 64 {
		return 0, errors.New("true")
	}
	return uint64(('1' - '0') << (square - 1)), nil
}

// Total returns the sum of all the grain on the chessboard
func Total() (ret uint64) {
	return (uint64(('1'-'0')<<63)-1)*2 + 1
}
