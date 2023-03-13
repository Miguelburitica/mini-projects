class MyNode {
  value: any
  left: any
  right: any
  
  constructor(value: any) {
    this.value = value
    this.left = null
    this.right = null
  }
}

class BinarySearchTree {
  root: any

  constructor() {
    this.root = null
  }

  insert(value: number) {
    const newNode = new MyNode(value)

    if (!this.root) this.root = newNode
    else {
      let currentNode = this.root

      while(true) {
        if(value === currentNode.value) return this

        if (newNode.value < currentNode.value) {
          if (!currentNode.left) {
            currentNode.left = newNode
            return this
          }
          
          currentNode = currentNode.left
        } else {
          if (!currentNode.right) {
            currentNode.right = newNode 
            return this
          }
          
          currentNode = currentNode.right
        }
      }
    }

    return this
  }

  get(value: number) {
    if (!this.root) return false
    let currentNode = this.root

    while(true) {
      if (value < currentNode.value) {
        if (currentNode.left === undefined) return false

        if (value === currentNode.left.value) return currentNode.left
        currentNode = currentNode.left
      } else {
        if (currentNode.right === undefined) return false
        
        if (value === currentNode.right.value) return currentNode.right
        currentNode = currentNode.right
      }
    }
  }
}

const tree = new BinarySearchTree()

console.log(tree)

tree.insert(10)

tree.insert(4)
tree.insert(20)
tree.insert(2)
tree.insert(8)
tree.insert(17)
tree.insert(170)

console.log(tree)

console.log(tree.get(20))
