package main

import (
	"fmt"
)

func main() {
	fmt.Println("Task1:")
	Task1()
	fmt.Println("Second Task:")
	Task2()
}

func Task1() {
	var count int = 20
	unitWeight := 0.4
	totalWeight := float64(count) * unitWeight
	fmt.Println(count, "cans weigh", totalWeight, "kilograms")
}

func Task2() {
	var pebbleWeight float64 = 0.1
	var rockWeight float64 = 1.2
	var boulderWeight float64 = 502.4
	var totalWeight float64 = pebbleWeight + rockWeight + boulderWeight
	fmt.Println(totalWeight)
}
