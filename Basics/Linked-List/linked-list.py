class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None

    def insert(self, data):
        new_node = Node(data)
        new_node.next = self.head
        self.head = new_node

    def delete(self, key):
        current = self.head

        if current is not None and current.data == key:
            self.head = current.next
            return

        prev = None
        while current is not None and current.data != key:
            prev = current
            current = current.next

        if current is None:
            return

        prev.next = current.next

# Example usage:
# ll = LinkedList()
# ll.insert(3)
# ll.insert(7)
# ll.delete(3)
