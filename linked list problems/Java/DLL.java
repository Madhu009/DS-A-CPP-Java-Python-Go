import java.util.*;  


class Node {

	int data;
	Node next;
	Node prev;

	Node(int data)
	{
		this.data = data;
		this.next = null;
		this.prev=null;
	}

}

class DLL
{
	public static void main(String args[])
	{
		Node head = create_DLL();
		printLinkedlist(head);

		Node test = insert_pos(head);
		printLinkedlist(test);
		Node test1 = delete_pos(head);
		printLinkedlist(test1);
		Node test2 = reverse(head);
		printLinkedlist(test2);
	}

	public static Node create_DLL()
	{
		Scanner sc = new Scanner(System.in);
		Node head = null;
		boolean done = false;
		System.out.println("please enter the data to create a DLL ");
		while(done==false)
		{

			String input = sc.next();
			if(input.equals("done"))
				break;
			int data = Integer.parseInt(input);
			Node new_node = new Node(data);

			if (head == null)
			{
				head = new_node;
			}
			else
			{
				Node temp = head;
				while(temp.next!=null)
					temp = temp.next;

				new_node.prev = temp;
				temp.next = new_node;
			}

			System.out.println("New node "+ data +" is added to the linked list and type done to stop it");
		}

		return head;

	}

	public static void printLinkedlist(Node head)
	{
		Node temp = head;
		Node last = null;
		System.out.println("---------------");
		while(temp!=null)
		{
			last = temp;
			System.out.print(temp.data+ " ");
			temp = temp.next;
		}
		System.out.println();

		while(last!=null)
		{
			System.out.print(last.data+ " ");
			last = last.prev;
		}
		System.out.println();
		System.out.println("---------------");

	}

	public static Node insert_pos(Node head)
	{
		Scanner sc = new Scanner(System.in);

		System.out.println("Please enter the data to insert ");
		int data = sc.nextInt(); 
		System.out.println("Plese enter the index where to be inserted");
		int pos = sc.nextInt();
		
		Node new_node = new Node(data);
		Node temp = head;

		int curr_pos = 1;

		if (pos == 0)
		{
			new_node.next = head;
			head.prev = new_node;
			head = new_node;
			return head;
		}

		while(temp.next!=null)
		{
			if (curr_pos== pos)
			{
				new_node.next = temp.next;
				temp.next.prev = new_node;
				new_node.prev = temp;
				temp.next = new_node;
				System.out.println("The data has been inserted successfully");
				return head;
			}

			temp= temp.next;
			curr_pos+=1;

			if (temp.next==null && curr_pos==pos)
			{
				temp.next = new_node;
				new_node.prev = temp;
				System.out.println("The data has been inserted successfully");
				return head;
			}
		}

		System.out.println("The index is out of range or invalid");
		return head;
	}

	public static Node delete_pos(Node head)
	{
		System.out.println("Please enter the index where the node should be deleted");
		Scanner sc = new Scanner(System.in);
		int pos = sc.nextInt();

		int curr_pos = 1;

		if (pos == 0)
		{
			head = head.next;
			head.prev = null;
			System.out.println("The data has been deleted successfully");
			return head;
		}

		Node temp = head;

		while(temp.next!=null)
		{
			if (curr_pos==pos)
			{
				temp.next = temp.next.next;
				if (temp.next!=null)
					temp.next.prev = temp;
				System.out.println("The data has been deleted successfully");
				return head;
			}
			temp = temp.next;
			curr_pos +=1;

			if (temp.next==null && curr_pos==pos)
			{
				temp.prev.next = null;
				temp.prev = null;
				System.out.println("The data has been delteted successfully");
				return head;
			}
		}

		System.out.println("The index is out of range or invalid");
		return head;
	}

	public static Node reverse(Node head)
	{
		Node prev = null;
		Node next = null;
		Node curr = head;

		while(curr!=null)
		{
			next = curr.next;
			curr.next = curr.prev;
			curr.prev = next;
			prev = curr;
			curr = next;
		}

		head = prev;
		return head;
	}
}