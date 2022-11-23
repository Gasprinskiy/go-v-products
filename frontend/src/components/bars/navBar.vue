<template>
  <div class="nav-bar-container">
    <div class="nav-bar-item nav-bar-left">
        <el-button 
            type="primary" 
            :icon="ArrowLeft"
            @click="router.back()"
            title="Назад"
        />
        <router-link
            to="/"
        >
            <el-button 
                type="primary" 
                :icon="HomeFilled"
                title="Главная"
            />
        </router-link>
    </div>
    <div class="nav-bar-item nav-bar-righ">
        <router-link
            v-for="route in mainRoutes"
            :to="route.path"
            :class="{'router-link-active': routeActive(route.name)}"
        >
            <el-button 
                type="primary" 
                :icon="route.icon"
                :title="route.title"
            />
        </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
// imports

import { 
    ArrowLeft,
    HomeFilled
} from '@element-plus/icons-vue'
import { useRouter, useRoute } from 'vue-router';
import { RouterLink } from 'vue-router'; 
import { mainRoutes } from '../../composables/main-routes';
/////////////////

// Vue use
const router = useRouter()
const route = useRoute()
/////////////////

// methods
const routeActive = (name: string) => {
    return route.name === name
}
/////////////////
</script>


<style scoped>
    .nav-bar-container {
        width: 100%;
        display: flex;
        justify-content: space-between;
        padding: 0.6rem;
        background-color: var(--bg-color-secondary);
    }

    .nav-bar-left {
        width: 100px;
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 5px;
    }

    .nav-bar-righ {
        width: 145px;
        display: grid;
        grid-template-columns: 1fr 1fr 1fr;
        gap: 5px;
    }
</style>
