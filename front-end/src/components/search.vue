<template>
  <div>
    <sun-header></sun-header>
    <el-row>
    <el-col class="margin-top" :span="24" >
                <el-form :inline="true" :model="searchInfo" class="demo-form-inline"  >
                    <el-form-item>
                        <el-input v-model="searchInfo.title" placeholder="艺术品名称"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-input v-model="searchInfo.description" placeholder="简介"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-input v-model="searchInfo.artist" placeholder="作者名"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="search(1)">搜索</el-button>
                    </el-form-item>
                </el-form>

            </el-col>
    </el-row>

      <div v-if="currentItem >= 1">
      <el-col :span="8" v-for="item in currentItem" :key="item">
        <sun-infobox :artwork="artworks[item-1]"></sun-infobox>
      </el-col>
      </div>
      <el-col>
    <div class="block">
      <span class="demonstration"></span>
        <el-pagination
        layout="prev, pager, next"
        :current-page="page"
        @current-change="handleCurrentChange"
        :total="total">
        </el-pagination>
    </div>
    </el-col>
  </div>
</template>

<script>
import SunHeader from './base-components/SunHeader'
import SunInfobox from './base-components/SunInfobox'
export default {
  components: {
    SunHeader,
    SunInfobox
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
      total: 0,
      page: 1,
      currentItem: 0,
      searchInfo: {
      }
    }
  },
  name: 'Search',

  mounted () {
    this.$store.commit('addFoot', '搜索')
    console.log(this.$store.state.foot)
  },
  methods: {
    search: function (page) {
      var that = this
      this.searchInfo['page'] = page
      this.axios.get('/search', {
        params: this.searchInfo

      })
        .then(
          function (res) {
            if (res.data) {
              const data = res.data
              if (data.state === 'T') {
                that.artworks = data.message
                that.total = data.total
                that.currentItem = data.message.length
                console.log(data)
              }
            }
          })
    },
    handleCurrentChange (page) {
      this.search(page)
    }
  }
}
</script>
