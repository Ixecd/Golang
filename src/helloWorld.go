package main

import (
	"fmt"
	"flag" // 命令行参数
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// 经过测试之后发现,不加下面这行代码也会执行成功,说明Go语言本身在main函数执行之前已经对命令行参数进行了解析
	flag.Parse() // 解析命令行参数
    fmt.Printf("hello, %s!\n", name)
}