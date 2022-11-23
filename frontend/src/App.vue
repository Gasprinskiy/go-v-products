<template>
  <nav-bar/>
  <div class="router-view-wrapper">
    <router-view/>
  </div>
</template>

<script setup lang="ts">
// imports
import navBar from './components/bars/navBar.vue'
import { RouterView } from 'vue-router';
///////////////////
</script>
