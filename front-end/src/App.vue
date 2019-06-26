<template>
  <div id="app">
  <sun-nav></sun-nav>
  <el-row type="flex" justify="center"  >
    <el-col :span="22"><router-view/></el-col>
  </el-row>

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
