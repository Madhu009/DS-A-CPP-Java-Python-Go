#include <iostream>

using namespace std;

struct Node
{
	int data;
	struct Node *next = NULL;
	struct Node *prev = NULL;
};


struct Node* create_dll()
{
	struct Node *head = NULL;
	bool done = false;

	cout<<"Please continue entering the data to the linkedlist and type done to stop it\n";

	while (!done){
		string input;
		cin>>input;
		
		if (input =="done")
			break;
		int data = stoi(input);

		struct Node *new_node = (struct Node*) malloc (sizeof(struct Node));
		new_node->data = data;
		new_node->next = NULL;
		new_node->prev = NULL;

		if (head==NULL)
			head = new_node;
		else
		{
			struct Node *temp = head;

			while(temp->next!=NULL)
				temp = temp->next;
			temp->next = new_node;
			new_node->prev = temp;	
		}

		cout<<"New node "<<data<<" is added to the linked list and type done to stop it"<<endl;

	}

	return head;

}

void printLinkedList(struct Node* head )
{
	struct Node* itr = head;
	struct Node* last = NULL;
	cout<<"-------------"<<endl;
	while(itr!=NULL)
	{
		last = itr;
		cout<<itr->data<<" ";
		itr = itr->next;
	}

	cout<<endl;

	while(last !=NULL)
	{
		cout<<last->data<<" ";
		last = last->prev;
	}

	cout<<endl;
	cout<<"-------------"<<endl;


}

struct Node* insert_pos(struct Node *head)
{
	int curr_pos = 1;
	struct Node *temp = head;
	cout<<"Please enter the data to be inserted"<<endl;
	string input;
	cin>>input;
	int data = stoi(input);

	cout<<"Please enter the index where you wanna insert the node"<<endl;
	cin>>input;

	int pos = stoi(input);

	struct Node *new_node = (struct Node*) malloc (sizeof(struct Node));
	new_node->next = NULL;
	new_node->prev = NULL;
	new_node->data = data;

	if (pos == 0)
	{
		new_node->next = head;
		head->prev = new_node;
		head = new_node;
		cout<<"The node has been inserted at "<<pos<<" index"<<endl;
		return head;
	}

	while(temp->next!=NULL)
	{
		if (curr_pos==pos)
		{
			new_node->next = temp->next;
			temp->next->prev = new_node;
			temp->next = new_node;
			new_node->prev = temp;

			cout<<"The node has been inserted at "<<pos<<" index"<<endl;
			return head;
		}
		temp = temp->next;
		curr_pos+=1;

		if (temp->next==NULL && curr_pos==pos)
		{
			temp->next = new_node;
			new_node->prev= temp;
			curr_pos+=1;
			cout<<"The node has been inserted at "<<pos<<" index"<<endl;
			return head;
		}
	}

	cout<<"index is out of range or invalid"<<endl;
	return head;
}

struct Node* delete_pos(struct Node *head)
{
	cout<<"Please enter the index where you wanna delete the node"<<endl;
	string input;
	cin>>input;

	int pos = stoi(input);
	int curr_pos=1;

	struct Node *temp = head;

	if (pos == 0)
	{
		head->next->prev = NULL;
		head = head->next;
		cout<<"The node at "<<pos<<" has been deleted"<<endl;
		return head;
	}

	while(temp->next!= NULL)
	{
		if (curr_pos==pos)
		{
			temp->next = temp->next->next;
			if (temp->next!=NULL)
				temp->next->prev = temp;
			cout<<"The node at "<<pos<<" has been deleted"<<endl;
			return head;
		}

		temp = temp->next;
		curr_pos+=1;

		if (temp->next==NULL && curr_pos==pos)
		{
			temp->prev->next = NULL;
			temp->prev= NULL;

			cout<<"The node at "<<pos<<" has been deleted"<<endl;
			return head;
		}
	}

	cout<<"index is out of range or invalid"<<endl;
	return head;

}

int main()
{
	struct Node *head = create_dll();
	printLinkedList(head);
	// struct Node *test = insert_pos(head);
	// printLinkedList(test);
	struct Node *test = delete_pos(head);
	printLinkedList(test);

	return 0;
}