<template>
    <div class="product-table">
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
                    width="300"
                />
                <el-table-column 
                    prop="tags" 
                    label="Теги"  
                    width="300"
                />
                <el-table-column 
                    prop="description" 
                    label="Описание" 
                    width="400"
                />
                <el-table-column fixed="right" width="140">
                    <template #default="scope">
                        <el-button 
                            link 
                            type="primary"
                            @click="emitRowId(scope.row.id)"
                        >
                            Подробности
                        </el-button>
                    </template>
                </el-table-column>
        </el-table>
    </div>
</template>

<script setup lang="ts">
// imports
import { defineEmits, defineProps } from 'vue';
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

// emits
const emits = defineEmits(['row-clicked'])

const emitRowId = (id: number) => {
    emits('row-clicked', id)
}
/////////////////
</script>