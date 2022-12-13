package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("./input.txt")
	if (err != nil) {
		panic(err)
	}
	defer f.Close()


	scanner := bufio.NewScanner(f)
	score := 0
	sscore := 0

	for scanner.Scan() {

		moves := strings.Split(scanner.Text(), " ")
		player1 := moves[0]
		player2 := moves[1]

		switch player1 {
		case "A":
			switch player2 {
			case "X":
				score += 3 + 1
				sscore += 0 + 3
			case "Y":
				score += 6 + 2
				sscore += 3 + 1
			case "Z": 
				score += 0 + 3
				sscore += 6 + 2
			}
		case "B":
			switch player2 {
			case "X":
				score += 0 + 1
				sscore += 0 + 1
			case "Y":
				score += 3 + 2
				sscore += 3 + 2
			case "Z":
				score += 6 + 3
				sscore += 6 + 3
			}
		case "C":
			switch player2 {
			case "X":
				score += 6 + 1
				sscore += 0 + 2
			case "Y":
				score += 0 + 2
				sscore += 3 + 3
			case "Z":
				score += 3 + 3
				sscore += 6 + 1
			}
		}


		if scanner.Text() == "" {
			fmt.Println("ciao")
		}


	}

	fmt.Printf("The total score if we follow the strategy given is: %d\n", score)
	fmt.Printf("The total score by following the correct strategy given is: %d\n", sscore)


}