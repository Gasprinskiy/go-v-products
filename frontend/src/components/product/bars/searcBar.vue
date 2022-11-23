<template>
    <form 
        class="product-search-bar"
        @submit.prevent="emitSubmitSearch"
    >
        <el-input 
            :placeholder="props.placeholder" 
            clearable 
            size="large"
            :prefix-icon="Search"
            @clear="emitResetList"
            @blur="emitRestIfValueClear"
            v-model="refs.query"
            @input="emitQueryChanges"
        />
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
import { useRoute } from 'vue-router'; 

import useVuelidate from '@vuelidate/core'
import { required } from '@vuelidate/validators'
////////////////

// props
const props = defineProps({
    placeholder: {
        type: String,
        default: "Поиск"
    }
})
///////////////////

// Vue use
const route = useRoute()
///////////////////
// computed
const routeQuerry = computed(() => route.params.querry ? route.params.querry : "")
///////////////////

// refs
const refs = ref({
    query: routeQuerry.value
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
const emits = defineEmits(['reset-list', 'submit-search', 'update-query'])

const emitSubmitSearch = () => {
    emits('submit-search')
}

const emitResetList = () => {    
    emits('reset-list')
}

const emitRestIfValueClear = () => {
    if (!refs.value.query) {
        emits('reset-list')
    }
}

const emitQueryChanges = (val: string | string[]) => {
    emits('update-query', val)
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