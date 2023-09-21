package list

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateSlise(len int) []int {
	slise := make([]int, len)
	fmt.Print()
	for i := range slise {
		slise[i] = rand.Intn(10)
	}
	return slise
}

func generateList(n int) (head *ListNode) {

	arg := generateSlise(n)
	var tempNode *ListNode
	for i := 1; i <= n; i++ {
		head = &ListNode{
			Val:  arg[n-i],
			Next: tempNode,
		}
		tempNode = head
	}
	return
}

func generateListArgs(arg ...int) (head *ListNode) {
	n := len(arg) - 1
	var tempNode *ListNode
	for i := range arg {
		head = &ListNode{
			Val:  arg[n-i],
			Next: tempNode,
		}
		tempNode = head
	}
	return
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%4.v ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func StartTask() {

	// head := generateList(7)
	// printList(head)
	// printList(task_19(head, 7))

	// head := generateListArgs(1, 4, 5, 23)
	// printList(head)
	// printList(task_206(head))

	// head1 := generateListArgs(1, 2, 3, 5)
	// printList(head2)
	// head2 := generateListArgs(1, 4, 5, 5)
	// printList(head2)
	// printList(task_21(head, head2))

	// head := generateListArgs(1, 2, 3, 3, 2, 1)
	// head := generateListArgs(1, 2, 3, 4, 3, 2, 1)
	// printList(head)
	// fmt.Println(task_234(head))

	head := generateListArgs(1, 2)
	printList(head)
	fmt.Println(task_141(head))
}
