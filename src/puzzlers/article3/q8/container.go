package main
import "fmt"

// continer 包中包含的容器
import "container/list"
// List 实现了双向链表, Element 实现了链表的节点
// List 中的四种方法 MoveBefore(), MoveAfter(), MoveToFront(), MoveToBack()
// 上面的四种方法都是 *Element 类型的方法, 用来在链表中操作元素

import "container/ring"
// Ring 实现了循环链表, List在内部就是一个循环链表

/**
	List和ring的区别
	1. Ring类型的数据结构仅由它自身即可代表,而List类型则需要有他以及Element类型才能代表
	2. Ring类型的值严格来讲,只代表了其所属的循环链表厚葬的一个元素,而一个List类型的值则代表了一个完整的链表,维度上不同
	3. 在创建和初始化一个Ring的时候,需要指定一个环的大小,对于一个List值来说不能也没有必要.Ring一旦创建,其长度就确定了
	4. 仅通过var r ring.Ring语句声明的的r变量是一个大小为1的循环链表,而通过var l list.List语句声明的l变量是一个大小为空的双向链表
		List中的根元素不存储任何值,所以在极端长度时并不会包含他
	5. Ring类型的Len()方法时间复杂度是O(n),而List类型的Len()方法的时间复杂度是O(1),原因是在List struct中,有一个单独的变量来记录链表的长度
**/

func printList(l *list.List) {
	for e := l.Front(); e!= nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func testList() {
	List := list.List{}
	List.PushBack(1)
	List.PushBack(2)
	List.PushFront(4)
	e := List.Front()
	e = e.Next()
	List.MoveBefore(e, List.Front())
	printList(&List)
}

func testList2() {
	var l list.List // 声明一个大小为0的链表, 可以开箱即用
	// Go语言中的链表具有延迟初始化操作,仅在实际需要的时候才进行初始化操作
	l.Front() // 这里会先判断链表的长度,如果是0,会进行初始化操作,之后返回nil

}

func testRing() {
	r := ring.Ring{}
	r.Value = 1
	fmt.Println("Ring Len :=", r.Len())
	fmt.Println("Value of Ring :=", r.Value)
}


func main() {

	testList()

	testList2()

	testRing()

	fmt.Println("Hello, world!");

}