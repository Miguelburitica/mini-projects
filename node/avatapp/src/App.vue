<template>
<div class="container">
  <div class="resume-container">
    <ResumeComponent :conceptItems="currentConceptsList" @removeItem="removeItem"></ResumeComponent>
  </div>
  <div class="concept-form">
    <div class="mode-concept">
      <b class="expenses" :class="{ active: expenseMode }" @click="expenseMode = true">Gasto</b>
      <b class="incomes" :class="{ active: !expenseMode }" @click="expenseMode = false">Ingreso</b>
    </div>
    <div class="concept-inputs">
      <input v-model="conceptName" type="text" placeholder="Ingresar concepto">
      <div>
        <input
          v-model="conceptAmountLabel"
          type="text"
          placeholder="Ingresar valor"
          @keypress.enter="createNewConceptItem"
          @input="realTimeFormat"
        >

        <button @click="createNewConceptItem">Add</button>
      </div>
    </div>
  </div>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ResumeComponent from '@/components/ResumeComponent.vue'
import { AccountItemList, AccountItem } from './entities'

export default defineComponent({
  name: 'App',
  components: {
    ResumeComponent
  },
  data: function () {
    return {
      currentConceptsList: new AccountItemList(),
      expenseMode: Boolean(),
      conceptName: String(),
      conceptAmount: Number(),
      conceptAmountLabel: String(),
      timeoutId: Number()
    }
  },
  methods: {
    createNewConceptItem () {
      if (!this.conceptName || !this.conceptAmount) return

      const typeConcept = this.expenseMode ? 'expense' : 'income'
      const conceptAmount = this.expenseMode ? this.conceptAmount * -1 : this.conceptAmount
      const conceptItem = new AccountItem(this.conceptName, conceptAmount, typeConcept)
      this.currentConceptsList.list.push(conceptItem)
    },
    removeItem (id: string) {
      const newConceptList = this.currentConceptsList.list.filter(item => item.id !== id)

      this.currentConceptsList.list = newConceptList
    },
    formatCurrency (amount: number): string {
      const formater = Intl.NumberFormat('es-CO', { style: 'currency', currency: 'COP' })

      return formater.format(amount)
    },
    realTimeFormat (event: Event) {
      clearTimeout(this.timeoutId)
      const input = event.target as HTMLInputElement

      const userInput = input.value
      const numericInput1 = userInput.replace(/[^\d-,]/g, '')
      const numericInput = numericInput1.replace(/[,]/g, '.')

      const number = parseFloat(numericInput)

      const formattedValue = this.formatCurrency(number)

      this.conceptAmount = number

      this.timeoutId = setTimeout(() => {
        this.conceptAmountLabel = formattedValue
      }, 1000)
    }
  }
})
</script>

<style lang="less">
* {
  box-sizing: border-box;
}

body {
  color-scheme: dark;
  background-color: var(--dark-color);
}

:root {
  --ok-color: #9FCC2E;
  --blue-color: #90D7FF;
  --border-color: #BFD0E0;
  --wrong-color: #E3170A;
  --dark-color: #2D1E2F;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 40px;
}

.flex-column-center {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.container {
  display: flex;
  justify-content: space-around;
  width: 100%;
  max-width: 1120px;
  margin: 0 auto;

  @media(max-width: 800px) {
    align-items: center;
    flex-direction: column-reverse;
  }
}

.resume-container {
  width: 40%;

  @media(max-width: 800px) {
    width: 95%;
  }
}

.concept-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 60%;

  .mode-concept {
    width: 80%;
    display: flex;
    justify-content: space-evenly;
    border-radius: 6px;
    overflow: hidden;
    margin-bottom: 30px;

    b {
      background-color: #bedbf8;
      color: var(--dark-color);
      padding: 8px 0;
      width: 50%;
      transition: all .3s ease;
      cursor: pointer;
      line-height: 1;

      &.incomes.active {
        background-color: var(--ok-color);
        color: ghostwhite;
      }

      &.expenses.active {
        background-color: var(--wrong-color);
        color: ghostwhite;
      }

      &:hover {
        background-color: #2c3e50aa;
        color: ghostwhite;
      }
    }
  }

  .concept-inputs {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 10px;
    width: 80%;

    input {
      width: 100%;
      border-radius: 4px;
      outline: none;
      border: var(--border-color) 1px solid;
      padding: 6px 10px;
    }

    div {
      display: flex;
      width: 100%;

      input {
        border-radius: 4px 0 0 4px;
        border-right: none;
      }

      button {
        background-color: var(--blue-color);
        color: var(--dark-color);
        border: var(--border-color) 1px solid;
        border-radius: 0 4px 4px 0;
        outline: none;
        transition: .3s all ease;
        cursor: pointer;

        &:hover {
          filter: brightness(108%);
        }
      }
    }
  }

  @media(max-width: 800px) {
    width: 100%;
    margin-bottom: 25px;
  }
}
</style>
