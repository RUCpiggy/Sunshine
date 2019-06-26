<template>
    <div>
        <el-row>
            <el-col :span="5">
                <div class="image">
                    <img :src="'http://localhost:8080/api/img/'+ artwork.imageFileName" width=100%>
                </div>
            </el-col>
            <el-col :span="5">
                <p>{{artwork.title}}</p>
                <p>{{artwork.artist}}</p>
            </el-col>
        </el-row>
        <el-row>
            <p class="description" v-html="artwork.description"></p>
        </el-row>
        <el-row>
            <el-col :span="12"><el-button>查看</el-button></el-col>
            <el-col :span="12"><el-button>添加到购物车</el-button></el-col>
        </el-row>
    </div>
</template>

<script>
export default {
  name: 'SunInfobox',
  props: ['artwork'],
  data () {
    return {

    }
  },
  mounted () {
    console.log(this.artwork)
  },
  methods: {

  }
}
</script>
