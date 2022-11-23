<template>
    <div 
        class="column-group-container"
        :style="gap"
    >
        <slot/>
    </div>
</template>
  
<script setup lang="ts">
// imports
import { defineProps, computed } from 'vue';
//////////////////

// props
const props = defineProps({
    gap: {
        type: Number,
        default: 5
    }
})
//////////////////

// computed
const gap = computed(()=> `gap:${props.gap}px;`)
//////////////////
</script>

<style scoped>
    .column-group-container {
        width: 100%;
        display: grid;
        grid-template-rows: auto;
        grid-template-columns: 100%;
    }
</style>
  