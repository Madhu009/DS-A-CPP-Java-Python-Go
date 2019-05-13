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

	printLinkedList(head)
	# add_at_pos(head)
	# del_at_pos(head)
	del_match(head)

def printLinkedList(head):
	temp = head
	while (temp!=None):

		print(temp.data,"",end='')
		temp = temp.next

	print()


if __name__== "__main__":
  createLinkedList()