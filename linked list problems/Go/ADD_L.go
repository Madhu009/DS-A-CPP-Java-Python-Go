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
	
	// var head* Node = createLinkedList()
	// printLinkedList(head)
	// add_at_pos(head)
	// del_at_pos(head)
	// del_match(head)
	// reverse(head)
	// remove_duplicates_1(head)
	// print_mid(head)
	// create_loop(head)
	// nth_from_last(head)
	// merge_2_nodewise(head,head)

	var head1* Node = createLinkedList()
	printLinkedList(head1)
	var head2* Node = createLinkedList()
	printLinkedList(head2)

	merge_2_nodewise(head1,head2)

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


func print_mid(head* Node) {
	slow:=head
	fast:=head
	for {
		if fast.next!=nil {
				
				fast = fast.next.next
				
				if fast==nil {
					fmt.Println("The middle node is ",slow.data)
					break
				}
				slow = slow.next

		} else {
				fmt.Println("The middle node is ",slow.data)
				break
			}
	}
}


func create_loop(head* Node) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the position to create a loop")
	scanner.Scan()
	text := scanner.Text()

	pos, err := strconv.Atoi(text)

	if err != nil {
		return
	}

	temp := head
	var loop* Node = nil
	curr_pos := 0

	for {

		if curr_pos==pos {
			loop = temp
		}

		if temp.next!=nil {
			temp = temp.next
			curr_pos+=1
		} else {
			temp.next = loop
			break
		}

	}

	if loop!=nil {
		fmt.Println("The loop is created and the loop node is ", loop.data)
		find_cycle(head)
		// printLinkedList(head)
	} else {
		fmt.Println("loop is not created")
		printLinkedList(head)

	}
}


func find_cycle(head* Node) {

	slow:=head
	fast:=head.next


	for {
		if fast != nil || fast.next!=nil {
			fast = fast.next.next
			slow = slow.next
			
			if slow == fast {
				fmt.Println(slow,fast)
				fmt.Println("the loop is found")
				break
			}
		} else  {
			fmt.Println("loop does not exist")
			break
		}
	}
}


func nth_from_last(head* Node){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the position to find the node from last")
	scanner.Scan()
	text := scanner.Text()

	pos, err := strconv.Atoi(text)
	curr_pos:=0

	if err!=nil {
		return
	}

	slow := head
	fast:=head
	for {
		if fast!=nil{

			for {
				if curr_pos > pos {
					break
				} else {
					fast = fast.next
					curr_pos+=1
				}
			}
				
			slow = slow.next
			fast = fast.next

		} else if slow!=head {
			fmt.Println("The nth from last is ",slow.data)
			break
		} else {
			fmt.Println("Out of range")
			break
		}
	}
}

func merge_2_nodewise(head1* Node,head2* Node) {

	curr1:=head1
	curr2:=head2
	var next1* Node = nil
	var next2* Node = nil

	for {

		if curr1!=nil && curr2!=nil {
			
			next1= curr1.next
			next2= curr2.next

			curr1.next =curr2
			curr2.next = next1

			curr1 = next1
			if next1!=nil{
				curr2 = next2
			} else {
				curr2.next = next2 
			}

		} else{
			break
		}
			
	}

	fmt.Println("Merged list is ")
	printLinkedList(head1)

}


func createLinkedList() *Node {

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

	return head
}
