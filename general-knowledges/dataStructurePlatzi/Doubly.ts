
// let singlyLinkedList = {
//   head: {
//     value: 1,
//     next: {
//       value: 2,
//       next: {
//         value: 3,
//         next: {
//           value: 4,
//           next: null
//         }
//       }
//     }
//   }
// }

  class MyNode {
    value: String|number|Array<any>
    next: any
    constructor(value: String|number|Array<any>) {
      this.value = value
      this.next = null
    }
  }

  // Singly linked list
  class MyLinkList {
    public head: MyNode
    public tail: MyNode
    public length: number
    constructor(value: String|number|Array<any>) {
      this.head = new MyNode(value)
      this.tail = this.head
      this.length = 1;
    }

    getTheIndex(index: number) {
      let cont = 0
      let currentNode = this.head

      while(cont < index) {
        currentNode = currentNode.next
        cont++
      }

      return currentNode
    }

    append(value: String|number|Array<any>) {
      const newNode = new MyNode(value)
      this.tail.next = newNode
      this.tail = newNode
      this.length++
      return this
    }

    preppend(value: String|number|Array<any>) {
      const newNode = new MyNode(value)
      newNode.next = this.head
      this.head = newNode
      this.length++
      return this
    }

    insert(index: number, value: String|number|Array<any>) {
      if(this.length < index) return this.append(value)

      const newNode = new MyNode(value)
      const backwardPointer = this.getTheIndex(index - 1)
      const holdingPointer = backwardPointer.next
      backwardPointer.next = newNode
      newNode.next = holdingPointer

      this.length++

      return this
    }
    
    remove(index: number, value: String|number|Array<any>) {
      if(this.length < index) return false

      const leftPointer = this.getTheIndex(index - 1)
      const rigthPointer = this.getTheIndex(index + 1)
      leftPointer.next = rigthPointer

      this.length--

      return this
    }
  }

let myLinkedList = new MyLinkList(1)