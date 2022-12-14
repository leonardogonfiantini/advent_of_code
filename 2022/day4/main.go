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

	score := 0
	for scanner.Scan() {
		
		line := strings.Split(scanner.Text(), ",")
		
		assignments1 := strings.Split(line[0], "-")
		assignments2 := strings.Split(line[1], "-")

		a1, _ := strconv.Atoi(assignments1[0])
		a2, _ := strconv.Atoi(assignments1[1])
		a3, _ := strconv.Atoi(assignments2[0])
		a4, _ := strconv.Atoi(assignments2[1])


		if ((a1 <= a3) && (a2 >= a4)) || ((a3 <= a1) && (a4 >= a2)) {
				score++
		}

	}

	fmt.Printf("The total number of same pair is: %d\n", score)



}