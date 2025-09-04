package main

import (
	"fmt"
)

func main() {
	var x interface{} = make(chan any)
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan any:
		fmt.Println("chan")
	default:
		fmt.Println("unsupported")
	}
}
