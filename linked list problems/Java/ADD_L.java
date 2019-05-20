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
		Node head = createLinkedList();
		System.out.println("the head node is "+head.data);
		create_loop(head);

	}

	public static Node createLinkedList()
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
			else
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
		// add_at_pos(head);
		// del_at_pos(head);
		// del_match(head);
		// find_match_position(head);
		// reverse(head);
		// remove_duplicates(head);
		// print_mid(head);
		return head;
		
	}

	public static void reverse(Node head)
	{
		Node curr = head;
		Node next = null;
		Node prev = null;

		while(curr.next!=null)
		{
			next = curr.next;
			curr.next = prev;
			prev = curr;
			curr = next;
		}
		curr.next = prev;
		head = curr;
		System.out.print("The reverse list is ");
		printLinkedList(head);
	}

	public static void print_mid(Node head) {
		Node slow = head;
		Node fast = head;
		while(fast.next!=null)
		{
			slow = slow.next;
			fast = fast.next.next;
			if (fast==null)
				break;
		}
		System.out.println("The middle element of the list is "+ slow.data);
	}
	public static void add_at_pos(Node head)
	{	
		Scanner scanner = new Scanner(System.in);
		System.out.println("Please enter the data to insert");
		int data = scanner.nextInt();
		System.out.println("Please enter the postion");
		int pos = scanner.nextInt();

		Node temp = head;
		int curr_pos = 0;
		Node new_node = new Node(data);

		while(temp!=null)
		{
			if (pos == 0)
			{
				new_node.next = head;
				head = new_node;
				break;
			}
			
			if (pos==curr_pos+1)
			{
				Node prev = temp.next;
				temp.next = new_node;
				new_node.next = prev;
				curr_pos+=1;
				break;
			}
			else
			{
				temp=temp.next;
				curr_pos+= 1;
			}

		}

		if((curr_pos!=pos) || (pos<0))
		{
			System.out.println("The index is out of bound");
		}
		else
		{
			System.out.println("The data has been inserted successfully");
		}

		printLinkedList(head);

	}

	public static void del_at_pos(Node head)
	{
		Scanner scanner = new Scanner(System.in);
		System.out.println("Please enter the postion to delete an item");
		int pos = scanner.nextInt();
		int curr_pos = 0;

		Node temp = head;

		while(temp.next!=null)
		{

			if (pos==0)
			{
				head = head.next;
				break;
			}

			if (curr_pos+1 == pos)
			{
				temp.next = temp.next.next;
				curr_pos+= 1;
				break;
			}
			else
			{
				temp = temp.next;
				curr_pos+= 1;
			}
		}

		if(curr_pos<0 || curr_pos!=pos)
			System.out.println("Index is out of bound");
		else
			System.out.println("The item at the position is deleted");

		printLinkedList(head);
	}

	public static void del_match(Node head)
	{
		Scanner scanner = new Scanner(System.in);
		System.out.println("Please enter the data to delete");
		int data = scanner.nextInt();

		if (head.data == data)
		{
			head = head.next;

			System.out.println("Data has been found and deleted");
			printLinkedList(head);
			return;
		}

		Node temp = head;
		
		while(temp.next!=null)
		{
			if (temp.next.data == data)
			{
				temp.next = temp.next.next;
				System.out.println("Data has been found and deleted");
				printLinkedList(head);
				return;
			}
			else
			{
				temp = temp.next;
			}
		}

		System.out.println("data is not found in the list");
	}

	public static void find_match_position(Node head)
	{
		Scanner scanner = new Scanner(System.in);
		System.out.println("Please enter the data to find the position");
		int data = scanner.nextInt();

		Node temp = head;
		int pos = 0;

		while(temp!=null)
		{
			if (temp.data == data)
			{
				System.out.println("Data has been found at pos "+pos);
				printLinkedList(head);
				return;
			}
			else
			{
				temp = temp.next;
				pos+=1;
			}
		}

		System.out.println("data is not found in the list");
	}


	public static void remove_duplicates(Node head)
	{
		Node curr = head;
		Node temp = head;

		while (curr.next!=null)
		{
			while(temp.next!=null)
			{
				if (temp.next.data == curr.data)
				{
					temp.next = temp.next.next;
				}
				else
				{
					temp = temp.next;
				}
			}

			if (curr.next!=null)
			{
				curr = curr.next;
				temp = curr;
			}
		}
		System.out.println("The final list is ");
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

	public static void create_loop(Node head)
	{
		int curr_pos =0;
		int index;
		Node loop_node = null;
		Node temp=head;
		System.out.println("Please enter the index where you wanna create a loop");
		Scanner sc = new Scanner(System.in);
		index = sc.nextInt();

		while(temp!= null)
		{
			if (curr_pos==index)
				loop_node = temp;

			temp = temp.next;
			curr_pos+=1;

			if (temp.next==null)
			{
				if (curr_pos==index)
					loop_node = temp;
				temp.next = loop_node;
				break;
			}
		}

		System.out.println("The loop node is "+loop_node.data);
		find_loop(head);
	}

	public static void find_loop(Node head)
	{
		Node slow = head;
		Node fast = head.next;
		while(fast.next!=null)
		{
			slow = slow.next;
			fast=fast.next.next;

			if(fast==slow)
			{
				System.out.println("The loop is found at "+fast.data);			
				break;
			}
		}
	}
}

