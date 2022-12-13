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

		elf1 := scanner.Text()
		scanner.Scan()
		elf2 := scanner.Text()
		scanner.Scan()
		elf3 := scanner.Text()

		var common byte = '0'

		for i := 0; i < len(elf1); i++ {
			for j := 0; j < len(elf2); j++ {
				if elf1[i] == elf2[j] {
					for l := 0; l < len(elf3); l++ {
						if elf1[i] == elf3[l] {
							common = elf1[i]
							break
						}
					}
				}
				if common != '0' {
					break
				}
			}
			if common != '0' {
				break
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