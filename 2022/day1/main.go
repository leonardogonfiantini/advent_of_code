package main

import (
	"fmt"
    "bufio"
	"os"
	"log"
	"strconv"
)

func main() {

	f, err := os.Open("./input.txt")
    if err != nil {
		panic(err)
	}	
	defer f.Close()

	scanner := bufio.NewScanner(f)

	max_elf_index := 0
	max_elf_calories := 0
	index := 0
	calories := 0

	for scanner.Scan() {

		if scanner.Text() == "" {

			if calories > max_elf_calories {
				max_elf_calories = calories
				max_elf_index = index
			}

			fmt.Printf("%s => %s \n", strconv.Itoa(index), strconv.Itoa(calories))


			index++
			calories = 0

			scanner.Scan()
		}

		tmp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		calories += tmp

	}

	fmt.Printf("\nThe elf with max calories is the elf => %s with %s calories", strconv.Itoa(max_elf_index), strconv.Itoa(max_elf_calories))

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}	