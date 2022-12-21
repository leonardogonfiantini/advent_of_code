package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

type point struct {
	x float64
	y float64
}


func isTouching(h *point, t *point) bool {

	if  (h.y == t.y && h.x+1 == t.x) || (h.y-1 == t.y && h.x == t.x) ||
		(h.y == t.y && h.x-1 == t.x) || (h.y+1 == t.y && h.x == t.x) ||
		(h.y+1 == t.y && h.x+1 == t.x) || (h.y-1 == t.y && h.x-1 == t.x) ||
		(h.y+1 == t.y && h.x-1 == t.x) || (h.y-1 == t.y && h.x+1 == t.x) {
						return true
					}

	return false
}

//i just need to know if need a diagonal move or a normal one so is bool
func pointDistance(p1 *point, p2 *point) bool {

	if math.Sqrt(math.Pow((p1.x-p2.x), 2)+(math.Pow((p1.y-p2.y), 2))) <= 2 {
		return true
	}
	
	return false
}


func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	scanner := bufio.NewScanner(f)

	h := &point {0,0,}
	t := &point {0,0,}

	p_arr := make([]point, 0)

	for scanner.Scan() {

		line := scanner.Text()	
		commands := strings.Split(line, " ")
		steps, _ := strconv.Atoi(commands[1])

		for i := 0; i < steps; i++ {

			switch commands[0] {
			case "U":
				h.y++
			case "D":
				h.y--
			case "L":
				h.x--
			case "R":
				h.x++
			}

			if h.x == t.x && h.y == t.y { continue }

			if !isTouching(h, t) {
				if pointDistance(h, t) {
					switch commands[0] {
					case "U":
						t.y++
					case "D":
						t.y--
					case "L":
						t.x--
					case "R":
						t.x++
					}
				} else {
					switch commands[0] {
					case "U":
						t.y = h.y-1
						t.x = h.x
					case "D":
						t.y = h.y+1
						t.x = h.x
					case "L":
						t.y = h.y
						t.x = h.x+1
					case "R":
						t.y = h.y
						t.x = h.x-1
					}
				}

			}

			p_arr = append(p_arr, *t)
		}

	}

	final_arr := make([]point, 0)
	for _, p := range p_arr {

		fmt.Println(p)
		flag := 0
		for _, f := range final_arr {
			if (f.x == p.x && f.y == p.y) {
				flag = 1
				break
			}
		}

		if flag == 0 {
			final_arr = append(final_arr, p)
		}

	}

	fmt.Println(len(p_arr))
	fmt.Println(len(final_arr))

}