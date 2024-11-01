package main

import "fmt"

func main() {
	word1 := "madam"
	word2 := "mom"
	word3 := "12211"
	result := ""


	isPalindrome(word1, result)
	isPalindrome(word2, result)
	isPalindrome(word3, result)

}

func isPalindrome(w, r string) {
	for i := len(w) - 1; i >= 0; i-- {
		r += string(w[i])
	}
	if w == r {
		fmt.Printf("%s is Palindrome!\n", w)
	} else {
		fmt.Printf("%s is not a Palindrome!\n", w)
	}
	
}
