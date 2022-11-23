<template>
    <div class="stock-table">
        <el-table
            height="100%"
            :data="props.list"
            v-loading="props.loadingflag"
        >
            <el-table-column 
                prop="id" 
                label="ID" 
                fixed="left"
            />
            <el-table-column 
                prop="name" 
                label="Название" 
            />
            <el-table-column 
                prop="location" 
                label="Адрес"
            />
            <el-table-column 
                label="Общее кол-во" 
            >
                <template #default="scope">
                    {{ noDataHandler(scope.row.amount) }}
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup lang="ts">
// imports
import { defineProps } from 'vue';
import { noDataHandler } from '../../../helpers';
/////////////////

// props
const props = defineProps({
    loadingflag: {
        type: Boolean,
        default: false
    },
    list: {
        type: Array,
        default: []
    }
})
/////////////////
</script>