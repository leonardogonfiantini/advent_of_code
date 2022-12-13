package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close();


	scanner := bufio.NewScanner(f)
	
	score := 0
	for scanner.Scan() {
		line := scanner.Text()

		var common byte = 'a'

		for i := 0; i < len(line)/2; i++ {
			c := line[i]
			for j := len(line)/2; j < len(line); j++ {
				if c == line[j] {
					common = c
				}
			}
		}


		if (int(common) < 97) {
			score += int(common) - 65 + 27
			fmt.Println(string(common), int(common) - 65 + 27)
		} else {
			score += int(common) - 96
			fmt.Println(string(common), int(common) - 96)
	
		}
	}

	fmt.Printf("The sum of all priorities is: %d", score)


}