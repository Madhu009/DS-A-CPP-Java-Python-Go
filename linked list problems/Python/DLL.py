
class Node ():

	def __init__(self,data):
		self.data = data
		self.next = None
		self.prev = None

def create_DLL():
	print("Please enter the data to create a DLL")
	done = False
	head = None

	while not done:
		data = input()

		if data =="done":
			break
		data = int(data)
		new_node = Node(data)

		if head == None:
			head = new_node
		else:
			temp = head

			while temp.next!=None:
				temp = temp.next
			
			temp.next = new_node
			new_node.prev = temp

		print("The data has been added type done to complete the list")

	return head


def printLinkedList(head):
	
	temp = head
	last = None
	print("-------------")
	while temp!=None:
		last = temp
		print(temp.data,"",end='')
		temp = temp.next
	print()

	while last!=None:
		print(last.data,"",end='')
		last = last.prev
	print()
	print("-------------")

def insert_pos(head):
	print("Please enter the data to be inserted")
	data = int(input())
	print("Please enter the index to be inserted")
	pos = int(input())

	curr_pos = 1

	new_node = Node(data)
	temp = head

	if pos==0:
		new_node.next = head
		head.prev = new_node
		head = new_node
		print("The node has been inserted at ",pos," index");
		return head

	while temp.next!=None:

		if curr_pos == pos:
			new_node.next = temp.next
			temp.next.prev = new_node
			temp.next = new_node
			new_node.prev = temp
			print("The node has been inserted at ",pos," index");
			return head
		temp = temp.next
		curr_pos +=1

		if temp.next==None and curr_pos==pos:
			temp.next = new_node
			new_node.prev = temp
			print("The node has been inserted at ",pos," index");
			return head

	print("index is out of range or invalid")
	return head

def delete_pos(head):
	print("please enter the index where the node to be deleted")
	pos = int(input())
	curr_pos =1

	if pos == 0:
		head.next.prev = None
		head = head.next
		print("The node has been deleted at ",pos," index");
		return head

	temp = head

	while temp.next!=None:

		if curr_pos ==pos:
			temp.next = temp.next.next
			if temp.next!=None:
				temp.next.prev = temp
			print("The node has been deleted at ",pos," index");
			return head
		
		temp = temp.next;
		curr_pos += 1

		if temp.next==None and curr_pos==pos:
			temp.prev.next = None
			temp.prev = None
			print("The node has been deleted at ",pos," index");
			return head

	print("index is out of range or invalid")
	return head

def reverse(head):

	prev, curr, next = None,head,None

	while curr!=None:
		next = curr.next
		curr.next = curr.prev
		curr.prev = next
		prev = curr
		curr = next

	head = prev

	return head


if __name__ == "__main__":
	head = create_DLL()
	printLinkedList(head)
	test = reverse(head)
	print("The reversed list is")
	printLinkedList(test)