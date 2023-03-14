<template>
<div class="resume-container">
  <h2>Resumen de cuentas</h2>
  <div class="concept-item" :key="item.concept" v-for="item of conceptItems.list">{{ item.amount }}</div>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { AccountItemList } from '@/entities'

export default defineComponent({
  name: 'ResumeComponent',
  props: {
    conceptItems: {
      type: AccountItemList,
      default: { list: [] }
    }
  }
})

</script>
