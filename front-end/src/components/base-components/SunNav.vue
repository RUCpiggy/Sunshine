<template>
<div>

<el-menu  class="el-menu-demo" mode="horizontal">
  <el-menu-item ><a href="/" >首页</a></el-menu-item>
  <el-menu-item ><a @click="routeTo('/search')" >搜索</a></el-menu-item>
  <el-menu-item ><a @click="routeTo('/detail')"  >详情</a></el-menu-item>
  <el-menu-item v-if="loginStatus=='0'"><sun-login v-bind:message="loginStatus" v-on:listenLoginStatus="status">登陆</sun-login></el-menu-item>
  <el-menu-item v-if="loginStatus=='0'"><sun-register v-bind:message="loginStatus" v-on:listenLoginStatus="status">注册</sun-register></el-menu-item>
    <el-menu-item v-if="loginStatus=='1'" ><a @click="routeTo('/profile')">个人页面</a></el-menu-item>
  <el-menu-item v-if="loginStatus=='1'" ><a @click="logout">登出</a></el-menu-item>
</el-menu>
<div class="line"></div>

</div>
</template>

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
      loginStatus: '0'
    }
  },
  mounted () {
    if (localStorage.tocken != null) {
      this.loginStatus = '1'
    }
  },
  methods: {
    routeTo: function (path) {
      this.$router.push({path: path})
    },
    status: function (status) {
      this.loginStatus = status
    },
    logout: function () {
      localStorage.removeItem('tocken')
      this.loginStatus = '0'
      this.openSucceess('登出成功')
    },
    openSucceess: function (message) {
      this.$notify({
        title: '成功',
        message: message,
        type: 'success'
      })
    }
  }

}
</script>
