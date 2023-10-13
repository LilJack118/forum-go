<template>
  <div class="container">
    <router-view :key="$route.fullPath" />
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {},
  mounted() {
    this.$store.dispatch("setUser", { "user": JSON.parse(localStorage.getItem("user")) })
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
