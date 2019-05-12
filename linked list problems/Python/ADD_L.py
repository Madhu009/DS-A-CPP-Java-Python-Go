class Node:

	def __init__(self, data):
		self.data = data
		self.next = None


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
			temp = head
			while (temp.next!= None):
				temp = temp.next

			temp.next = Node(data)
		
		print("New node {} is added to the linked list and type done to stop it".format(data))

	printLinkedList(head)

def printLinkedList(head):
	temp = head
	while (temp!=None):

		print(temp.data,""end='')
		temp = temp.next

	print()


if __name__== "__main__":
  createLinkedList()