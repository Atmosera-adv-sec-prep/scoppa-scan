package main

import (
	"fmt"
	"os"
	"github.com/andrew-scoppa/go-modules/wordgame"
)

// this is the entry point
func main() {
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scanf("%s\n", &word)

	if is_pallindrome := wordgame.IsPalindrome(word); !is_pallindrome {
		fmt.Printf("%s is not a palindrome\n", word)
	}
	fmt.Printf("%s is a palindrome\n", word)
	os.Exit(0)
}
