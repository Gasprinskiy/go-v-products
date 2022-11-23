<template>
    <div class="pagination-bar">
        <div class="pagination-count">
            <el-pagination 
                v-model:currentPage="paginationParams.offset"
                layout="prev, pager, next" 
                :total="paginationTotal"
                background
                @current-change="emitOffsetChange"
            />
        </div>
        <div class="pagination-offset-select">
            <el-select
                v-model="paginationParams.limit"
                placeholder="количество"
                size="small"
                @change="emitLimitChange"
            >
                <el-option
                    v-for="option in props.limitOptions"
                    :key="option"
                    :label="option"
                    :value="option"
                />
            </el-select>
        </div>
    </div>
</template>

<script setup lang="ts">
// imports
import { defineProps, defineEmits } from 'vue';
import { computed, ref } from '@vue/reactivity';
/////////////////

// props
const props = defineProps({
    listTotalPage: {
        type: Number,
        default: 0
    },
    paginationParams: {
        type: Object,
        default: {}
    },
    limitOptions: {
        type: Array,
        default: []
    },
})
//////////////////

// refs

// computed
const paginationTotal = computed(() => props.listTotalPage * 10)
//////////////////

// emits 
const emits = defineEmits(['limit-change', 'offset-change'])

const emitOffsetChange = (e: any) => {
    emits('offset-change', e)
}

const emitLimitChange = (e: any) => {
    emits('limit-change', e)
} 
//////////////////
</script>

<style scoped>
    .pagination-bar {
        width: 100%;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
</style>