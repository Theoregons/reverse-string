package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Mengubah string menjadi slice rune untuk menangani karakter multibyte
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "Hello, World!"
	reversed := reverseString(input)
	fmt.Println(reversed) // Output: !dlroW ,olleH
}
