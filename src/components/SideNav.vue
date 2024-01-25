<script setup lang="ts">
import {HomeFilled, ShoppingBag} from '@element-plus/icons-vue';

import logo from '../assets/images/logo-simple.png';

const iconSize = 26;
</script>

<template>
  <el-menu style="height: 100%">
    <router-link to="/">
      <el-menu-item index="logo">
        <img :src="logo" />
      </el-menu-item>
    </router-link>

    <router-link to="/">
      <el-menu-item
        index="home"
        class="sidenav-item"
      >
        <el-icon
          :size="iconSize"
          :color="'primary'"
        >
          <HomeFilled />
        </el-icon>
      </el-menu-item>
    </router-link>

    <router-link to="/orders">
      <el-menu-item
        index="orders"
        class="sidenav-item"
      >
        <el-icon
          :size="iconSize"
          :color="'primary'"
        >
          <ShoppingBag />
        </el-icon>
      </el-menu-item>
    </router-link>
  </el-menu>
</template>

<style scoped>
img {
  padding: 10px !important;
  width: 50px;
}

.sidenav-item {
  display: flex;
  flex-direction: column;
  line-height: 1;
  align-items: center;
  justify-content: center;
  height: 50px;
}
</style>
