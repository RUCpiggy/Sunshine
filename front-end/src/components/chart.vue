<template>
    <div>
        <sun-header></sun-header>
        <el-col :span="8" v-for="item in artworks.length " :key="item">
          <chart-infobox :artwork="artworks[item-1]"></chart-infobox>
        </el-col>
        <el-col>total: {{total}}</el-col>
        <el-button @click="pay">下单</el-button>
    </div>
</template>

<script>
import SunHeader from './base-components/SunHeader'
import ChartInfobox from './base-components/ChartInfobox'
export default {
  components: {
    SunHeader,
    ChartInfobox
  },
  data () {
    return {
      artworks: [
        {ImageFileName: ''},
        {ImageFileName: ''},
        {ImageFileName: ''},
        {ImageFileName: ''},
        {ImageFileName: ''},
        {ImageFileName: ''},
        {ImageFileName: ''},
        {ImageFileName: ''},
        {ImageFileName: ''},
        {ImageFileName: ''}
      ],
      total: 0

    }
  },
  mounted () {
    this.$store.commit('addFoot', '购物车')
    this.getArtworks()
  },
  methods: {
    getArtworks: function () {
      var that = this
      this.axios.get('/chart', {
        params: {
          'tocken': localStorage.getItem('tocken')
        }}).then(
        function (res) {
          if (res.data) {
            const data = res.data.message
            that.artworks = data
          }
        })
    },
    pay: function () {

    }

  },
  watch: {
    artworks: function () {
      for (var i = 0; i < this.artworks.length; i++) {
        this.total = this.total + parseInt(this.artworks[i].price)
      }
    }
  }
}
</script>
