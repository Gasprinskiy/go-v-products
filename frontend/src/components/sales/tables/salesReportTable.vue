<template>
    <div class="sales-report-table">
        <el-table :data="props.list" v-loading="loadingflag">
            <el-table-column 
                label="ID" 
                prop="id"
                fixed="left" 
            />
            <el-table-column 
                label="Название" 
                prop="name"
                width="300" 
            />
            <el-table-column 
                label="Форма выпуска" 
                prop="type" 
            />
            <el-table-column 
                label="Ед. изм" 
                prop="unit" 
            />
            <el-table-column 
                label="Кол-во" 
                prop="amount" 
            />
            <el-table-column 
                label="Дата продажи"
                width="300" 
            >
                <template #default="scope">
                    {{ parseDate(scope.row.soldDate) }}
                </template>
            </el-table-column>
            <el-table-column 
                label="Склад" 
                prop="stock" 
            />
            <el-table-column 
                fixed="right" 
            >
                <template #default="scope">
                    <el-button
                        link
                        type="primary"
                        @click="emitRowDetails(scope.row)"
                    >
                        К продукту
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup lang="ts">
// imports
import { defineProps, defineEmits } from 'vue';
import { parseDate } from '../../../helpers'
//////////////////

// props 
const props = defineProps({
    list: {
        type: Array,
        default: []
    },
    loadingflag: {
        type: Boolean,
        default: false
    }
})
//////////////////

// emits
const emits = defineEmits(['get-details'])

const emitRowDetails = (row: any) => {
    emits('get-details', row)
}
//////////////////
</script>