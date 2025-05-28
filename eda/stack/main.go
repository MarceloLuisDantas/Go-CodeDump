package main

import (
	"fmt"
	"stack/stack"
)

func main() {
	var s1 = stack.Stack{}
	var _, err = s1.Pop()
	if err != nil {
		fmt.Println(err)
	}

	_, err = s1.Get(2)
	if err != nil {
		fmt.Println(err)
	}

	len := s1.Len()
	if len != 0 {
		fmt.Println("sei la como isso aconteceu")
	}

	s1.Push(10)
	s1.Push(10)
	s1.Push(10)

	var x, _ = s1.Get(2)
	if x != 10 {
		fmt.Println("Ok isso já é estranho")
	}

	x, _ = s1.Pop()
	if x != 10 {
		fmt.Println("Eu to com sono")
	}

	if s1.Len() != 2 {
		fmt.Println("ainda to com sono")
	}
}
