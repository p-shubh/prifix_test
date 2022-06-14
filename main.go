package main

import (
	"fmt"
	"strings"
)

func main() {

	// var str1 []string

	var count int

	Input := []string{"pay", "attention", "practice", "attend"}

	find := strings.Split("at", "")

	// fmt.Println(find)

	for i := 0; i < len(Input); i++ {

		str1 := strings.Split(Input[i], "")
		// fmt.Println(str1)
		count = 0

		for j := 0; j < len(find); j++ {
			if find[j] == str1[j] {
				count = count + 1
			}
		}

	}

	fmt.Println("the count is ", count)

	// 	// for i := 0; i < len(Input); i++ {
	// 	// 	for j := 0; j < len("at"); j++ {

	// 	// 	}
	// 	fmt.Println(strings.Contains(str1, "at"), str1)

	// 	if strings.Contains(str1, "at") == true {
	// 		count++
	// 	}

	// 	fmt.Println(count)

}

//
