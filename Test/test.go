package main
import "fmt"

func main() {
	length := len([]rune("你好,hello"))
	fmt.Println(length)
	fmt.Println(len("你好,hello"))
}
