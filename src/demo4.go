package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "demo4", "test_demo")
}

func main() {
	flag.Parse()
	fmt.Println("Hello, " + name + "!")
	hello(name)
}
