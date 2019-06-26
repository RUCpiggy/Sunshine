<template>
  <div>
    <div class="block ">
      <el-carousel height=600px >
        <el-carousel-item v-for="item in 4" :key="item" :style="'background-image: url(http://localhost:8080/api/img/'+artworks[item-1].imageFileName+');'">
          <div>
            <p class="title">{{ artworks[item-1].title }}</p>
            <p v-html="artworks[item-1].description"></p>
            <p class="more" @click="routeTo(artworks[item-1].artworkID)">learn more</p>
          </div>
        </el-carousel-item>
      </el-carousel>
      <el-row>
        <el-col :span="8" v-for="item in 3" :key="item">
          <sun-newcard :artwork="newArtworks[item-1]"></sun-newcard>
        </el-col>
      </el-row>
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
import SunNewcard from './base-components/SunNewcard'

export default {
  name: 'Index',
  components: {
    SunNav,
    SunNewcard
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
      ],
      newArtworks: [
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
    this.getNewArtworks()
    this.$store.commit('addFoot', '首页')
    console.log(this.newArtworks)
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
    getNewArtworks: function () {
      var that = this
      this.axios.get('/new').then(
        function (res) {
          if (res.data) {
            const data = res.data.message
            that.newArtworks = data
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
