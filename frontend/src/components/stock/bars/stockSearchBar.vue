<template>
    <form 
        class="product-search-bar"
        @submit.prevent="emits('submit-search')"
    >
        <el-select 
            placeholder="Поиск продукта" 
            clearable 
            size="large"
            :prefix-icon="Search"
            @change="emitQueryChanges"
            @clear="emits('reset-list')"
            v-model="refs.query"
            filterable
            remote
            :remote-method="emitProductSearch"
        >
            <el-option
                v-for="product in props.productOptions"
                :key="product.id"
                :label="product.name"
                :value="product.id"
            />
        </el-select>
        <el-button
            type="primary"
            native-type="submit"
            :icon="Search"
            size="large"
            :disabled="!queryValid"
        />
    </form>
</template>

<script setup lang="ts">
// imports
import { Search } from '@element-plus/icons-vue';
import { defineEmits, defineProps, ref, computed } from 'vue';

import useVuelidate from '@vuelidate/core'
import { required, minLength } from '@vuelidate/validators'
////////////////

// props
const props = defineProps({
    productOptions: {
        type: Array<any>,
        default: []
    }
})

// refs
const refs = ref({
    query: ""
})
///////////////

// validators
const validators = {
    query: {required}
}
const v$ = useVuelidate(validators, refs)
/////////////

// computed
const queryValid = computed(() => !v$.value.query.$dirty && !v$.value.query.$invalid)
///////////////////

// emits
const emits = defineEmits(['reset-list', 'submit-search', 'update-query', 'search-product'])


const emitQueryChanges = (val: string) => {
    emits('update-query', val)
}

const emitProductSearch = (val: string) => {
    if (val){
        emits('search-product', val)
    }
}
///////////////
</script>

<style scoped>
     .product-search-bar {
        display: grid;
        grid-template-columns: calc(93% - 5px) 7%;
        gap: 5px;
    }
</style>