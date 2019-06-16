<template>
<div>
<el-button type="text" @click="dialogFormVisible = true"><slot></slot></el-button>
  <el-dialog title="用户注册" :visible.sync="dialogFormVisible">
    <el-form :model="form" :rules="rules" ref="form">
      <el-form-item label="用户名" :label-width="formLabelWidth" prop="name">
        <el-input v-model="form.name" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="密码" :label-width="formLabelWidth" prop="inputPassword">
        <el-input type="password" v-model="form.inputPassword" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="确认密码" :label-width="formLabelWidth" prop="checkPassword">
        <el-input type="password" v-model="form.checkPassword" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="邮箱" :label-width="formLabelWidth" prop="email">
        <el-input v-model="form.email" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="电话号码" :label-width="formLabelWidth" prop="tel">
        <el-input v-model="form.tel" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="地址" :label-width="formLabelWidth" prop="address">
        <el-input v-model="form.address" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="验证码" :label-width="formLabelWidth" prop="captcha">
        <el-row>
           <el-col :span="12">
               <el-input v-model="form.captcha" auto-complete="off"></el-input>
            </el-col>
            <el-col :span="12">
              <div @click="refreshCode">
                <sun-captcha :identifyCode="identifyCode"></sun-captcha>
              </div>
            </el-col>

        </el-row>

      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button @click="resetForm('form')">重置</el-button>
      <el-button type="primary" @click="postRegisterInfo('form')">确 定</el-button>
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
import crypto from 'crypto'
import SunCaptcha from './SunCaptcha'
export default {
  name: 'SunRegister',
  components: {
    SunCaptcha
  },
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
    var vaildatePassword2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.form.inputPassword) {
        callback(new Error('两次输入密码不一致'))
      } else {
        callback()
      }
    }
    var vaildateEmail = (rule, value, callback) => {
      var reg = /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
      if (value === '') {
        callback(new Error('请输入邮箱'))
      } else if (!reg.test(value)) {
        callback(new Error('邮箱格式不正确'))
      } else {
        callback()
      }
    }
    var vaildateTel = (rule, value, callback) => {
      var reg = /^1(3|4|5|7|8)\d{9}$/
      if (value === '') {
        callback(new Error('请输入手机号'))
      } else if (!reg.test(value)) {
        callback(new Error('手机号格式不正确'))
      } else {
        callback()
      }
    }
    var vaildateAddress = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入收货地址'))
      } else {
        callback()
      }
    }
    var vaildateCatptha = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入验证码'))
      } else if (value !== this.identifyCode) {
        callback(new Error('验证码输入错误'))
      } else {
        callback()
      }
    }
    return {
      dialogFormVisible: false,
      form: {
        name: '',
        password: '',
        inputPassword: '',
        checkPassword: '',
        email: '',
        tel: '',
        captcha: '',
        address: ''
      },
      identifyCode: '',
      identifyCodes: 'ABC1234567890',
      rules: {
        name: [
          {validator: vaildateName, trigger: 'blur', required: true}
        ],
        inputPassword: [
          {validator: vaildatePassword, trigger: 'blur', required: true}
        ],
        checkPassword: [
          {validator: vaildatePassword2, trigger: 'blur', required: true}
        ],
        email: [
          {validator: vaildateEmail, trigger: 'blur', required: true}
        ],
        tel: [
          {validator: vaildateTel, trigger: 'blur', required: true}
        ],
        address: [
          {validator: vaildateAddress, trigger: 'blur', required: true}
        ],
        captcha: [
          {validator: vaildateCatptha, trigger: 'blur', required: true}
        ]
      },
      formLabelWidth: '120px'
    }
  },
  mounted () {
    // init captcha
    this.identifyCode = ''
    this.makeCode(this.identifyCodes, 4)
  },

  methods: {
    postRegisterInfo: function (form) {
      this.form.password = this.getSha1(this.form.inputPassword)
      this.$refs[form].validate((valid) => {
        if (valid) {
          var that = this
          this.axios.post('/register', this.form).then(
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
    resetForm: function (form) {
      this.$refs[form].resetFields()
    },
    refreshCode: function () {
      this.identifyCode = ''
      this.makeCode(this.identifyCode, 4)
    },
    makeCode: function (o, l) {
      for (let i = 0; i < l; i++) {
        this.identifyCode += this.identifyCodes[this.randomNum(0, this.identifyCodes.length)]
      }
      console.log(this.identifyCode)
    },
    randomNum: function (min, max) {
      return Math.floor(Math.random() * (max - min) + min)
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
    },
    getSha1: function (date) {
      const hash = crypto.createHash('sha1')
      hash.update(date)
      return hash.digest('hex')
    }
  }
}
</script>
