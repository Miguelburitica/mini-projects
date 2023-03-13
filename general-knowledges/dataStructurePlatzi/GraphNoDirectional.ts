/* 
   2 - 0
  / \
 1 - 3


*/

// // Edge List

// const graph = [
//   [0, 2],
//   [2, 3],
//   [2, 1],
//   [1, 3]
// ]

// // Adjacent List

// const graphh = [[2], [2, 3], [0, 1, 3], [1, 2]]

// const graphhh = {
//   '0': [2],
//   '1': [2, 3],
//   '2': [0, 1, 3],
//   '3': [1, 2]
// }

class Graph {
  nodes: number
  adjacentList: Object

  constructor() {
    this.nodes = 0
    this.adjacentList = {}
  }

  addVertex(node: any) {
    this.adjacentList[node] = []
    this.nodes++
  }

  addEdge(node1: any, node2: any) {
    this.adjacentList[node1].push(node2)
    this.adjacentList[node2].push(node1)
  }
}
