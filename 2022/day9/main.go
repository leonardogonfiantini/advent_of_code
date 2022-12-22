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

	if  (h.y == t.y && h.x == t.x) ||
		(h.y == t.y && h.x+1 == t.x) || (h.y-1 == t.y && h.x == t.x) ||
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

	f, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	scanner := bufio.NewScanner(f)

	knots := make([]*point, 10)
	for i := 0; i < len(knots); i++ {
		knots[i] = &point{0,0,}
	}

	p_arr := make([]point, 0)

	for scanner.Scan() {

		line := scanner.Text()	
		commands := strings.Split(line, " ")
		steps, _ := strconv.Atoi(commands[1])

		for j := 0; j < steps; j++ {

			
			switch commands[0] {
			case "U":
				knots[0].y++
			case "D":
				knots[0].y--
			case "L":
				knots[0].x--
			case "R":
				knots[0].x++
			}
			
			

			for i := 1; i < len(knots); i++ {


				if !isTouching(knots[i-1], knots[i]) {
					if pointDistance(knots[i-1], knots[i]) {
						if knots[i-1].x == knots[i].x && knots[i-1].y == (knots[i].y+2) {
							knots[i].y++
						} 

						if knots[i-1].x == knots[i].x && knots[i-1].y == (knots[i].y-2) {
							knots[i].y--
						}

						if knots[i-1].x == (knots[i].x-2) && knots[i-1].y == knots[i].y {
							knots[i].x--
						}

						if knots[i-1].x == (knots[i].x+2) && knots[i-1].y == knots[i].y {
							knots[i].x++
						}
					} else {

						//1 1
						if (knots[i-1].x == knots[i].x+1 && knots[i-1].y-1 == knots[i].y+1) ||
							(knots[i-1].x-1 == knots[i].x+1 && knots[i-1].y == knots[i].y+1) {
								knots[i].x--
								knots[i].y--
						}

						
						//-1 1
						if (knots[i-1].x+1 == knots[i].x-1 && knots[i-1].y == knots[i].y+1) ||
							(knots[i-1].x == knots[i].x-1 && knots[i-1].y-1 == knots[i].y+1) {
								knots[i].x++
								knots[i].y--
						}


						//-1 -1
						if (knots[i-1].x == knots[i].x-1 && knots[i-1].y+1 == knots[i].y-1) ||
							(knots[i-1].x+1 == knots[i].x-1 && knots[i-1].y == knots[i].y-1) {
								knots[i].x++
								knots[i].y++
						}


						//1 -1
						if (knots[i-1].x == knots[i].x+1 && knots[i-1].y+1 == knots[i].y-1) ||
							(knots[i-1].x-1 == knots[i].x+1 && knots[i-1].y == knots[i].y-1) {
								knots[i].x--
								knots[i].y++
						}
					
					}

				}



			}

			p_arr = append(p_arr, *knots[len(knots)-1])

		}

	}

	final_arr := make([]point, 0)
	for _, p := range p_arr {

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