<template>
  <div class="app">
    <div class="div_head">
      <h2>Galen test</h2>
    </div>
    <div class="div_daohang">
      <RouterLink to="/json2struct" class="RouterLink" active-class="active_link_daohang">json2struct</RouterLink>
      <br />
      <RouterLink to="/jsonfile2struct" class="RouterLink" active-class="active_link_daohang">jsonfile2struct
      </RouterLink>
      <br />
      <RouterLink to="/urlescap" class="RouterLink" active-class="active_link_daohang">urlEscap</RouterLink>
      <br />
      <RouterLink to="/redisshell" class="RouterLink" active-class="active_link_daohang">redis</RouterLink>
      <br />
      <RouterLink to="/configF" class="RouterLink" active-class="active_link_daohang">config</RouterLink>
    </div>
    <div class="div_content">
      <RouterView></RouterView>
    </div>
  </div>
</template>

<script setup lang="ts" name="app">
import { RouterView, RouterLink } from 'vue-router'
</script>

<style>
.div_head {
  text-align: center;
}

.div_content {
  margin-left: 17%;
}


.div_daohang {
  float: left;
  width: 15%;

  border: 2px solid darkseagreen;
  border-radius: 5px;
}

.active_link_daohang {
  background-color: aqua;
  text-decoration: none;
}

.RouterLink {
  text-decoration: none;
  margin-left: 10px;
}

button {
  margin-left: 5px;
  border-radius: 3px;
}
</style>
