package main

import (
	"fmt"
	"github.com/khangta/go-commons/greet"
	"github.com/khangta/go-commons/string"
)

func main() {
	fmt.Println(greet.Greeting("Khang"))

	fmt.Println(string.Trim(" Hello world!!!"))
}

