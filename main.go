package main

import (
	"fmt"
	"ipneschadin/learning-artem/internal/calculator"
	"ipneschadin/learning-artem/internal/names"

	"github.com/pkg/errors"
)

func main() {
	name := names.New("test")
	err := errors.New("test")
	fmt.Println(err)
	fmt.Println(calculator.Add(10, 20))
	fmt.Println("Hello " + name.ToString())
}
