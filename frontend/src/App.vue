<template>
  <Navbar v-if="isLoggedIn" />

  <div class="container">
    <router-view :key="$route.fullPath" />
  </div>
</template>

<script>
import Navbar from './components/ui/Navbar.vue';

export default {
  name: 'App',
  components: {
    Navbar
  },
  mounted() {
    this.$store.dispatch("setUser", { "user": JSON.parse(localStorage.getItem("user")) })
  },
  computed: {
    isLoggedIn() {
      let user = this.$store.getters["getUser"];
      if (user) {
        return true;
      } else {
        return false;
      }
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;
}
</style>
