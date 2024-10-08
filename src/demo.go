package main

import (
	"fmt"
	"flag"
	"os"
)


var name string

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)


func init() {
	// flag.StringVar(&name, "name", "everyone", "The name of the person")

	cmdLine.StringVar(&name, "name", "everyone", "The name of the person")

}

func main() {

	// 重新设置命令行参数
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	// flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	// flag.CommandLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }


	cmdLine.Parse(os.Args[1:])
	// flag.Parse()

	fmt.Printf("Hello %s!\n", name)
}