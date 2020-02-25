package main

import (
	"fmt"
)

func sum(values ...int) int {
	retSum := 0
	for _, value := range values {
		retSum += value
	}

	return retSum

}

func sumOfOddNumbers(splSum func(values ...int) int, receivedSlice ...int) int {

	var oddNumbersList []int
	for _, value := range receivedSlice {
		if value%2 != 0 {
			oddNumbersList = append(oddNumbersList, value)
		}
	}

	return (splSum(oddNumbersList...))

}

func main() {
	mainSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	oddNumberSum := sumOfOddNumbers(sum, mainSlice...)

	fmt.Println("sum of odd numbers is: ", oddNumberSum)
}
