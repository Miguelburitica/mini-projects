
class MyNode {
  value: any
  next: any
  
  constructor(value: any) {
    this.value = value
    this.next = null
  }
}

class Stack {
  top: any
  bottom: any
  length: number

  constructor() {
    this.top = null
    this.bottom = null
    this.length = 0
  }

  peek() {
    return this.top
  }

  push(value: any) {
    const newNode = new MyNode(value)
    const holdingPointer = this.top

    this.top = newNode

    if (this.length === 0) this.bottom = newNode
    else this.top.next = holdingPointer

    this.length++
    return this
  }
}