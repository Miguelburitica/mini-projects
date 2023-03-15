<template>
<div class="resume-container">
  <h2>Resumen de cuentas</h2>
  <div class="items-container">
    <div class="concept-item" :key="accountItem.concept+accountItem.type" v-for="accountItem of conceptItems.list">
      <h6>{{ accountItem.concept }}</h6>
      <caption>{{ formatCurrency(accountItem.amount) }}</caption>
      <i class="material-icons-round" @click="$emit('removeItem', accountItem.id)">delete</i>
    </div>
  </div>

  <div class="flex-column-center">
    <b class="resume">Total: {{ totalResume.subtotal }}</b>
    <b class="resume">De un total de {{ totalResume.totalItems }} conceptos</b>
    <button class="confirm" @click="exportExcel">Exportar consolidado</button>
  </div>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { AccountItemList } from '@/entities'
import * as XLSX from 'xlsx'

export default defineComponent({
  name: 'ResumeComponent',
  props: {
    conceptItems: {
      type: AccountItemList,
      default: { list: [] }
    }
  },
  computed: {
    totalResume () {
      let total = Number()
      for (const item of this.conceptItems.list) {
        total += Number(item.amount)
      }

      return { subtotal: this.formatCurrency(total), totalItems: this.conceptItems.list.length }
    }
  },
  methods: {
    formatCurrency (amount: number): string {
      const formater = Intl.NumberFormat('es-CO', { style: 'currency', currency: 'COP' })

      return formater.format(amount)
    },
    exportExcel () {
      // Incomes Table
      let incomeTotal = 0
      const incomes = this.conceptItems.list
        .filter(item => item.type === 'income')
        .map(item => {
          const { concept, amount } = item
          incomeTotal += Number(amount)
          return { Concepto: concept, Monto: amount }
        })

      const incomesSheet = [
        ['Ingresos'],
        Object.keys(incomes[0]),
        ...incomes.map(item => Object.values(item)),
        ['Subtotal:', incomeTotal]
      ]

      // Expenses table
      let expenseTotal = 0
      const expenses = this.conceptItems.list
        .filter(item => item.type === 'expense')
        .map(item => {
          const { concept, amount } = item
          expenseTotal += Number(amount)
          return { Concepto: concept, Monto: amount }
        })

      const expensesSheet = [
        ['Gastos'],
        Object.keys(expenses[0]),
        ...expenses.map(item => Object.values(item)),
        ['Subtotal:', expenseTotal],
        ['Total:', expenseTotal + incomeTotal]
      ]
      const worksheet = XLSX.utils.aoa_to_sheet([])
      XLSX.utils.sheet_add_aoa(worksheet, expensesSheet, { origin: 'A1' })
      XLSX.utils.sheet_add_aoa(worksheet, incomesSheet, { origin: 'E1' })

      const workbook = XLSX.utils.book_new()
      XLSX.utils.book_append_sheet(workbook, worksheet, 'Data')

      XLSX.writeFile(workbook, 'file.xlsx')
    }
  }
})

</script>

<style lang="less" scoped>

.resume-container {
  border: var(--border-color) 1px solid;
  box-shadow: 0px 10px 15px 0px #2c3e50;
  -webkit-box-shadow: 0px 10px 15px 0px #2c3e50;
  -moz-box-shadow: 0px 10px 15px 0px #2c3e50;
  padding: 20px 10px;
  border-radius: 6px;
  width: 100%;

  h2 {
    margin: 0;
    color: var(--blue-color);
    margin-bottom: 10px;
  }
}

.items-container {
  display: flex;
  flex-direction: column;
  gap: 6px;
  @media(min-width: 800px) {
    max-height: 75vh;
    overflow: auto;
  }
}

.concept-item {
  display: flex;
  width: 100%;
  align-items: center;
  border-radius: 6px;
  padding: 4px 0;
  background-color: ghostwhite;
  transition: .3s ease all;

  &:hover {
    background-color: darken(ghostwhite, 20%);
  }

  h6 {
    width: 40%;
    margin: 0;
  }

  caption {
    width: 50%;
  }

  i {
    background-color: transparent;
    color: var(--wrong-color);
    border-radius: 8px;
    padding: 4px 4px;
    font-size: 20px;
    cursor: pointer;
    transition: .3S ease all;

    &:hover {
      background-color: var(--wrong-color);
      color: ghostwhite;
    }
  }
}

.flex-column-center {
  border-top: var(--border-color) 1px solid;
  padding: 15px 0 0;
}

.resume {
  margin-bottom: 10px;
  color: var(--blue-color);
}

button.confirm {
  background-color: var(--blue-color);
  color: var(--dark-color);
  outline: none;
  border: none;
  border-radius: 6px;
  padding: 4px 30px;
  transition: .3s ease all;
  cursor: pointer;

  &:hover {
    filter: brightness(110%);
  }
}

</style>
