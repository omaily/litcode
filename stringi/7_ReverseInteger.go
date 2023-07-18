package stringi

import (
	"errors"
	"fmt"
	"math"
)

func task_7() {
	fmt.Printf("%d to ", 146746549)
	println(reverse(146746549))
}

var ErrOverflow = errors.New("integer_32 overflow")

func add32(a, b int) (int, error) {
	if a > math.MaxInt32-b {
		return math.MaxInt32, ErrOverflow
	}
	return a + b, nil
}

func reverse(x int) int {

	var y, signed int

	if x > 0 {
		signed = 1
	} else {
		x *= -1
		signed = -1
	}

	for i := x; i > 0; i /= 10 {
		var err error
		y, err = add32(y*10, i%10)
		if err != nil {
			return 0
		}
	}
	return y * signed
}
