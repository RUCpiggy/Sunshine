<template>
    <div>
        <el-button size="medium" @click="addToChart()">添加到购物车</el-button>
    </div>
</template>

<script>
export default {
  props: ['artworkID'],
  methods: {
    addToChart: function () {
      var that = this
      var form = 'ArtworkID=' + this.artworkID + '&tocken=' + localStorage.getItem('tocken')

      this.axios.post('/chart', form).then(
        function (res) {
          if (res.status === 200) {
            if (res.data.state === 'T') {
              that.openSucceess(res.data.message)
            } else {
              that.openError(res.data.message)
            }
          }
        })
    },
    openError: function (message) {
      this.$notify.error({
        title: '错误',
        message: message
      })
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
