<script setup lang="ts">
import {ref} from 'vue';

import SplashScreen from './components/SplashScreen.vue';
import DashboardLayout from './layouts/DashboardLayout.vue';

const isErrorPage = window.location.pathname.startsWith('/error');
const initialized = ref(isErrorPage);

setTimeout(() => {
  initialized.value = true;
}, 1000);
</script>

<template>
  <el-config-provider>
    <dashboard-layout v-if="initialized"></dashboard-layout>
    <splash-screen v-else></splash-screen>
  </el-config-provider>
</template>
