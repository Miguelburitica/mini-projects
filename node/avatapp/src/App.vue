<template>
<div class="container">
  <ResumeComponent :conceptItems="currentConceptsList"></ResumeComponent>
  <div class="concept-form">
    <div class="mode-concept">
      <b :class="{ active: expenseMode }" @click="expenseMode = true">Gasto</b>
      <b :class="{ active: !expenseMode }" @click="expenseMode = false">Ingreso</b>
    </div>
    <div class="concept-inputs">
      <input type="text" placeholder="Ingresar concepto">
      <div>
        <input type="tel" placeholder="Ingresar valor">
        <button>Add</button>
      </div>
    </div>
  </div>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ResumeComponent from '@/components/ResumeComponent.vue'
import { AccountItemList } from './entities'

export default defineComponent({
  name: 'App',
  components: {
    ResumeComponent
  },
  data: function () {
    return {
      currentConceptsList: AccountItemList,
      expenseMode: Boolean()
    }
  }
})
</script>

<style lang="less">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.container {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 100%;
  max-width: 1120px;
  margin: 0 auto;
}

.concept-form {
  display: flex;
  flex-direction: column;
  width: 60%;
  align-items: center;

  .mode-concept {
    width: 60%;
    display: flex;
    justify-content: space-evenly;
    border-radius: 6px;
    overflow: hidden;
    margin-bottom: 30px;

    b {
      background-color: #2c3e5033;
      padding: 8px 0;
      width: 50%;
      transition: all .3s ease;
      cursor: pointer;

      &.active {
        background-color: #2c3e50dd;
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
    width: 60%;

    input {
      width: 100%;
      border-radius: 4px;
      outline: none;
      border: 1px solid #2c3e50;
    }

    div {
      display: flex;
      width: 100%;
    }
  }
}
</style>
