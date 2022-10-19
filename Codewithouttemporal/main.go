package main

import (
	"fmt"
	"os"
	greeting "github.com/genos1998/Temporal-Research/Codewithouttemporal/greeting/greeting"
)

func main() {
	name := os.Args[1]
	greeting := greeting.GreetSomeone(name)
	fmt.Println(greeting)
}
