package main

import (
	"fmt"
	"puzzlers/article3/q2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The Greeting of object")
}

func main() {
	flag.Parse()
	lib5.Hello(name)
}


