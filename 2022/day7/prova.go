package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	father *node
	childrens []*node
	name string
	size int
}

func newNode(value string, father *node) *node {
	
	childrens := make([]*node, 0)
	
	return &node{
		father,
		childrens,
		value,
		0,
	}
}

func addChildrens(childrens []string, father *node) {

	for _, c := range childrens {
		child := newNode(c, father)
		father.childrens = append(father.childrens, child)
	}

}

func findChild(child string, father *node) *node {
	for _, c := range father.childrens {
		if c.name == child {
			return c
		}
	}

	return nil
}

func main() {


	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	scanner := bufio.NewScanner(f)
	scanner.Scan()

	for {
		line := scanner.Text()

		commands := strings.Split(line, " ")
		if commands[0] == "$" {

			if commands[1] == "cd" {
				if commands[2] == ".." {
					fmt.Println("..")
				} else {
					fmt.Println(commands[2])
				}
			}

			scanner.Scan()

			if commands[1] == "ls" {
		
				for {

					line := scanner.Text()
					if line == "" {break}


					commands := strings.Split(line, " ")
					if commands[1] == "cd" {
						break
					} 

					fmt.Println(line)
					scanner.Scan()
					
				}
			}

		}

		if line == "" { break }

	}




}