package task2

import "fmt"

func Run() {
	var sum float64
	sum += perimeter(8.2, 10)
	sum += perimeter(5, 5.4)
	sum += perimeter(6.2, 4.5)

	fmt.Printf("You'll need %0.2f meters of fencing\n", sum)
}

func perimeter(l float64, w float64) float64 {
	return l*2 + w*2
}
