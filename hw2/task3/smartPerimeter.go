package task3

import "fmt"

func Run() {
	perimeter, err := smartPerimeter(8.2, -10)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Println(perimeter)
}

func smartPerimeter(l float64, w float64) (float64, error) {
	if l < 0 {
		return 0, fmt.Errorf("a %s of %0.2f is invalid", "length", l)
	}

	if w < 0 {
		return 0, fmt.Errorf("a %s of %0.2f is invalid", "width", w)
	}

	return l*2 + w*2, nil
}
