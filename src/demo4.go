package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "demo4", "test_demo")
}

func main() {
	flag.Parse()
	println("Hello, " + name + "!")
	hello(name)
}