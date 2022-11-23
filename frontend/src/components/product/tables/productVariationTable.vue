<template>
    <div class="product-variation-info">
        <el-table :data="props.variationList">
            <el-table-column type="expand">
                <template #default="props">
                    <el-table :data="props.row.stock_availability">
                        <el-table-column 
                            label="Название склада" 
                            prop="name"
                        />
                        <el-table-column 
                            label="Адрес склада" 
                            prop="location"
                        />
                        <el-table-column 
                            label="Кол-во" 
                            prop="amount" 
                        />
                    </el-table>
                </template>
            </el-table-column>
            <el-table-column 
                label="ID" 
                prop="id" 
            />
            <el-table-column label="Цена">
                <template #default="scope">
                    {{ noDataHandler(scope.row.price) }}
                </template>
            </el-table-column>
            <el-table-column 
                label="Форма выпуска" 
                prop="type" 
            />
            <el-table-column 
                label="Ед. изм" 
                prop="unit" 
            />
            <el-table-column fixed="right" width="180">
                <template #default="scope">
                    <el-button 
                        link 
                        :icon="Plus" 
                        type="primary"
                        @click="emits('price-event', scope.row.id)"
                    >
                        Цена
                    </el-button>
                    <el-button 
                        link 
                        :icon="Plus" 
                        type="primary"
                        @click="emits('stock-event', scope.row.id)"
                    >
                        В склад
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup lang="ts">
// imports
import { Plus } from '@element-plus/icons-vue';
import { defineProps, defineEmits } from 'vue';
import { noDataHandler } from '../../../helpers'
//////////////////

// props 
const props = defineProps({
    variationList: {
        type: Array,
        default: []
    }
})
//////////////////

// emts
const emits = defineEmits(['price-event', 'stock-event'])
//////////////////
</script>