// Vladislav Ebert
// 241RDB316
package main

import (
	"fmt"
	"strings"
)

func main() {
	var First_word, Second_word string

	fmt.Print("Enter first word: ")
	fmt.Scan(&First_word)
	fmt.Print("Enter second word: ")
	fmt.Scan(&Second_word)

	if Anagramma(First_word, Second_word) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func Anagramma(First_word, Second_word string) bool {

	First_word = strings.ToLower(First_word)
	Second_word = strings.ToLower(Second_word)

	if len(First_word) != len(Second_word) {
		return false
	}
	frekvence1 := make(map[rune]int)
	frekvence2 := make(map[rune]int)

	for _, burts := range First_word {
		frekvence1[burts]++
	}
	for _, burts := range Second_word {
		frekvence2[burts]++
	}
	if len(frekvence1) != len(frekvence2) {
		return false
	}
	for burts, dauzdzums1 := range frekvence1 {
		if dauzdzums2, exists := frekvence2[burts]; !exists || dauzdzums2 != dauzdzums1 {
			return false
		}
	}
	return true
}
