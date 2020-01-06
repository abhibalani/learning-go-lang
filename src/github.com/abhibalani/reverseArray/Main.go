package main

import "fmt"

func main() {

	i := reverse([]string{"1", "2", "3", "4", "5"})
	fmt.Println(i)
}

func reverse(text []string) []string {
	last := len(text) - 1
	for i := 0; i < len(text)/2; i++ {
		text[i], text[last-i] = text[last-i], text[i]
	}
	return text
}
