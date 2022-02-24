package task1

import (
	"fmt"
	"strings"
)

func Run() {
	students := []string{"Jermaine", "Jessie", "Lamar"}
	scores := map[string][]float64{
		"Jermaine": {95.40, 82.30, 74.60},
		"Jessie":   {79.30, 99.10, 82.50},
		"Lamar":    {82.20, 95.40, 77.60},
	}

	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n", "Name", "Score 1", "Score2", "Score 3", "Average")
	fmt.Println(strings.Repeat("-", 55))

	for _, student := range students {
		scoreSummary(student, scores[student])
	}
}

func scoreSummary(student string, scores []float64) {
	sum := 0.00
	for _, score := range scores {
		sum += score
	}
	avg := sum / 3
	fmt.Printf("%10s | %8.2f | %8.2f | %8.2f | %8.2f\n", student, scores[0], scores[1], scores[2], avg)
}
