<template>
    <div>
        <div class="block ">
            <el-carousel height=600px >
            <el-carousel-item v-for="item in 4" :key="item">
            <h3>{{ item }}</h3>

        </el-carousel-item>
        </el-carousel>
        </div>
      </div>
</template>

<style scoped>
.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
  }

.el-carousel__item:nth-child(4n+3) {
  background-color:aqua;
  background-image: url('http://localhost:8080/api/img/102.jpg');
  background-size: cover;
  background-position: center center;
  }

.el-carousel__item:nth-child(2n+1) {

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
      artwork: {
        title: '',
        description: '',
        artworkID: '',
        artist: '',
        imageFileName: ''
      },
      artworks: ['', '', '', '']
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
          for (var i = 0; i < res.data.message.length; i++) {
            that.artwork.artist = res.data.message[i].Artist
            that.artwork.title = res.data.message[i].Title
            that.artwork.description = res.data.message[i].Description
            that.artwork.artworkID = res.data.message[i].ArtworkID
            that.artwork.imageFileName = res.data.message[i].ImageFileName
            that.artworks[i] = that.artwork
          }
        })
      console.log(this.artworks)
    }

  },
  watch: {

  }

}

</script>
