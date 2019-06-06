<template>
  <div>
<el-button type="text" @click="dialogFormVisible = true"><slot></slot></el-button>
  <el-dialog title="用户登陆" :visible.sync="dialogFormVisible">
    <el-form :model="form" :rules="rules" ref="form">
      <el-form-item label="用户名" :label-width="formLabelWidth" prop="name">
        <el-input v-model="form.name" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="密码" :label-width="formLabelWidth" prop="password">
        <el-input type="password" v-model="form.password" auto-complete="off"></el-input>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" @click="postLoginInfo('form')">确 定</el-button>
    </div>
  </el-dialog>

  </div>
</template>

<style scoped>
.el-form-item{
  max-width: 500px;
}
</style>

<script>
export default {
  name: 'SunLogin',
  data () {
    var vaildateName = (rule, value, callback) => {
      var reg = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,}$/
      if (value === '') {
        callback(new Error('请输入用户名'))
      } else if (value.length < 6) {
        callback(new Error('用户名至少为6位'))
      } else if (!reg.test(value)) {
        callback(new Error('用户名不能为纯数字或纯字母'))
      } else {
        callback()
      }
    }
    var vaildatePassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else if (value === this.form.name) {
        callback(new Error('密码不能与用户名相同'))
      } else if (value.length < 6) {
        callback(new Error('密码至少为6位'))
      } else {
        if (this.form.checkPassword !== '') {
          this.$refs.form.validateField('checkPassword')
        }
        callback()
      }
    }
    return {
      dialogFormVisible: false,
      form: {
        name: '',
        password: ''
      },
      rules: {
        name: [
          {validator: vaildateName, trigger: 'blur', required: true}
        ],
        password: [
          {validator: vaildatePassword, trigger: 'blur', required: true}
        ]
      },
      formLabelWidth: '120px'
    }
  },
  methods: {
    postLoginInfo: function (form) {
      this.$refs[form].validate((valid) => {
        if (valid) {
          var that = this
          this.axios.post('/login', this.form).then(
            function (res) {
              if (res.status === 200) {
                if (res.data.state === 'T') {
                  that.openSucceess(res.data.message)
                  that.dialogFormVisible = false
                } else {
                  that.openError(res.data.message)
                }
              }
            })
        } else {
          this.openError('请检查所有信息是否合法')
          return false
        }
      })
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
