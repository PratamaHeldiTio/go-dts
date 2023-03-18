package main

import (
	"fmt"
	"testing"
)

func TestChallengeTwo(t *testing.T) {
	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i =", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			name := "Heldi Tio Pratama"
			for i, unicode := range name {
				if i%2 == 0 {
					fmt.Printf("character %#U starts at byte position %d \n", unicode, i)
				}
			}
		}
		fmt.Println("Nilai j =", j)
	}
}
