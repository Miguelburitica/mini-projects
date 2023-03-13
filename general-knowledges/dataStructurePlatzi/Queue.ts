class MyNode {
  value: any
  next: any
  
  constructor(value: any) {
    this.value = value
    this.next = null
  }
}

class Queue {
  first: any
  last: any
  length: number

  constructor() {
    this.first = null
    this.last = null
    this.length = 0
  }

  peek() {
    return this.first
  }

  enqueue(value: any) {
    const newNode = new MyNode(value)

    
    if (this.length === 0) this.first = newNode
    else this.last.next = newNode
    
    this.last = newNode

    this.length++
    return this
  }

  dequeue() {
    const newFirst = this.first.next

    this.first = newFirst
    this.length--

    return this
  }
}

const myQueue = new Queue()

myQueue.enqueue(1)
myQueue.enqueue(2)
myQueue.enqueue(3)