<template>
    <div class="product-info">
        <el-table :data="[props.info]">
            <el-table-column
                label="Название"
                prop="name"
            />
            <el-table-column
                label="Описание"
                prop="description"
            />
            <el-table-column
                label="Теги"
                prop="tags"
            />
        </el-table>
    </div>
</template>

<script setup lang="ts">
// imports
import { defineProps } from 'vue';
//////////////

// props
const props = defineProps({
    info: {
        type: Object,
        default: {}
    }
})
//////////////
</script>