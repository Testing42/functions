package main

import "fmt"

/*
a function can be a parameter for another function
*/

func main() {
	numbers := []int{1, 2, 3, 4}
	double := transformNumbers(&numbers, double)
	triple := transformNumbers(&numbers, triple)

	fmt.Println(double)
	fmt.Println(triple)
}

/*
TRANSFORM PARAMETER BELOW functions are first class functions
functions are just vaules in go and can be used
*/
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
