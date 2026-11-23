package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[1] + fibs[2]
	fibs[4] = fibs[2] + fibs[3]

	fibs = append(fibs, 8)

	fibs = append(fibs, 13)
	fibs = append(fibs, 21)
	fibs = append(fibs, 34)

	fmt.Println(fibs)
}
