<template>
<div>

<el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal">
  <el-menu-item index="1"><a href="/" >首页</a></el-menu-item>
  <el-menu-item index="2"><a href="/search" >搜索</a></el-menu-item>
  <el-menu-item index="3"><a href="/detail" >详情</a></el-menu-item>
  <el-menu-item index="4"><sun-login>登陆</sun-login></el-menu-item>
  <el-menu-item index="5"><sun-register>注册</sun-register></el-menu-item>
</el-menu>
<div class="line"></div>

</div>
</template>

<script>
import SunLogin from './SunLogin'
import SunRegister from './SunRegister'
export default {
  name: 'SunNav',
  components: {
    SunLogin,
    SunRegister
  },
  data () {
    return {
      activeIndex: '1'
    }
  },
  methods: {
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
