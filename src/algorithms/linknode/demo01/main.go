package main

import (
	"fmt"
)

// HeroNode is struct
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 指向下一个节点
}

// 给链表插入一个节点
// 编写第一种插入方法，在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil { //表示找到了最后
			break
		}
		temp = temp.next // 使temp不断指向下一个节点
	}
	// 将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}

func ListHeroNode(head *HeroNode) {
	temp := head

	// 判断链表是否为空
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}

	// 遍历链表
	for {
		fmt.Printf("[%d, %s, %s] --> ", temp.next.no, temp.next.name, temp.next.nickname)

		// 判断是否是链表尾
		temp = temp.next
		if temp.next == nil {
			break
		}
	}

}

func main() {
	// 创建头节点
	head := &HeroNode{}

	// 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "吴用",
		nickname: "智多星",
	}

	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)

	ListHeroNode(head)
}
