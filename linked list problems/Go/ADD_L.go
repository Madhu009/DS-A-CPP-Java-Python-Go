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

func del_at_pos(head* Node) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the position to delete");
	scanner.Scan()
	text := scanner.Text()
	pos, err := strconv.Atoi(text)
	temp := head
	if err != nil {
		return
	}
	curr_pos := 0

	for {

		if pos == 0 {
			head = head.next
			break
		}
		if temp.next !=nil {
			if curr_pos+1 == pos {
				curr_pos+=1
				temp.next = temp.next.next
				break
			} else {
				temp = temp.next
				curr_pos+=1
			}
		} else {
			break
		}
	}

	if pos<0 || curr_pos !=pos {
		fmt.Println("the index is out of bound")
	} else {
		fmt.Println("the data at the position is deleted")
	}

	printLinkedList(head)

}
func add_at_pos(head* Node) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the data to insert");
	scanner.Scan()
	text := scanner.Text()
	data, err := strconv.Atoi(text)
	fmt.Println("Please enter the position")
	scanner.Scan()
	pos_text := scanner.Text()
	pos, err := strconv.Atoi(pos_text)
	
	if err != nil {
		return
	}
	curr_pos := 0
	var new_node = Node{data,nil}
	temp := head

	for {
		if pos == 0 {
			new_node.next = head
			head = &new_node
			break
		}

		if temp != nil {

			if pos == curr_pos+1 {
				new_node.next = temp.next
				temp.next = &new_node
				curr_pos+=1
				break
			} else {
				temp = temp.next
				curr_pos+=1
			}
		} else {
			break
		}
	}

	if pos != curr_pos || pos<0 {
		fmt.Println("The index is out bound")
	}else {
		fmt.Println("The data has been inserted successfully")
	}
	printLinkedList(head)

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
	// add_at_pos(head)
	// del_at_pos(head)
	// del_match(head)
	// reverse(head)
	remove_duplicates_1(head)
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

func del_match(head* Node) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the data to delete");
	scanner.Scan()
	text := scanner.Text()
	data, err := strconv.Atoi(text)

	if err != nil {
		return
	}

	curr_pos := 0
	temp := head

	if head.data == data {
		head = head.next
		fmt.Println("The data is found at ",curr_pos," and deleted")
		printLinkedList(head)
		return
	}

	for {
		if temp.next!=nil {

			if temp.next.data == data {
				curr_pos+=1;
				temp.next = temp.next.next
				fmt.Println("The data is found at ",curr_pos, " and deleted")
				printLinkedList(head)
				return
			} else {
				temp = temp.next
				curr_pos+=1
			}
		} else {
			break
		}
	}

	fmt.Println("The data is not found in the list")
}

func reverse(head* Node) {

	var next* Node = nil
	var prev* Node = nil

	curr:= head

	for {
		if curr.next!=nil {
			next = curr.next
			curr.next = prev
			prev = curr
			curr = next
		} else {
			curr.next = prev
			head = curr
			break
		}
	}

	fmt.Println("The reversed list is ")
	printLinkedList(head)
}


func remove_duplicates_1(head* Node) {

	temp:= head
	curr:= head

	for {
		if curr.next!=nil {
			for {
				fmt.Println("curr2",curr,temp)
				if temp.next!= nil{
					if curr.data == temp.next.data {
						temp.next = temp.next.next
					} else {
						temp = temp.next
					}
				} else {
					break
				}
			}
			if curr.next!=nil {
				curr = curr.next
				temp = curr
			}
			
		} else {
			break
		}
	}

	fmt.Println("The final list is ")
	printLinkedList(head)
}
