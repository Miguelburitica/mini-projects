import { v4 } from 'uuid'
export class AccountItem {
  constructor (concept: string, amount: number, type: string) {
    this.concept = concept
    this.amount = amount
    this.type = type
    this.id = v4()
  }

  concept: string
  amount: number
  type: string
  id: string
}

export class AccountItemList {
  constructor () {
    this.list = []
  }

  list: AccountItem[]
}
