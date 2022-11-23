<template>
    <form class="sales-search-bar" @submit.prevent="emitSubmitSearch" @reset="emitResetList">
        <div class="sales-search-bar-inputs">
            <el-date-picker 
                v-model="refs.start" 
                placeholder="Промежуток (начало)" 
                size="large" 
                type="datetime"
                :class="{ error: startValid }" 
            />
            <el-date-picker 
                v-model="refs.end" 
                placeholder="Промежуток (конец)" 
                size="large" 
                type="datetime"
                :class="{ error: endValid }" 
            />
            <el-input 
                v-model="refs.name" 
                placeholder="Название продукта" 
                ssize="large" 
                type="datetime" 
            />
            <el-select
                v-model="refs.stock" 
                placeholder="ID склада" 
                size="large"
                @focus="emitGetStockOptions" 
                clearable
            >
                <el-option
                    v-for="stock in props.stockOptions"
                    :key="stock.id"
                    :label="`${stock.name} на ${stock.location}`"
                    :value="stock.id"
                    size="large"
                />
            </el-select>
            <el-button 
                size="large" 
                type="primary" 
                :icon="Search" 
                native-type="submit" 
            />
            <el-button 
                size="large" 
                type="warning" 
                :icon="RefreshRight" 
                native-type="reset"
                style="padding: 0; margin: 0;"
            />
        </div>
    </form>
</template>

<script setup lang="ts">
// imports
import { Search, RefreshRight } from '@element-plus/icons-vue';
import { defineEmits, defineProps, ref, computed } from 'vue';

import useVuelidate from '@vuelidate/core'
import { required } from '@vuelidate/validators'
////////////////

// props
const props = defineProps({
    stockOptions: {
        type: Array<any>,
        default: []
    }
})

// refs
const refs = ref({
    start: null,
    end: null,
    stock: null,
    name: null
})
///////////////

// validators
const validators = {
    start: { required },
    end: { required },
}
const v$ = useVuelidate(validators, refs)
/////////////

// computed
const startValid = computed(() => v$.value.start.$dirty && v$.value.start.$invalid)
const endValid = computed(() => v$.value.end.$dirty && v$.value.end.$invalid)
///////////////////

// emits
const emits = defineEmits(['reset-list', 'submit-search', 'get-stock-options'])

const emitSubmitSearch = () => {
    v$.value.$validate()
    if (!v$.value.$invalid) {
        emits('submit-search', refs.value, true)
    }
}

const emitResetList = () => {
    (Object.keys(refs.value) as (keyof typeof refs.value)[]).forEach((key) => {
        refs.value[key] = null
    })
    v$.value.$reset()
    emits('reset-list', refs.value, false)
}

const emitGetStockOptions = () => {
    emits('get-stock-options')
}
///////////////
</script>

<style scoped>
.sales-search-bar-inputs {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 0.2fr 0.2fr;
    gap: 10px;
}

.sales-search-bar-buttons {
    margin-left: auto;
}
</style>