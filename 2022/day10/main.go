package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func printCycle(cycle int, x int) int {
	for i := 20; i <= 220; i+=40 {
		if cycle == i {
			fmt.Println(i, x, ":", x*i)
			return x*i
		}
	}

	return 0
}

func printImage(cycle int, x int, matrix []string) {
	
	col := (cycle/40)


	fmt.Println(40*col)

	row := 40 * col

	if (row == 0) {
		row = 40
	}

	for i:=col*40; i < row; i++ {
		
		if x == col || x == col+1 || x == col+2 {
			matrix[i] = "#"
		} else {
			matrix[i] = "."
		}

	}

}


func main() {

	matrix := make([]string, 240)


	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cycle := 0
	x := 1
	res := 0
	for scanner.Scan() {

		line := scanner.Text()
		commands := strings.Split(line, " ")

		
		switch commands[0] {
		case "noop":
			cycle++
			res += printCycle(cycle, x)
			printImage(cycle, x, matrix)
		case "addx":
			for i := 1; i <= 2; i++ {
				cycle++
				res += printCycle(cycle, x)
				printImage(cycle, x, matrix)
			}

			val, err := strconv.Atoi(commands[1])
			if err != nil {
				panic(err)
			}

			x += val
		}

	}


	for i:=0; i < 240; i++ {

		if i != 0 {
			if (i%40 == 0) {
				fmt.Println(" ")
				
			}
		}

		fmt.Print(matrix[i])

	}


}