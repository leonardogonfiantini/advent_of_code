package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	matrix := make([][]int, 99)

	len := len(matrix)

	index := 0
	for scanner.Scan() {

		line := scanner.Text()

		for _, l := range line {
			li, _ := strconv.Atoi(string(l))
			matrix[index] = append(matrix[index], li)
		}

		index++
	}

	res := 0

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {

			flag := 0
			val := matrix[i][j]

			//check left
			if i == 0 {
				flag = 1
			} else {
				for p := 0; p < j; p++ {
					if matrix[i][p] < val {
						flag = 1
					} else {
						flag = 0
						break
					}
				}
			}

			//check right
			if i == len-1 || flag == 1 {
				flag = 1
			} else {
				for p := j + 1; p < len; p++ {
					if matrix[i][p] < val {
						flag = 1
					} else {
						flag = 0
						break
					}
				}
			}

			//check up
			if j == 0 || flag == 1 {
				flag = 1
			} else {
				for p := 0; p < i; p++ {
					if matrix[p][j] < val {
						flag = 1
					} else {
						flag = 0
						break
					}
				}
			}

			//check bottom
			if j == len-1 || flag == 1 {
				flag = 1
			} else {
				for p := i + 1; p < len; p++ {
					if matrix[p][j] < val {
						flag = 1
					} else {
						flag = 0
						break
					}
				}

			}

			if flag == 1 {
				res++
			}

		}
	}

	fmt.Println(res)

	res = 0
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {

			val := matrix[i][j]

			//check left
			left := 0
			for p := j-1; p >= 0; p-- {
				left++
				if matrix[i][p] >= val {
					break
				} 
			}
			
			//check right
			right := 0
			for p := j+1; p < len; p++ {
				right++
				if matrix[i][p] >= val {
					break
				} 
			}
			
			//check up
			up := 0
			for p := i-1; p >= 0; p-- {
				up++
				if matrix[p][j] >= val {
					break
				} 
			}
			

			//check down
			down := 0	
			for p := i+1; p < len; p++ {
				down++
				if matrix[p][j] >= val {
					break
				} 
			}

			pres := up * left * down * right

			if res < pres {
				res = pres
				
				fmt.Println(i,j,"=>",up, left, down, right, "=>", res)
			}

		}
	}

	fmt.Println(res)


}
