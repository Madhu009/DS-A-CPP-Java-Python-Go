class Node:

	def __init__(self, data):
		self.data = data
		self.next = None

def add_at_pos (head):
	print("please enter the data to insert")
	data = int(input())
	print("Please enter the positon")
	pos = int(input())
	curr_pos = 0

	temp = head
	new_node = Node(data)

	while temp!=None:

		if pos == 0:
			new_node.next= head
			head = new_node
			break
		if curr_pos+1 == pos:
			new_node.next = temp.next
			temp.next = new_node
			curr_pos+= 1
			break
		else:
			temp = temp.next
			curr_pos+= 1

	if pos != curr_pos or  pos<0:
		print("The index is out of bound")
	else:
		print("The data has been inserted at the positon")

	printLinkedList(head)

def del_at_pos(head):

	print("Please enter the positon to delete")
	pos = int(input())
	curr_pos = 0

	temp = head


	while temp.next is not None:
		
		if pos == 0:
			head = head.next
			break

		if pos == curr_pos+1:
			temp.next = temp.next.next
			curr_pos+=1
			break	
		else:
			temp= temp.next
			curr_pos+= 1
	
	if pos<0 or curr_pos!=pos:
		print("the index is out of bound")
	else:
		print("The data at the positon is deleted")
	printLinkedList(head)


def del_match(head):
	print("Please enter the data to delete")
	data = int(input())
	curr_pos = 0

	temp = head
	if head.data == data:
		head = head.next
		print("The data is found at {} and deleted".format(curr_pos))
		printLinkedList(head)
		return None
	while temp.next is not None:
		if temp.next.data == data:
			curr_pos+=1
			temp.next = temp.next.next
			print("The data is found at {} and deleted".format(curr_pos))
			printLinkedList(head)
			return None
		else:
			temp= temp.next
			curr_pos+=1
	print("The data is not found in the list")


def createLinkedList():

	head,done = None, True
	print("Please continue entering the data to the linkedlist and type done to stop it")

	while done:
		
		inp = input()

		if inp == "done":
			break

		data = int(inp)

		if head == None:
			head = Node(data)
		else:
			temp =head
			while (temp.next!= None):
				temp = temp.next

			temp.next = Node(data)

		print("New node {} is added to the linked list and type done to stop it".format(data))

	# printLinkedList(head)
	# add_at_pos(head)
	# del_at_pos(head)
	# del_match(head)
	# reverse(head)
	# remove_duplicates(head)
	# print_mid(head)
	return head




def print_mid(head):
	slow, fast = head, head
	while fast.next!=None:
		slow = slow.next
		fast = fast.next.next
		if fast == None:
			break
	print("the middle element of the list is ",slow.data)

def printLinkedList(head):
	temp = head
	while (temp!=None):

		print(temp.data,"",end='')
		temp = temp.next

	print()

def reverse(head):
	curr,next,prev = head,None,None

	while curr.next!=None:
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next

	curr.next = prev
	head = curr
	print("the reverse list is ")
	printLinkedList(head)

def remove_duplicates(head):
	temp, curr = head,head

	while curr.next!=None:
		while temp.next!=None:
			if curr.data == temp.next.data:
				temp.next = temp.next.next
			else:
				temp = temp.next
		if curr.next!=None:
			curr = curr.next
			temp = curr
	print("The final list is ")
	printLinkedList(head)

def create_loop(head):
	
	loop_node, temp = None,head
	print("please enter the index where you wanna create a loop")
	curr_pos = 0
	index = int(input())

	while temp.next!=None:
		if curr_pos==index:
			loop_node = temp
		temp = temp.next
		curr_pos+=1

	if curr_pos==index:
		loop_node = temp

	temp.next = loop_node
	print("the loop node is ", loop_node.data)
	find_loop(head)

def find_loop(head):
	slow,fast = head,head.next
	while fast.next!=None:
		slow = slow.next
		fast = fast.next.next
		if fast == slow:
			print("The loop is found at ",fast.data)
			break

if __name__== "__main__":
  head = createLinkedList()
  create_loop(head)
