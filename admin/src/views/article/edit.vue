<template>
  <div class="app-container">
    <el-form ref="form" :model="form" :rules="rules" label-width="120px">
      <el-input v-show="false" v-model="form.article_id" />
      <el-form-item label="标题" prop="title">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="选择标签">
        <el-checkbox-group v-model="form.tag_list">
          <el-checkbox
            v-for="(tag, index) in init_tag_list"
            :key="index"
            :label="tag.id"
            name="tag"
          >
            <el-tag type="success">{{ tag.tag_name }}</el-tag>
          </el-checkbox>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="创建时间">
        <el-date-picker
          v-model="form.create_time"
          type="datetime"
          placeholder="选择日期时间"
        />
      </el-form-item>
      <el-form-item label="最近更新时间">
        <el-date-picker
          v-model="form.update_time"
          type="datetime"
          placeholder="选择日期时间"
        />
      </el-form-item>
      <el-form-item label="内容">
        <mavon-editor v-model="form.content" />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          :loading="isLoading"
          @click="onSubmit"
        >提交</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { editArticle, detailArticle } from '@/api/article'
import { getTagList } from '@/api/tag'

export default {
  data() {
    return {
      isLoading: false,
      init_tag_list: [],
      form: {
        title: '',
        tag_list: [],
        content: '',
        create_time: '',
        update_time: '',
        article_id: 0,
        html: ''
      },
      rules: {
        title: [
          { required: true, message: '请输入活动名称', trigger: 'blur' },
          {
            min: 3,
            max: 50,
            message: '长度在 3 到 20 个字符',
            trigger: 'blur'
          }
        ],
        tag_list: [
          {
            type: 'array',
            required: true,
            message: '请至少选择一个标签',
            trigger: 'change'
          }
        ]
      }
    }
  },
  created() {
    this.isLoading = true
    this.initTagList()
    this.getDetail(this.$route.query.id)
  },
  methods: {
    initTagList() {
      getTagList({ page: 1 }).then((response) => {
        this.init_tag_list = response.data
      })
    },
    getDetail(id) {
      detailArticle({ article_id: Number(id) }).then((response) => {
        this.isLoading = false
        this.form.article_id = response.data.id
        this.form.title = response.data.title
        if (response.data.tag_list != null) {
          this.form.tag_list = response.data.tag_list
        }
        this.form.content = response.data.content
        this.form.create_time = response.data.create_time
        this.form.update_time = response.data.update_time
      })
    },
    onSubmit() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.isLoading = true
          editArticle(this.form).then((response) => {
            this.isLoading = false
            if (response.data !== '') {
              this.$message({
                showClose: true,
                message: response.data,
                type: 'primary'
              })
            } else {
              this.$message({
                showClose: true,
                message: '成功',
                type: 'success'
              })
            }
          })
        } else {
          return false
        }
      })
    }
  }
}
</script>

<style scoped>
.line {
  text-align: center;
}
</style>

