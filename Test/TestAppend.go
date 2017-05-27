package main

import "fmt"

func main() {
	d := []string{"Welcome", "for", "Tianjin", "Have", "a", "good", "journey"}
	insertSlice := []string{"It", "is", "a", "big", "city"}
	insertSliceIndex := 3
	d = append(d[:insertSliceIndex], append(insertSlice, d[insertSliceIndex:]...)...)
	fmt.Printf("result:%v\n", d)

	a := []rune{'a','c','m','龢'}
	b := [10]rune{}
	a = append(a,b[:]...)
	fmt.Println(len(a))
	fmt.Println(a)
}
