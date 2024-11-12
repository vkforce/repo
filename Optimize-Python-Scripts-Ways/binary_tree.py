class Node:
    def __init__(self, data):
        self.left = None
        self.right = None
        self.data = data

    def print_tree(self):
        if self.left:
            self.left.print_tree()
        print(self.data, end='\n')
        if self.right:
            self.right.print_tree()

# Creating the root node
root = Node(10)
root.left = Node(5)
root.right = Node(15)

# Inserting more nodes
root.left.left = Node(3)
root.left.right = Node(7)
root.right.left = Node(12)
root.right.right = Node(18)

# Printing the tree
root.print_tree()