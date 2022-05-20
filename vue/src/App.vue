<template>
  <div id="app">
    <navbar />
    <b-container>
      <router-view />
    </b-container>
  </div>
</template>

<script>
import Navbar from '@/views/layout/Navbar.vue';

export default {
  components: { Navbar },
  data() {},
};
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  // text-align: center;
  // color: #2c3e50;
  // margin-top: 60px;
}
</style>
