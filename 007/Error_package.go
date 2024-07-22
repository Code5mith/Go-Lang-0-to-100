package main

import (
	"fmt",
	"errors"
)

// the errors module is a built in module that exposes some additional functionallity 
// with error handling.

func divide(x, y float64) (float64, error) {

	if y == 0 {
		return nil, errors.New("invalid denominator 0 is not allowed!")
	}
	return fmt.Sprintf("The quotent is : $%v", x/y), nil
}

func main() {
	divide(2.5,0);

}
