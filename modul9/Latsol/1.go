package main

import "fmt"

func main() {
	var motor, totalmotor int
	fmt.Scan(&motor)
	totalmotor = motor / 2
	if motor%2 != 0 {
		totalmotor += 1
	}
	fmt.Println(totalmotor)
}
