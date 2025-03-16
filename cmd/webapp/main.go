package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "hello my name is sankalpa. sankalpa is a programmer. programmers are genreally a person who writes code"

	results := strings.Split(data, " ")
	mapper := make(map[string]int)

	for _, word := range results {
		mapper[word]++
	}
	fmt.Printf("%t", mapper)
	fmt.Println(mapper)
}
