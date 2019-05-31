<template>
  <div id="app">
  <SunNav></SunNav>
    <router-view/>
  </div>
</template>

<script>
import SunNav from './components/base-components/SunNav'
export default {
  name: 'App',
  components: {SunNav}
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
