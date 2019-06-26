<template>
  <div>
    <sun-header></sun-header>
    <profile-nav></profile-nav>
    <el-row>
      <el-col :span="8">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>基本信息</span>
          </div>
          <div v-for="i in attr.length"  :key="i" class="text item">
            {{ attr[i-1] +": " + user[attr[i-1]] }}
          </div>
          <el-button style="float: right; padding: 3px 0" type="text">充值</el-button>

        </el-card>
      </el-col>

      <el-col :span="16">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>我的艺术品列表</span>
          </div>
          <el-table
      :data="myArtworks"
      style="width: 100%">
      <el-table-column
        prop="title"
        label="名称"
        width="180">
      </el-table-column>
      <el-table-column
        prop="timeReleased"
        label="上传日期"
        width="180">
      </el-table-column>
       <el-table-column
      fixed="right"
      label="操作"
      width="120">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="deleteRow(scope.$index, tableData)"
          type="text"
          size="small">
          移除
        </el-button>
      </template>
    </el-table-column>
    </el-table>
          <el-button style="float: right; padding: 3px 0" type="text">充值</el-button>

        </el-card>
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>我的订单列表</span>
          </div>
          <el-table
      :data="myArtworks"
      style="width: 100%">
      <el-table-column
        prop="title"
        label="名称"
        width="180">
      </el-table-column>
      <el-table-column
        prop="timeReleased"
        label="上传日期"
        width="180">
      </el-table-column>
       <el-table-column
      fixed="right"
      label="操作"
      width="120">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="deleteRow(scope.$index, tableData)"
          type="text"
          size="small">
          移除
        </el-button>
      </template>
    </el-table-column>
    </el-table>
        </el-card>
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>我的卖出列表</span>
          </div>
          <el-table
      :data="myArtworks"
      style="width: 100%">
      <el-table-column
        prop="title"
        label="名称"
        width="180">
      </el-table-column>
      <el-table-column
        prop="timeReleased"
        label="上传日期"
        width="180">
      </el-table-column>
       <el-table-column
      fixed="right"
      label="操作"
      width="120">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="deleteRow(scope.$index, tableData)"
          type="text"
          size="small">
          移除
        </el-button>
      </template>
    </el-table-column>
    </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
  .text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }

.box-card{
  margin:20px
}
</style>

<script>
import SunHeader from './base-components/SunHeader'
import ProfileNav from './base-components/ProfileNav'
export default {
  name: 'Profile',
  components: {
    SunHeader,
    ProfileNav
  },
  data () {
    return {
      tocken: localStorage.getItem('tocken'),
      user: {},
      attr: ['name', 'email', 'tel', 'address', 'balance'],
      myArtworks: [
        {

        }
      ],
      myOrders: [
        {

        }
      ],
      soldArtworks: [
        {

        }
      ]
    }
  },
  mounted () {
    this.$store.commit('addFoot', '个人页面')
    this.getUserInfo()
  },
  methods: {
    getUserInfo: function () {
      var that = this
      var form = 'tocken=' + this.tocken
      this.axios.post('/user', form).then(
        function (res) {
          if (res.status === 200) {
            if (res.data.state === 'T') {
              that.user = res.data.message
            } else {
              that.openError(res.data.message)
            }
          }
        })
    },

    deleteRow (index, rows) {
      rows.splice(index, 1)
    },

    openSucceess: function (message) {
      this.$notify({
        title: '成功',
        message: message,
        type: 'success'
      })
    },
    openError: function (message) {
      this.$notify.error({
        title: '错误',
        message: message
      })
    }
  }

}
</script>
