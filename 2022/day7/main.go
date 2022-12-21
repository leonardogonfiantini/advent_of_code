package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	father *node
	childrens []*node
	name string
	size int
}

func newNode(name string, father *node) *node {
	
	childrens := make([]*node, 0)
	s := 0


	return &node{
		father,
		childrens,
		name,
		s,
	}
}

func (father *node) addChildren(name string) {

	child := newNode(name, father)
	father.childrens = append(father.childrens, child)

}

func (father *node) findChild(child string) *node {
	for _, c := range father.childrens {
		if c.name == child {
			return c
		}
	}

	return nil
}

func (father *node) printFS(spaces int) {

	for _, c := range father.childrens {

		for i := 0; i < spaces; i++ {
			fmt.Print(" ")
		}

		fmt.Println(c.name, c.size)
		c.printFS(spaces+2)
	}


}

func (father *node) addSize(s int) {
	father.size += s
}

func (father *node) adjustsize() {


	for _, c := range father.childrens {

		c.adjustsize()
	}
	
	if father.childrens != nil {
		fsize := 0

		for _, c := range father.childrens {
			fsize += c.size
		}

		father.addSize(fsize)
	}
	
}

func (father *node) quest1() int {

	res := 0

	for _, c := range father.childrens {

		if (c.size < 100000) {
			res += c.size
		}

		res += c.quest1()
	}

	return res

}

func (father *node) quest2(ss int) {

	for _, c := range father.childrens {

		if (c.size >= ss) {
			fmt.Println(c.size)
		}

		c.quest2(ss)

	}

}

func main() {


	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	scanner := bufio.NewScanner(f)
	scanner.Scan()

	root :=  newNode("/", nil)
	father := root

	scanner.Scan()
	for {

		line := scanner.Text()

		commands := strings.Split(line, " ")
		if commands[0] == "$" {

			if commands[1] == "cd" {
				if commands[2] == ".." {
					father = father.father
				} else {
					father = father.findChild(commands[2])
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


					if commands[0] == "dir" {
						father.addChildren(commands[1])
					} else {
						size, _ := strconv.Atoi(commands[0])
						father.addSize(size)
					}


					scanner.Scan()
					
				}
			}
		}

		if line == "" { break }

	}



	fmt.Println(root.name)

	root.adjustsize()
	root.printFS(2)

	fmt.Println(root.quest1())
	
	fmt.Println("root size:", root.size)
	val := 70000000 - root.size
	val = 30000000 - val
	fmt.Println("val needed:", val)
	root.quest2(val)

}