// Example 1:
// Input: words = ["pay","attention","practice","attend"], pref = "at"
// Output: 2
// Explanation: The 2 strings that contain "at" as a prefix are: "attention" and "attend".

// Example 2:
// Input: words = ["leetcode","win","loops","success"], pref = "code"
// Output: 0
// Explanation: There are no strings that contain "code" as a prefix.

package main

import "fmt"

func main() {

	prif := []string{"at"}

	count := 0

	var bool bool

	Input := []string{"pay", "attention", "practice", "attend"}

	for i := 0; i < len(Input); i++ {
		for j := 0; j < len(prif); j++ {

			if Input[i] == prif[j] {
				bool = true
			}
			if bool == true {
				count++
			}

		}

	}

	Output := count

	fmt.Println("output = ", Output)
}
