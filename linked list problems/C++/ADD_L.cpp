#include <iostream>

using namespace std;

struct Node
{
	int data;
	struct Node *next = NULL;
};

void printLinkedList(struct Node* head )
{
	struct Node* itr = head;
			
	while(itr!=NULL)
	{
		cout<<itr->data<<" ";
		itr = itr->next;
	}
	cout<<endl;
}

void reverse(struct Node* head)
{
	struct Node* curr = head;
	struct Node* prev = NULL;
	struct Node* next = NULL;

	while(curr->next!=NULL)
	{
		next = curr->next;
		curr->next = prev;
		prev = curr;
		curr = next;
	}
	curr->next = prev;
	head = curr;
	cout<<"The reversed list is "<<endl;
	printLinkedList(head);
}

void add_at_pos(struct Node *head)
	{
		cout<<"Please enter the data to insert"<<endl;
		int data;
		cin>>data;
		cout<<"please enter the position"<<endl;
		int pos;
		cin>>pos;

		struct Node *temp = head;
		int curr_pos = 0;
		struct Node *new_node = (struct Node*) malloc (sizeof(struct Node));
		new_node->data = data;
		new_node->next = NULL;

		while(temp!=NULL)
		{
			if (pos==0)
			{
				new_node->next = head;
				head = new_node;
				break;
			}

			if(curr_pos+1==pos)
			{
				new_node->next = temp->next;
				temp->next = new_node;
				curr_pos+=1;
				break;
			}
			else
			{
				temp = temp->next;
				curr_pos += 1;
			}

		}

		if ( (pos<0) || (curr_pos!=pos) )
			cout<<"The index is out bound"<<endl;
		else
			cout<<"The data has been inserted successfully"<<endl;
	
		printLinkedList(head);

	}


void del_at_pos(struct Node *head)
{

	cout<<"please enter the position to delete an item"<<endl;
	int pos;
	cin>>pos;
	int curr_pos = 0;

	struct Node *temp = head;

	while(temp->next!=NULL)
	{
		if (pos==0)
		{
			head = head->next;
			break;
		}
		if (curr_pos+1==pos)
		{
			temp->next = temp->next->next;
			curr_pos+= 1;
			break;
		}
		else
		{
			temp= temp->next;
			curr_pos+= 1;
		}
	}

	if ((pos<0) || (curr_pos!=pos) )
		cout<<"the index is out of bound"<<endl;
	else
		cout<<"The data at the position is deleted"<<endl;

	printLinkedList(head);

}

void del_match(struct Node *head)
{
	cout<<"please enter the data to delete"<<endl;
	int data;
	cin>>data;
	int pos=0;
	if(head->data == data)
	{
		head = head->next;
		cout<<"The data has been found at pos "<<pos<<" and deleted"<<endl;
		printLinkedList(head);
		return;
	}
	struct Node *temp = head;

	while(temp->next!=NULL)
	{

		if (temp->next->data == data)
		{	
			pos+=1;
			temp->next = temp->next->next;
			cout<<"The data has been found at pos "<<pos<<" and deleted"<<endl;
			printLinkedList(head);
			return;
		}
		else
		{
			temp = temp->next;
			pos+= 1;

		}		
	}
	
	cout<<"data is not found in the list"<<endl;

}

void remove_duplicates(struct Node *head)
{
	struct Node *curr = head;
	struct Node *temp = head;

	while (curr->next!=NULL)
	{
		while(temp->next!=NULL)
		{
			if (temp->next->data == curr->data)
			{
				temp->next = temp->next->next;
			}
			else
			{
				temp = temp->next;
			}
		}

		if (curr->next!=NULL)
		{
			curr = curr->next;
			temp = curr;
		}
	}
	cout<<"The final list is "<<endl;
	printLinkedList(head);
}

void print_mid(struct Node *head) 
{
	struct Node *slow = head;
	struct Node *fast = head;

	while (fast->next!=NULL)
	{
		slow = slow->next;
		fast = fast->next->next;
		if (fast==NULL)
			break;
	}

	cout<<"the middle element in the list is " <<slow->data<<endl;
}


struct Node* createLinkedList()
{
	struct Node *head = NULL;
	bool done = false;

	cout<<"Please continue entering the data to the linkedlist and type done to stop it\n";
	
	while(!done)
	{
		string input;	
		cin >> input;

		if (input=="done")
			break;

		int data = stoi(input);

		struct Node *temp = (struct Node*) malloc (sizeof(struct Node));
		temp->data = data;
		temp->next = NULL;

		if(head==NULL)
		{
			head= temp;
		}
		else
		{
			struct Node* itr = head;
			
			while(itr->next!=NULL)
			{
				itr = itr->next;
			}

			itr->next = temp;
		}

		cout<<"New node "<<data<<" is added to the linked list and type done to stop it"<<endl;

	}

	// printLinkedList(head);
	// add_at_pos(head);
	// del_at_pos(head);
	// del_match(head);
	// reverse(head);
	// remove_duplicates(head);
	return head;
	// print_mid(head);

}

void find_loop(struct Node *head)
{
	struct Node *slow = head;
	struct Node *fast = head->next;

	while(fast->next!=NULL)
	{
		slow = slow->next;
		fast = fast->next->next;

		if (fast==slow)
		{
			cout<<"The loop is found at "<<fast->data<<endl;
			break;
		}
	}
}
void create_loop(struct Node *head)
{
	struct Node *loop_node = NULL;
	struct Node *temp = head;
	cout<<"Please enter the index where you wanna create a loop"<<endl;

	int index;
	int curr_pos =0;
	cin>>index;

	while(temp->next!=NULL)
	{
		if (curr_pos==index)
		{
			loop_node = temp;
		}
		temp = temp->next;
		curr_pos+=1;
	}

	if (curr_pos==index)
		loop_node = temp;

	temp->next = loop_node;
	cout<<"The loop node is "<<loop_node->data<<endl;
	find_loop(head);
}

int main()
{
	struct Node *head = createLinkedList();
	printLinkedList(head);
	create_loop(head);

	return 0;
}