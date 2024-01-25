<script setup lang="ts">
import _ from 'lodash';
import {computed, ref} from 'vue';

import SideNav from '../components/SideNav.vue';

const isErrorPage = window.location.pathname.startsWith('/error');
const loading = ref(!isErrorPage);
const loadingText = computed(() =>
  _.sample([
    'Shuffling reality cards...',
    'Dancing with binary stars...',
    'Decrypting the meaning of life...',
    'Assembling parallel universes...',
    'Juggling quantum particles...',
    'Unraveling the fabric of time...',
    'Syncing with cosmic vibrations...',
    'Randomizing existence...',
    'Whispering to parallel dimensions...',
    'Syncing with the cosmic heartbeat...',
    'Decoding the language of galaxies...',
    "...at least you're not on hold...",
    "It's not you. It's me.",
    'Riding the cosmic waves...',
    'Chasing photons in a quantum maze...',
    'Harmonizing with the strings of existence...',
    'Quantum leaping through dimensions...',
    'Deconstructing the cosmic jigsaw puzzle...',
    'Navigating the quantum sea...',
    'Tuning in to the cosmic symphony...',
    'Diving into the pool of infinite possibilities...',
    'Sculpting time like a cosmic artist...',
    'Riding the photon waves...',
    'Chasing the elusive neutrino...',
    'Quantum entangling reality...',
    'Blurring the lines of existence...',
  ]),
);

setTimeout(() => {
  loading.value = false;
}, 2000);
</script>

<template>
  <el-container
    v-loading="loading"
    :element-loading-text="loadingText"
    class="h-100"
  >
    <el-aside class="w-auto position-fixed h-100">
      <side-nav></side-nav>
    </el-aside>
    <el-main class="h-100 p-3 ml-sidenav">
      <el-scrollbar max-height="100vh">
        <router-view class="content"></router-view>
      </el-scrollbar>
    </el-main>
  </el-container>
</template>

<style scoped>
.ml-sidenav {
  margin-left: 90px;
}

.content {
  min-height: 100vh;
}
</style>
