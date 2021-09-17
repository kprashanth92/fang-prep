package main

import "fmt"

func main() {
	fmt.Println("pairs.go")
	input := []int{10, 5, 2, 3, -6, 9, 11}
	sum := 4

	output := pairOutput(input, sum)
	fmt.Println(output)
}

func pairOutput(input []int, sum int) []int {
	var output []int
	m := make(map[int]struct{})
	for _, item := range input {
		x := sum - item
		if _, ok := m[x]; ok {
			output = append(output, x)
			output = append(output, item)
			return output
		}

		m[item] = struct{}{}
	}

	return output
}
