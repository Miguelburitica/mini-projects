export class AccountItem {
  constructor (concept: string, amount: number) {
    this.concept = concept
    this.amount = amount
  }

  concept: string
  amount: number
}

export class AccountItemList {
  constructor () {
    this.list = []
  }

  list: AccountItem[]
}
