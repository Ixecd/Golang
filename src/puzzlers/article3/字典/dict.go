package main

import "fmt"
// 字典,存储的是KV键值对,本质上是一个哈徐表的特定实现,键的类型是受限的,而元素却可以是任意类型的
// map中的key必须是可以比较的,也就是!= ==
// Go语言中键值不可以是 函数/字典/切片类型
func testMap() {
	aMap := map[string]int{}
	aMap["one"] = 1
	aMap["two"] = 2
	aMap["three"] = 3
	fmt.Println(aMap)

	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Println(k, ":", v)
	} else {
		fmt.Println("Not found!")
	}
}

func testMap2() {
	// 这里interface是指针类型,可以存储任意类型的值
	// 注意:最好不要把字典的键类型设置为任意接口类型,如果非要这么做,那么一定要确保带吧在可用的范围之内
	aMap := map[interface{}]int{
		"1" : 1,
		// []int{2} : 2, // panic
		3 : 3,
	}
	fmt.Println(aMap)
}

// Q: 在值为nil的字典上执行读操作会成功吗,写操作呢?
// A: 字典是引用类型,所以当我们声明而不初始化一个字典类型的变量的时候,它的值会是nil
// 除了添加键值对,我们在一个值为nil的字典上做任何操作都不hi引起错误,当我们试图在一个值为nil的字典中添加键值的时候,Go语言的运行是系统就会立即抛出一个panic
func testMap3() {
	// 声明一个dict,这里是nil
	var aMap map[string]int
	aMap = make(map[string]int)
	if aMap == nil {
		fmt.Println("aMap is nil")
	}
	aMap["one"] = 1
	fmt.Println(aMap)
	aMap["two"] = 2 // panic: assignment to entry in nil map
}

func main() {
	testMap()
	testMap2()
	testMap3()
}



