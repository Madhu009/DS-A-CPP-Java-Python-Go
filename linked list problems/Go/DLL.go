package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type Node struct {
	data int
	next *Node
	prev *Node

}


func create_DLL() *Node {

	scanner:= bufio.NewScanner(os.Stdin)
	var head* Node = nil
	
	fmt.Println("Please entet the data to create a DLL")

	for {
		scanner.Scan()
		text := scanner.Text()

		if text == "done" {
			break
		}

		data, err := strconv.Atoi(text)

		if err != nil {
			return head
		} 

		var new_node = Node{data,nil,nil}

		if head == nil {
			head = &new_node
		} else {

			temp := head

			for {

				if temp.next!=nil {
					temp = temp.next
				} else {
					break
				}
			}

			new_node.prev = temp
			temp.next = &new_node

		}

		fmt.Println("New node",data," is added to the linked list and type done to stop it")

	}
	return head
}

func printLinkedlist(head* Node) {
	var temp* Node = head
	var last* Node = nil

	fmt.Println("------------")
	for {
		if temp!=nil {
			last = temp
			fmt.Print(temp.data," ")
			temp = temp.next
		} else {
			break
		}
	}
	fmt.Println()
	for {
		if last!=nil {
			fmt.Print(last.data," ")
			last = last.prev
		} else {
			break
		}
	}
	fmt.Println()
	fmt.Println("------------")

}

func insert_pos(head* Node) *Node{
	fmt.Println("Please enter the data to insert")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := scanner.Text()
	data, err := strconv.Atoi(text)

	var new_node = Node{data,nil,nil}

	fmt.Println("Please enter the index where the node should be inserted")
	scanner.Scan()
	text_pos := scanner.Text()
	pos, err := strconv.Atoi(text_pos)

	if err!=nil {
		return head
	}
	curr_pos := 1

	if pos == 0 {
		new_node.next = head
		head.prev = &new_node
		head = &new_node
		fmt.Println("The node has been inserted successfully")
		return head
	}	
	var temp* Node = head

	for {
		if temp.next!=nil {

			if curr_pos ==pos {
				new_node.next = temp.next
				new_node.prev = temp
				temp.next.prev = &new_node
				temp.next = &new_node
				fmt.Println("The node has been inserted successfully")
				return head
			}
			temp = temp.next
			curr_pos+=1

			if temp.next==nil && curr_pos==pos {
				temp.next = &new_node
				new_node.prev = temp
				fmt.Println("The node has been inserted successfully")
				return head
			}
		} else {
			break
		}
	}

	fmt.Println("the index is out of range or invalid")
	return head

}

func delete_pos(head* Node) *Node{

	fmt.Println("Please enter the index where the node should be deleted")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := scanner.Text()
	pos, err := strconv.Atoi(text)
	curr_pos:=1
	if err!=nil {
		return head
	}

	if pos == 0 {
		head = head.next
		head.prev = nil
		
		fmt.Println("The node has been deleted successfully")
		return head
	}

	var temp* Node = head



	for {
		if temp.next!=nil {

			if curr_pos == pos {
				temp.next = temp.next.next
				if temp.next!=nil {
					temp.next.prev = temp
				}
				fmt.Println("The node has been deleted successfully")
				return head
			}
			temp = temp.next
			curr_pos+=1

			if temp.next==nil && curr_pos==pos {
				temp.prev.next = nil
				temp.prev = nil
				fmt.Println("The node has been deleted successfully")
				return head
			}
		} else {
			break
		}
	}

	fmt.Println("the index is out of range or invalid")
	return head


}

func reverse(head* Node) *Node {
	var prev* Node = nil
	var curr* Node = head
	var next* Node = nil

	for {
		if curr!=nil {
			next = curr.next
			curr.next = curr.prev
			curr.prev = next
			prev = curr
			curr = next
		} else {
			head = prev
			return head
		}
	}
}

func main() {
	head := create_DLL()
	printLinkedlist(head)
	// test := insert_pos(head)
	// printLinkedlist(test)
	test := delete_pos(head)
	printLinkedlist(test)
	test1 := reverse(head)
	printLinkedlist(test1)
}