<template>
  <div>
    <sun-header></sun-header>
    <button @click="tmpclick">tmp</button>
    <el-row>
      <el-col :span="18">
        <el-row>
            <h2>{{artwork.Title}}</h2>
        </el-row>
        <el-row>
          <p>{{artwork.Artist}}</p>
        </el-row>
        <el-row type="flex" justify="space-between"  align="middle">
          <el-col :span="10">
            <div class="image">
              <img :src="'http://localhost:8080/api/img/'+ artwork.ImageFileName">
            </div>
          </el-col>
          <el-col :span="14">
            <el-row>
              <em>Price</em><span> : {{artwork.Price}}$</span>
            </el-row>
            <el-row>
              <el-button size="medium">添加到购物车</el-button>
            </el-row>
            <el-row type="flex" justify="center">
            <el-col :span="18">
              <p>Product Details</p>
              <el-divider></el-divider>
              <el-row ><el-col :span="10">Date</el-col><el-col :span="14">{{artwork.YearOfWork}}</el-col></el-row>
              <el-divider></el-divider>
              <el-row ><el-col :span="10">Genre</el-col><el-col :span="14">{{artwork.Genre}}</el-col></el-row>
              <el-divider></el-divider>
              <el-row ><el-col :span="10">Dimensions</el-col><el-col :span="14">{{artwork.Width}}cm X {{artwork.Height}}cm</el-col></el-row>
              <el-divider></el-divider>
              <el-row ><el-col :span="10">Hot</el-col><el-col :span="14">{{artwork.View}} views</el-col></el-row>
              <el-divider></el-divider>
            </el-col>
            </el-row>

          </el-col>
        </el-row>
        <el-row>
          <p class="description" v-html="artwork.Description"></p>
        </el-row>
      </el-col>
      <el-col :span="6">
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>

img {
  width: 100%;
}

.el-divider {
  margin-top: 5px;
  margin-bottom: 5px;
}
.image {
  margin: 20px;
}
.description {
  text-align: justify;
}
.description:before {
  content: "\2002\2002";
}

</style>

<script>
import SunHeader from './base-components/SunHeader'
export default {
  name: 'Detail',
  components: {
    SunHeader
  },
  data () {
    return {
      artwork: {
        ArtworkID: this.$route.query.ArtworkID,
        ImageFileName: '446.jpg'
      }
    }
  },
  mounted () {
    this.getArtworkInfo()
  },
  methods: {
    getArtworkInfo: function () {
      var that = this
      this.axios.get('/detail', {
        params: {
          'ArtworkID': this.artwork.ArtworkID
        }
      })
        .then(
          function (res) {
            if (res.data) {
              const data = res.data.message[0]
              that.artwork = data
              console.log(data)
            }
          })
    },
    tmpclick: function () {
      console.log(this.artwork.Description)
    }
  }
}
</script>
