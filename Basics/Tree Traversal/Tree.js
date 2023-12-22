class TreeNode {
    constructor(value) {
      this.value = value;
      this.left = null;
      this.right = null;
    }
  }
  
  class BinaryTree {
    constructor() {
      this.root = null;
    }
  
    insert(value) {
      this.root = this._insert(this.root, value);
    }
  
    _insert(root, value) {
      if (root === null) {
        return new TreeNode(value);
      }
  
      if (value < root.value) {
        root.left = this._insert(root.left, value);
      } else {
        root.right = this._insert(root.right, value);
      }
  
      return root;
    }
  
    inOrderTraversal() {
      const result = [];
      this._inOrderTraversal(this.root, result);
      return result;
    }
  
    _inOrderTraversal(node, result) {
      if (node !== null) {
        this._inOrderTraversal(node.left, result);
        result.push(node.value);
        this._inOrderTraversal(node.right, result);
      }
    }
  
    preOrderTraversal() {
      const result = [];
      this._preOrderTraversal(this.root, result);
      return result;
    }
  
    _preOrderTraversal(node, result) {
      if (node !== null) {
        result.push(node.value);
        this._preOrderTraversal(node.left, result);
        this._preOrderTraversal(node.right, result);
      }
    }
  
    postOrderTraversal() {
      const result = [];
      this._postOrderTraversal(this.root, result);
      return result;
    }
  
    _postOrderTraversal(node, result) {
      if (node !== null) {
        this._postOrderTraversal(node.left, result);
        this._postOrderTraversal(node.right, result);
        result.push(node.value);
      }
    }
  }
  
  // Example usage:
  const myTree = new BinaryTree();
  myTree.insert(5);
  myTree.insert(3);
  myTree.insert(7);
  myTree.insert(2);
  myTree.insert(4);
  myTree.insert(6);
  myTree.insert(8);
  
  console.log("In-order traversal:", myTree.inOrderTraversal());
  console.log("Pre-order traversal:", myTree.preOrderTraversal());
  console.log("Post-order traversal:", myTree.postOrderTraversal());
  