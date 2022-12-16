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
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	line := scanner.Text()

	for i := 0; i < len(line)-14; i++ {

		f := 0
		for j := i; j < i + 14; j++ {
			for l := 1; l < 14; l++ {
				if j != i+ l {
					if line[j] == line[i+l] {
						f = 1
						break
					}
				}
			}
			if f == 1 {
				break
			}
		}

		if f == 0 {
			fmt.Println(i+14)
			break
		}

	}


}