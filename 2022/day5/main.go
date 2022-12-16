package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {


	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	scanner := bufio.NewScanner(f)

	crates := make([][]string, 9)
	for i:=0; i<9; i++ {
		crates[i] = make([]string, 0)
	}

	t := 0
	for scanner.Scan() {
		j := 0
		for i := 1; i <= 33; i+=4 { 
			if string(scanner.Text()[i]) != " " {
				crates[j] = append(crates[j], string(scanner.Text()[i]))
			}
			j++
		}

		t++
		if t == 8 {
			break
		}

	}

	scanner.Scan()
	scanner.Scan()

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		move, _ := strconv.Atoi(line[1])
		from, _ := strconv.Atoi(line[3])
		to, _ := strconv.Atoi(line[5])

		crate := make([]string, 0)
		for move != 0 {
			crate = append(crate, crates[from-1][0])
			crates[from-1] = crates[from-1][1:]
			move--
		}

		crates[to-1] = append(crate, crates[to-1]...)
	}

	for i:=0; i<9; i++ {
		fmt.Println(crates[i][0])
	}
}