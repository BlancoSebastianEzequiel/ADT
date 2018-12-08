package main

import (
	"fmt"
	"packets"
)

func isAnagram(word1 string, word2 string) bool {
	var dict1 = packets.Dictionary{}.NewDicc()
	var dict2 = packets.Dictionary{}.NewDicc()
	for _, char := range word1 {
		dict1.Put(string(char))
	}
	for _, char := range word2 {
		dict2.Put(string(char))
	}
	return dict1.Equals(dict2)
}

func main() {
	isOsoAnagramOfRat := isAnagram("oso", "rat")
	fmt.Printf("is oso anagram of rat: %t\n", isOsoAnagramOfRat)
}