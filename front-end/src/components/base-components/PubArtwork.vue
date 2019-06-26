<template>
  <div>
<el-button type="text" @click="dialogFormVisible = true"><slot></slot></el-button>
  <el-dialog title="发布艺术品" :visible.sync="dialogFormVisible">
    <el-form :model="form" :rules="rules" ref="form">
      <el-form-item label="名称" :label-width="formLabelWidth" prop="title">
        <el-input v-model="form.title" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="作者" :label-width="formLabelWidth" prop="artist">
        <el-input v-model="form.artist" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="简介" :label-width="formLabelWidth" prop="description">
        <el-input v-model="form.description" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="年份" :label-width="formLabelWidth" prop="yearOfWork">
        <el-input v-model="form.yearOfWork" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="流派" :label-width="formLabelWidth" prop="genre">
        <el-input v-model="form.genre" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="宽度（cm）" :label-width="formLabelWidth" prop="width">
        <el-input v-model="form.width" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="长度（cm）" :label-width="formLabelWidth" prop="height">
        <el-input v-model="form.height" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="价格" :label-width="formLabelWidth" prop="price">
        <el-input v-model="form.price" auto-complete="off"></el-input>
      </el-form-item>

      <el-upload
  class="avatar-uploader"
  action="http://localhost:8080/api/img"
  :show-file-list="false"
  :on-success="handleAvatarSuccess"
  :before-upload="beforeAvatarUpload">
  <img v-if="imageUrl" :src="imageUrl" class="avatar">
  <i v-else class="el-icon-plus avatar-uploader-icon"></i>
</el-upload>
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
.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
</style>

<script>
import crypto from 'crypto'
export default {
  name: 'SunLogin',
  data () {
    var vaildateYear = (rule, value, callback) => {
      var reg = /^-?[1-9]\d*$/
      if (value === '') {
        callback(new Error('年份不能为空'))
      } else if (!reg.test(value)) {
        callback(new Error('年份必须为整数'))
      } else {
        callback()
      }
    }
    var vaildateWidth = (rule, value, callback) => {
      var reg = /^(0|[1-9][0-9]*)(\.\d+)?$/
      if (value === '') {
        callback(new Error('宽度不能为空'))
      } else if (!reg.test(value)) {
        callback(new Error('宽度必须为正数'))
      } else {
        callback()
      }
    }
    var vaildateHeight = (rule, value, callback) => {
      var reg = /^(0|[1-9][0-9]*)(\.\d+)?$/
      if (value === '') {
        callback(new Error('宽度不能为空'))
      } else if (!reg.test(value)) {
        callback(new Error('宽度必须为正数'))
      } else {
        callback()
      }
    }
    var vaildatePrice = (rule, value, callback) => {
      var reg = /^[1-9]\d*$/
      if (value === '') {
        callback(new Error('价格不能为空'))
      } else if (!reg.test(value)) {
        callback(new Error('价格必须为正整数'))
      } else {
        callback()
      }
    }

    return {
      imageUrl: '',
      dialogFormVisible: false,
      form: {
        title: '',
        artist: '',
        description: '',
        yearOfWork: '',
        genre: '',
        width: '',
        height: '',
        price: ''

      },
      rules: {
        title: [
          {trigger: 'blur', required: true, message: '名称不能为空'}
        ],
        artist: [
          {trigger: 'blur', required: true, message: '作者不能为空'}
        ],
        description: [
          {trigger: 'blur', required: true, message: '简介不能为空'}
        ],
        yearOfWork: [
          {validator: vaildateYear, trigger: 'blur', required: true}
        ],
        genre: [
          {trigger: 'blur', required: true, message: '流派不能为空'}
        ],
        width: [
          {validator: vaildateWidth, trigger: 'blur', required: true}
        ],
        height: [
          {validator: vaildateHeight, trigger: 'blur', required: true}
        ],
        price: [
          {validator: vaildatePrice, trigger: 'blur', required: true}
        ]
      },
      formLabelWidth: '120px'
    }
  },
  methods: {
    postLoginInfo: function (form) {
      this.form['imageFileName'] = this.imageFileName
      this.form['tocken'] = localStorage.getItem('tocken')
      this.$refs[form].validate((valid) => {
        if (valid) {
          var that = this
          this.axios.post('/artwork', this.form).then(
            function (res) {
              if (res.status === 200) {
                if (res.data.state === 'T') {
                  that.openSucceess(res.data.message)
                  that.dialogFormVisible = false
                  console.log(res.data)
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
    handleAvatarSuccess (res, file) {
      this.imageUrl = URL.createObjectURL(file.raw)
      this.imageFileName = file.name
    },
    beforeAvatarUpload (file) {
      const isJPG = file.type === 'image/jpeg'
      const isLt2M = file.size / 1024 / 1024 < 2

      if (!isJPG) {
        this.$message.error('上传图片只能是 JPG 格式!')
      }
      if (!isLt2M) {
        this.$message.error('上传图片大小不能超过 2MB!')
      }
      return isJPG && isLt2M
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
