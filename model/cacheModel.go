package model

import "fmt"

const size = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add %s\n", n.Val)
	temp := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n
	c.Queue.Length++
	if c.Queue.Length > size {
		c.Remove(c.Queue.Tail.Left)
	}
}
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove : %s\n", n.Val)
	left := n.Left
	right := n.Right

	right.Left = left
	left.Right = right
	c.Queue.Length -= 1

	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d-[", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}
