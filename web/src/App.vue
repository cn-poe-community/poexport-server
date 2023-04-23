<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/item">翻译</RouterLink>
        <RouterLink to="/query">查询</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
nav>*:nth-child(n+2){
  margin-left: 10px;
}
nav{
  margin:20px auto;
}
</style>
