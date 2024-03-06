package main

import (
	"fmt"

	"github.com/Exitialis/package-test/calculator"
	"github.com/Exitialis/package-test/names"

	"github.com/pkg/errors"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	name := names.New("test")
	err := errors.New("test")
	fmt.Println(err)
	fmt.Println(calculator.Add(10, 20))
	fmt.Println("Hello " + name.ToString())
	fmt.Println("update")
}
