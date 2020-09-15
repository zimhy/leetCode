package main

import "fmt"

func longestPalindrome(s string) string {
	allRunes := []rune(s)
	var result = make([]rune, 0)
	for i := 0; i < len(allRunes); i++ {
		runes := findByCenterIndex(allRunes, i)
		if len(runes) > len(result) {
			result = runes
		}
	}
	return string(result)
}

func findByCenterIndex(runes []rune, index int) (palindrome []rune) {
	if index > len(runes) {
		return make([]rune, 0)
	} else if index < len(runes) {
		palindrome = runes[index : index+1]
		for i := 0; i <= index && i+index < len(runes); i++ {
			if runes[index-i] == runes[index+i] {
				palindrometmp := runes[index-i : index+1+i]
				if len(palindrometmp) > len(palindrome) {
					palindrome = palindrometmp
				}
			} else {
				break
			}

		}
	}
	if index < len(runes)-1 {
		if runes[index] == runes[index+1] {
			for i := 0; i <= index && i+index+1 < len(runes); i++ {
				if runes[index-i] == runes[index+1+i] {
					palindrometmp := runes[index-i : index+2+i]
					if len(palindrometmp) > len(palindrome) {
						palindrome = palindrometmp
					}
				} else {
					break
				}

			}
		}
	}
	return palindrome

}
func main() {
	fmt.Println(longestPalindrome("asdasddd"))
}
