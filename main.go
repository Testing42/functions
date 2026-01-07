package main

import "fmt"

/*
a function can be a parameter for another function

for better organization its best to make a custom type and then use that instead of the fucntion
directly in the other function. this makes things clearer and less messy
*/

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	double := transformNumbers(&numbers, double)
	triple := transformNumbers(&numbers, triple)

	fmt.Println(double)
	fmt.Println(triple)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformNumbers1 := transformNumbers(&numbers, transformFn1)

	moreTransformedNumbers := transformNumbers(&numbers, transformFn2)

	fmt.Println(transformNumbers1)
	fmt.Println(moreTransformedNumbers)
}

/*
TRANSFORM PARAMETER BELOW functions are first class functions
functions are just vaules in go and can be used

to make a variadiac value which allows for calling a function with any number of parameters like

func sumup(numbers ...int) int {

That only works if you pass a list of parameters which
isn't setup to be an array of something
like

sum := sumup(1, 10,22, 43, -95)

you can make an array into a slice using
anothersum := sumup(numbers...)

*/
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
