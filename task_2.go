package main

import (
	"fmt"
	"strings"
)

func isAlphabet(ch string) bool {
	return (ch >= "a" && ch <= "z") || (ch >= "A" && ch <= "Z")
}

func wordFrequencyCount(word string) map[string]int {
	counter := map[string]int{}
	lowerCaseString := strings.ToLower(word)

	// extract rune from word and map to counter
	for _, rune := range lowerCaseString {
		ch := string(rune)
		if isAlphabet(ch) {
			counter[string(rune)] += 1
		}
	}

	fmt.Println(counter)
	return counter
}

func palindromeCheck(word string) bool {
	wordLength := len(word)
	l, r := 0, wordLength-1
	word = strings.ToLower(word)

	for l <= r {
		leftCh := string(word[l])
		rightCh := string(word[r])

		// check right ch
		if !isAlphabet(rightCh) {
			r -= 1
			continue
		}

		// check left ch
		if !isAlphabet(leftCh) {
			l += 1
			continue
		}

		if leftCh != rightCh {
			return false
		}

		l += 1
		r -= 1
	}

	return true
}

func main() {
	wordFrequencyCount("Hello3#word")

	fmt.Println(palindromeCheck("abab"))
}
