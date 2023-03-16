package main

import (
	"fmt"
	"testing"
)

func TestChallengeOne(t *testing.T) {
	i := 21
	j := true
	var k float64 = 123.456
	fmt.Println(i)
	fmt.Printf("menampilkan tipe data i = %T \n", i)
	fmt.Printf("menampilkan tanda %s \n", "%")
	fmt.Printf("menampilkan nilai bolean j : true = %t \n", j)
	fmt.Printf("menampilkan nilai bolean j : false = %t \n", !j)
	fmt.Printf("menampilkan unicode rusia : Я (ya) = %s \n", string(rune('Я')))
	fmt.Printf("menampilkan nilai base 10 : 21 = %d \n", 21)
	fmt.Printf("menampilkan nilai base 8 : 25 = %o \n", 25)
	fmt.Printf("menampilkan nilai base 16 : f = %x \n", "f")
	fmt.Printf("menampilkan nilai base 16 : F = %x \n", "F")
	fmt.Printf("menampilkan unicode karakter Я : U+042F = %U \n", 'Я')
	fmt.Printf("menampilkan float : 123.456000 = %.6f \n", k)
	fmt.Printf("menampilkan float scientific : 1.234560E+02 = %E \n", k)
}
