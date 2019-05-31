<template>
<div>

<el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
  <el-menu-item index="1"><a href="/" target="_blank">首页</a></el-menu-item>
  <el-menu-item index="2"><a href="/" target="_blank">搜索</a></el-menu-item>
  <el-menu-item index="3"><a href="/" target="_blank">详情</a></el-menu-item>
  <el-menu-item index="4"><Login></Login></el-menu-item>
  <el-menu-item index="5"><a href="/" target="_blank">注册</a></el-menu-item>
</el-menu>
<div class="line"></div>
<el-row>
  <el-col class="margin-top" :span="6">
    <img src="../../assets/logo.jpg" height="70px">
  </el-col>
  <el-col class="margin-top" :span="6">

  </el-col>
</el-row>
</div>
</template>

<script>
import Login from './Login'
export default {
  name: 'SunNav',
  components: {Login},
  data () {
    return {
      activeIndex: '1'
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>

<style scoped>

.el-menu-demo
{
  display: flex;
  justify-content: flex-end;
}
a {
  text-decoration: none;
}

.margin-top {
  margin-top:10px;
}

</style>
