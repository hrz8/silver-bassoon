<script setup lang="ts">
import {useRoute} from 'vue-router';

const route = useRoute();
</script>

<template>
  <div class="wrapper">
    <div class="container">
      <el-row>
        <el-col class="text-center">
          <h4>Here are some reasons why you may be seeing this page</h4>

          <ul class="list-unstyled">
            <li v-if="route.params.errorType === 'not-found'">
              <p>This page does not exist.</p>
              <router-link :to="'/'">
                <el-button
                  size="small"
                  type="success"
                  >Back to home</el-button
                >
              </router-link>
            </li>
          </ul>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<style scoped>
.wrapper {
  position: fixed;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  display: flex;
  place-content: center;
  place-items: center;
  background: #4987e4;
  z-index: 999;
}

.container {
  border: 1px solid #eee;
  border-radius: 10px;
  padding: 40px;
  margin: 20px;
  box-shadow: 0 2px 8px 0 rgba(0, 0, 0, 0.3);
  background: #fff;
}
</style>
