<template>
  <div id="app">
    <Header />
    <router-view class="my-2"></router-view>
  </div>
</template>

<script>
import Vue from "vue";

import Header from "./components/Header";
export default {
  name: "app",
  components: {
    Header,
  },
};
</script>

<style lang="scss">
</style>