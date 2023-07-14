package main

import (
	"fmt"
	"strings"
)
// reverseWords reverses the order of words in a given string.
//
// s: the input string to reverse the order of words in.
// return: the string with its words reversed.
func reverseWords(s string) string {
    words := strings.Fields(s)
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
}