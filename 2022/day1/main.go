package main

import (
	"fmt"
    "bufio"
	"os"
	"log"
	"strconv"
)

func takeMin(arr []int) int {

	min := arr[0]

	tmp_index := 0
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			tmp_index = i
		}
	}

	return tmp_index

}

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

	rank := make([]int, 3)

	for scanner.Scan() {

		if scanner.Text() == "" {

			if calories > max_elf_calories {
				max_elf_calories = calories
				max_elf_index = index
			}


			fmt.Printf("%d => %d \n", index, calories)

			min := takeMin(rank)
			if calories > rank[min] {
				rank[min] = calories
			}

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

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	fmt.Printf("\nThe elf with max calories is the elf => %d with %d calories", max_elf_index, max_elf_calories)

	sum_rank := 0;
	for _, i := range rank {
		sum_rank += i
	}
	fmt.Printf("\nThe sum of calories of the three elf with most calories is: %d", sum_rank)

}	