package main

import (
	"errors"
	"fmt"
)

type Func func(param int) error

func Wrapper(aFunc Func) Func {
	return func(param int) error {
		fmt.Println("Before calling the function", param)
		err := aFunc(param)
		fmt.Println("After calling the function", err)

		return errors.New("must wrong")
	}
}

func main() {
	err := Wrapper(func(param int) error {
		fmt.Println("Inside the function", param)
		return nil
	})(1)
	fmt.Println(err)
}
