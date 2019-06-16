<template>
    <div>
        <el-row type="flex" justify="space-between" align="middle">
            <el-col class="margin-top" :span="8">
                <img src="../../assets/logo.jpg" height="80px">
            </el-col>
            <el-col class="margin-top" :span="10" >
                <el-form :inline="true" :model="searchInfo" class="demo-form-inline"  >
                    <el-form-item>
                        <el-input v-model="searchInfo.keywords" placeholder="请输入关键词"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">搜索</el-button>
                    </el-form-item>
                </el-form>

            </el-col>
        </el-row>
    </div>
</template>

<style scoped>
.el-form-item {
    margin-bottom: 0px;
}

</style>

<script>
export default {
  name: 'SunSearchbar',
  data () {
    return {
      searchInfo: {
        keywords: ''
      }
    }
  },
  methods: {
    onSubmit () {
      console.log('submit!')
    }
  }
}
</script>
