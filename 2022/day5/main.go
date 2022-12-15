package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	value string
	next *node
}

type stack struct {
	start *node
}

func (s *stack) push_init(val string) {

	new_node := &node {
		val, 
		nil,
	}

	if s.start == nil {
		s.start = new_node
	} else {
		tmp := s.start
		for tmp != nil {
			tmp = tmp.next
		}
		tmp = new_node
	}
}

func main() {


	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	scanner := bufio.NewScanner(f)

	scanner.Scan()

	for i := 1; i < 33; i+=4 { 
		fmt.Println(string(scanner.Text()[i]))
	}





}