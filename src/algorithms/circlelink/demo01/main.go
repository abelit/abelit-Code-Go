package main

import (
	"fmt"
)

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	// 判断是不是第一次猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成环形
		return
	}

	// 定义一个临时变量，找到环形的最后节点
	temp := head

	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	// 加入环形链表
	temp.next = newCatNode
	newCatNode.next = head

}

func ListCircleLink(head *CatNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("空链表")
		return
	}

	for {
		fmt.Printf("猫的信息是=%s %s ", temp.name, "->")
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func DelCircleLink(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("空链表")
		return head
	}

	// 只有一个节点
	if temp.next == head {
		temp.next = nil
		return head
	}

	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true

	// 多个节点
	for {
		if temp.next == head { //还差最后一个需要比较
			break
		}

		if temp.no == id { //找到了
			if temp == head { //删除头节点

			}
			helper.next = temp.next
			fmt.Println("不存在", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Println("猫猫", id)
		} else {
			fmt.Println("不存在", id)
		}
	}
	return head
}

func main() {
	// 初始化环形链表的头节点
	head := &CatNode{}

	// 创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	ListCircleLink(head)

	// DelCircleLink(head, 1)
	// fmt.Println()
	// ListCircleLink(head)
}
