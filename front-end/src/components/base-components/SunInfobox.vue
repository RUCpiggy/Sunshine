<template>
    <div class="infobox">
        <el-row>
            <el-col :span="5">
                <div class="image">
                    <img :src="'http://localhost:8080/api/img/'+ artwork.imageFileName" width=100%>
                </div>
            </el-col>
            <el-col :span="12">
                <p>{{artwork.title}}</p>
                <p>{{artwork.artist}}</p>
            </el-col>
        </el-row>
        <el-row>
            <p class="description" v-html="artwork.description"></p>
        </el-row>
        <el-row>
            <el-col :span="12"><el-button @click="routeTo(artwork.artworkID)">查看</el-button></el-col>
            <el-col :span="12">
                 <chart-button :artworkID="artwork.artworkID"></chart-button>
        </el-col>
        </el-row>
    </div>
</template>

<style scoped>
.infobox{
    margin:20px;
    height: 600px;
}
</style>

<script>
import ChartButton from './ChartButton'
export default {
  name: 'SunInfobox',
  props: ['artwork'],
  components: {
    ChartButton
  },
  data () {
    return {

    }
  },
  methods: {
    routeTo: function (ArtworkID) {
      this.$router.push({path: '/detail', query: { 'ArtworkID': ArtworkID }})
    }
  }
}
</script>
