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

func canIDrink(age int) bool {

	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {

	repeatMe("hyun", "asdf", "qwerty", "zxcv")

	fmt.Println(canIDrink(14))
}
