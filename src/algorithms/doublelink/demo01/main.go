package main

import (
	"fmt"
)

// HeroNode is struct
type HeroNode struct {
	no       int
	name     string
	nickname string
	prev     *HeroNode // 指向前一个节点
	next     *HeroNode // 指向下一个节点
}

// 给双向链表插入一个节点
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
	newHeroNode.prev = temp
}

// InsertHeroNodeByOrder is a function
func InsertHeroNodeByOrder(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true

	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("对象已经存在,no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.prev = temp
		if temp.next != nil {
			temp.next.prev = newHeroNode
		}
		temp.next = newHeroNode
	}
}

// DelHeroNode is function
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false

	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			// exists
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.prev = temp
		}
	} else {
		fmt.Println("Can not delete the node which is not exists")
	}

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

func ListHeroNodeRerverse(head *HeroNode) {
	temp := head

	// 判断链表是否为空
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}

	// 让temp定位到双向链表的最后节点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 遍历链表
	for {
		fmt.Printf("[%d, %s, %s] --> ", temp.no, temp.name, temp.nickname)

		// 判断是否是链表头
		temp = temp.prev
		if temp.prev == nil {
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
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	// InsertHeroNode(head, hero1)
	// InsertHeroNode(head, hero2)
	// InsertHeroNode(head, hero3)

	InsertHeroNodeByOrder(head, hero1)
	InsertHeroNodeByOrder(head, hero2)
	InsertHeroNodeByOrder(head, hero3)
	ListHeroNode(head)
	fmt.Println()
	ListHeroNodeRerverse(head)

	DelHeroNode(head, 3)
	fmt.Println()
	ListHeroNode(head)
	fmt.Println()
	ListHeroNodeRerverse(head)
}
