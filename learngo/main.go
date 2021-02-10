package main

import (
	"fmt"
	"strings"
)

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	totalLength, upperName := lenAndUpper("hyun")
	fmt.Println(totalLength, upperName)

	repeatMe("hyun", "asdf", "qwerty", "zxcv")

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
