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

void createLinkedList()
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

	printLinkedList(head);
	
}



int main()
{
	createLinkedList();
	return 0;
}