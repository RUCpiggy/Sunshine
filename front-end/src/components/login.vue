<template>
    <div id="login">
        <el-row>
        <el-button type='success' v-on:click="count++">You clicked me {{ count }} times.</el-button>
        </el-row>
    </div>
</template>

<script>
export default ('login', {
  data: function () {
    this.axios.get('/users/master').then((response) => {
      console.log(response.data)
    })
    console.log('tmp')
    return {
      count: 0
    }
  }
})
</script>
