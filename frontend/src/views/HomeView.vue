<template>
    <div class="home-view-container router-inner">
        <column-group :gap="25">
            <div class="home-view-logo">
                <img src="../assets/logo.svg">
            </div>
            <div class="home-view-router-button-group">
                <column-group :gap="15">
                    <el-button
                        class="home-view-router-button"
                        v-for="route in mainRoutes"
                        @click="router.push(route.path)"
                        :icon="route.icon"
                        style="width: 100%; padding: 3rem; margin: 0;"
                    >
                        {{route.title}}
                    </el-button>
                </column-group>
            </div>
        </column-group>
    </div>
</template>
  
<script setup lang="ts">
// imports
import columnGroup from '../components/templates/columnGroup.vue';
import { useRouter } from 'vue-router';
import { mainRoutes } from '../composables/main-routes';
/////////////////

// Vue use
const router = useRouter()
/////////////////
</script>
  
<style scoped>
    .home-view-container {
        height: 100%;
        width: 100%;
    }

    .home-view-logo {
        text-align: center;
        padding: 2rem 0;
    }
    .home-view-logo img {
        width: 150px;
    }
    
    .home-view-router-button-group {
        width: 100%;    
        /* padding: 0 9rem; */
    }

    .home-view-router-button {
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        outline: 1px solid var(--bg-color-secondary);
        border-radius: var(--br-radius);
        color: var(--default-color);
        font-size: 18px !important;
    }

    .home-view-router-button:hover {
        box-shadow: var(--elements-shadow);
        transition: var(--transtion-duration) var(--transtion-type);
        color: var(--color-primary);
    }
</style>
  