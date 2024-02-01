package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Comment struct {
	id      int
	content string
	next    *Comment
}

func main() {

	var qte int
	fmt.Scanf("%d", &qte)
	outputs := make([]map[int]*Comment, qte)

	for i := 0; i < qte; i++ {
		var worker int
		fmt.Scanf("%d", &worker)
		outputs[i] = readGroup(worker)
	}

	i := 0
	for i < qte {
		groupComment(outputs[i], outputs[i][-1])
		if i == qte-1 {
			return
		}
		i++
		fmt.Print("\r\n")
	}

}

func readGroup(n int) map[int]*Comment {

	comentsHash := make(map[int]*Comment)
	scanner := bufio.NewScanner(os.Stdin)

	for it := 0; it <= n-1; it++ {
		scanner.Scan()
		line := scanner.Text()
		lineParam := strings.SplitN(line, " ", 3)
		id, _ := strconv.Atoi(lineParam[0])
		parentId, _ := strconv.Atoi(lineParam[1])

		temp := &Comment{
			id:      id,
			content: lineParam[2],
		}

		if head, ok := comentsHash[parentId]; !ok {
			comentsHash[parentId] = temp
		} else {
			comentsHash[parentId] = sortAscending(head, temp)
		}
	}
	return comentsHash
}

func sortAscending(head *Comment, temp *Comment) *Comment {

	if head.id > temp.id {
		temp.next = head
		return temp
	}

	if head.next != nil {
		head.next = sortAscending(head.next, temp)
		return head
	}

	head.next = temp
	return head
}

func groupComment(tree map[int]*Comment, head *Comment) {
	for head != nil {
		fmt.Println(head.content)
		treeComents(tree, tree[head.id], 3)
		if head.next != nil {
			fmt.Print("\r\n")
		}
		head = head.next
	}
}

func treeComents(tree map[int]*Comment, node *Comment, nesting int) {

	if node == nil {
		return
	}

	fmt.Println(outputsNode(nesting))
	fmt.Println(outputsNode(nesting) + "--" + node.content)

	if node.next != nil {
		treeComents(tree, tree[node.id], nesting<<1^1)
	} else {
		treeComents(tree, tree[node.id], nesting<<1^3)
	}

	treeComents(tree, node.next, nesting)

}

func outputsNode(x int) string {

	var res string

	if x == 1 {
		return ""
	}
	if x == 2 {
		return " "
	}
	if x == 3 {
		return "|"
	}

	if x&0x1 == 1 {
		res = "  |" + res
	} else {
		res = "   " + res
	}

	res = outputsNode(x>>1) + res

	return res
}
