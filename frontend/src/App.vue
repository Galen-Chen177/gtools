<template>
  <div class="app">
    <div class="center">
      <h2>Galen test</h2>
    </div>
    <div class="daohang">
      <RouterLink to="/json2struct" active-class="daohanglink">json2struct</RouterLink>
      <br />
      <RouterLink to="/jsonfile2struct" active-class="daohanglink">jsonfile2struct</RouterLink>
      <br />
      <RouterLink to="/urlescap" active-class="daohanglink">urlEscap</RouterLink>
    </div>
    <div class="content">
      <RouterView></RouterView>
    </div>
  </div>
</template>

<script setup lang="ts" name="app">
import { RouterView, RouterLink } from 'vue-router'
</script>

<style>
nav {
  float: left;
  width: 15%;
}

.content {
  margin-left: 17%;
}

.center {
  text-align: center;
}

.daohang {
  float: left;
  width: 15%;
}

.daohanglink {
  background-color: aqua;
}
</style>
