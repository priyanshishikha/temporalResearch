package main

import (
	"fmt"
	"os"

	greeting "greeting/greeting"
)

func main() {
	name := os.Args[1]
	greeting := greeting.GreetSomeone(name)
	fmt.Println(greeting)
}
