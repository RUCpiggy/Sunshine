<template>
  <div>
    <div class="block ">
      <el-carousel height=600px >
        <el-carousel-item v-for="item in 4" :key="item" :style="'background-image: url(http://localhost:8080/api/img/'+artworks[item-1].ImageFileName+');'">
          <div>
            <p class="title">{{ artworks[item-1].Title }}</p>
            <p v-html="artworks[item-1].Description"></p>
            <p class="more" @click="routeTo(artworks[item-1].ArtworkID)">learn more</p>
          </div>

        </el-carousel-item>
      </el-carousel>
    </div>
    <div>

    </div>
  </div>
</template>

<style scoped>
.el-carousel__item div {
  color: white;
  font-size: 14px;
  opacity: 1;
  margin: 0;
  margin-top: 100px
  }

.el-carousel__item {
  background-color:black;
  background-size: cover;
  background-position: center center;
  }
.title{
  font-size:30px;
}

.more:hover {
  cursor:pointer;
}

</style>

<script>
import SunNav from './base-components/SunNav'
export default {
  name: 'Index',
  components: {
    SunNav
  },
  data () {
    return {
      bannerHeight: 500,
      artworks: [
        {
          ImageFileName: ''
        },
        {
          ImageFileName: ''
        },
        {
          ImageFileName: ''
        },
        {
          ImageFileName: ''
        }
      ]
    }
  },
  mounted () {
    this.getHotArtworks()
  },

  methods: {
    getHotArtworks: function () {
      var that = this
      this.axios.get('/hot').then(
        function (res) {
          if (res.data) {
            const data = res.data.message
            that.artworks = data
          }
        })
    },
    routeTo: function (ArtworkID) {
      this.$router.push({path: '/detail', query: { 'ArtworkID': ArtworkID }})
    }

  },

  watch: {

  }

}

</script>
