package main

import "fmt"

func main(){
	wordWindow := []string{}
	wordWindow = append(wordWindow,"a")
	wordWindow = append(wordWindow,"b")
	fmt.Println(len(wordWindow))
	wordWindow = wordWindow[1:]
	fmt.Println(len(wordWindow))

	wordWindow = append(wordWindow,"c")

	for _, neighbor := range wordWindow{
		fmt.Println(neighbor)
	}
	fmt.Println(len(wordWindow))
}
