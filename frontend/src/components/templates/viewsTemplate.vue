<template>
    <div class="views-template-container">
        <column-group>
            <div class="views-template-top template-top">
                <div class="template-top-title">
                    <h2>{{props.topTitle}}</h2>
                </div>
                <div class="templete-top-slot">
                    <slot name="top-right"/>
                </div>
            </div>
            <div class="views-template-bar" v-if="hideBar">
                <slot name="bar"/>
            </div>
            <div class="views-template-body">
                <el-scrollbar :height="scrollbarHeigth">
                    <slot name="body"/>
                </el-scrollbar>
            </div>
            <div class="views-template-bottom" v-if="!fullBody">
                <slot name="bottom"/>
            </div>
        </column-group>
    </div>
</template>
  
<script setup lang="ts">
// imports
import columnGroup from './columnGroup.vue'
import { defineProps } from 'vue';
import { computed } from '@vue/reactivity';
//////////////////

// props
const props = defineProps({
    noBar: {
        type: Boolean,
        default: false
    },
    fullBody: {
        type: Boolean,
        default: false
    },
    topTitle: {
        type: String,
        default: ""
    }
})
//////////////////

// computed
const scrollbarHeigth = computed(()=> {
    switch (props.fullBody) {
        case false:
            return !props.noBar ? "calc(100vh - 340px)" : "calc(100vh - 240px)"
        case true:
            return "calc(100vh - 160px)"
    }
})
const hideBar = computed(()=> !props.noBar && !props.fullBody)
//////////////////
</script>

<style scoped>
    .views-template-container {
        width: 100%;
        height: 100%;
        color: var(--default-color);
        font-size: 18px;
    }
    
    .views-template-top {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 1rem 0;
    }
    .views-template-bar {
        padding: 2rem 0;
    }
    
</style>