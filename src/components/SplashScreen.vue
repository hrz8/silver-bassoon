<script setup lang="ts">
import logo from '../assets/images/logo.png';
</script>

<template>
  <div>
    <img
      :src="logo"
      key="logo"
    />
  </div>
</template>

<style scoped>
div {
  display: flex;
  place-content: center;
  place-items: center;
  position: fixed;
  z-index: 999;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

div img {
  transition: all ease-in-out 0.25s;
}
</style>
