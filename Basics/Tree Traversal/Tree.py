class TreeNode:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

class BinaryTree:
    def __init__(self):
        self.root = None

    def insert(self, value):
        self.root = self._insert(self.root, value)

    def _insert(self, root, value):
        if root is None:
            return TreeNode(value)

        if value < root.value:
            root.left = self._insert(root.left, value)
        else:
            root.right = self._insert(root.right, value)

        return root

    def in_order_traversal(self):
        result = []
        self._in_order_traversal(self.root, result)
        return result

    def _in_order_traversal(self, node, result):
        if node is not None:
            self._in_order_traversal(node.left, result)
            result.append(node.value)
            self._in_order_traversal(node.right, result)

    def pre_order_traversal(self):
        result = []
        self._pre_order_traversal(self.root, result)
        return result

    def _pre_order_traversal(self, node, result):
        if node is not None:
            result.append(node.value)
            self._pre_order_traversal(node.left, result)
            self._pre_order_traversal(node.right, result)

    def post_order_traversal(self):
        result = []
        self._post_order_traversal(self.root, result)
        return result

    def _post_order_traversal(self, node, result):
        if node is not None:
            self._post_order_traversal(node.left, result)
            self._post_order_traversal(node.right, result)
            result.append(node.value)

# Example usage:
my_tree = BinaryTree()
my_tree.insert(5)
my_tree.insert(3)
my_tree.insert(7)
my_tree.insert(2)
my_tree.insert(4)
my_tree.insert(6)
my_tree.insert(8)

print("In-order traversal:", my_tree.in_order_traversal())
print("Pre-order traversal:", my_tree.pre_order_traversal())
print("Post-order traversal:", my_tree.post_order_traversal())
