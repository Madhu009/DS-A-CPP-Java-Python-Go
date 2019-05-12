package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type Node struct {
	data int
	next *Node
}
func main() {
	createLinkedList()
}

func createLinkedList() {

	var head* Node = nil
	fmt.Println("Please continue entering the data to the linkedlist and type done to stop it\n");

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()

		if text == "done" {
			break
		}
		data, err := strconv.Atoi(text)
		
		if err == nil {
        	var new_node = Node{data,nil}

        	if head == nil {
        		head = &new_node
        	} else {
        		temp := head
        		for{
        			if temp.next==nil {
	        			temp.next = &new_node
	        			break
        			}
        			temp = temp.next
        			
        		}
        	}
		fmt.Println("New node",data," is added to the linked list and type done to stop it")

    	}

	}

	printLinkedList(head)
}

func printLinkedList(head *Node){

	var temp *Node = head

	for {
		if temp==nil {
			break
		}
		fmt.Print(temp.data," ")
		temp = temp.next
	}
	fmt.Println()

}