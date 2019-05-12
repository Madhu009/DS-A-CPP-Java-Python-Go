import java.util.*;  

class Node
{
	int data;
	Node next;

	Node(int data)
	{
		this.data = data;
		this.next = null;
	}
}

class ADD_L 
{

	public static void main(String args[])
	{
		createLinkedList();
	}

	public static void createLinkedList()
	{
		Node head = null;
		System.out.println("Please continue entering the data to the linkedlist and type done to stop it ");

		Scanner scanner = new Scanner(System.in);
		boolean done = false;
		int data;

		while (!done)
		{
			String input = scanner.next();
			if (input.equals("done"))
				break;

			data = Integer.parseInt(input);

			if (head == null)
			{
				Node temp = new Node(data);
				head = temp;
			}
			else@
			{
				Node temp = head;
				while(temp.next!=null)
				{
					temp=temp.next;
				}

				Node new_node = new Node(data);
				temp.next = new_node;
			}
			System.out.println("New node "+ data +" is added to the linked list and type done to stop it");
		}

		printLinkedList(head);
		
	}

	public static void printLinkedList(Node head)
	{
		Node temp = head;

		while(temp!= null)
		{
			System.out.print(temp.data+ " ");
			temp= temp.next;
		}
		System.out.println();
	}
}

