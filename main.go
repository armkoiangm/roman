package main

import (
	"armen/roman/roman"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		input := os.Args[1]

		rc := roman.NewConverter()
		result := rc.Convert(input)

		fmt.Println(result)
		fmt.Println(input)
	}
}
