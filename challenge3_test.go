package main

import (
	"fmt"
	"testing"
)

func TestChallangeThree(t *testing.T) {
	word := "selamat malam"

	data := make(map[string]int)

	for _, value := range word {
		fmt.Println(string(value))
		data[string(value)]++
	}
	fmt.Println(data)

}
