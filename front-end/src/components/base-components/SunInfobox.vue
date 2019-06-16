<template>
    <div>
        <el-row>
            <el-col :span="12">
                <div class="image">
                    <img :src="'http://localhost:8080/api/img/'+ artwork.ImageFileName">
                </div>
            </el-col>
            <el-col :span="12">
                <p>{{artwork.Title}}</p>
                <p>{{artwork.Artist}}</p>
            </el-col>
        </el-row>
        <el-row>
            <p class="description" v-html="artwork.Description"></p>
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
  data () {
    return {

    }
  }
}
</script>
