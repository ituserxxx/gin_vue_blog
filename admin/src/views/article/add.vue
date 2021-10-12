<template>
  <div class="app-container">
    <el-form ref="form" :model="form" :rules="rules" label-width="120px">
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
      <el-form-item label="内容">
        <mavon-editor v-model="form.content" />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          :loading="isLoading"
          @click="publish"
        >提交</el-button>
        <el-button
          type="warning"
          :loading="save_draft_isLoading"
          @click="save_draft"
        >存草稿</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { createArticle } from '@/api/article'
import { getTagList } from '@/api/tag'
export default {
  data() {
    return {
      init_tag_list: [],
      isLoading: false,
      save_draft_isLoading: false,
      form: {
        status: 2, // 1-发布，2-草稿，3-下架
        title: '',
        tag_list: [],
        desc: '',
        content: ''
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
    this.initTagList()
  },
  methods: {
    initTagList() {
      getTagList({ page: -1 }).then((response) => {
        this.init_tag_list = response.data
      })
    },
    publish() {
      this.form.status = 1
      this.submit()
    },
    save_draft() {
      this.save_draft_isLoading = true
      this.form.status = 2
      this.submit()
    },
    submit() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.isLoading = true
          createArticle(this.form).then((response) => {
            this.tableData = response.data.id
            this.isLoading = false
            this.save_draft_isLoading = false
            if (response.msg === '') {
              this.$message({
                showClose: true,
                message: '操作成功',
                type: 'success'
              })
              this.$router.go(0)
            } else {
              this.$message({
                showClose: true,
                message: response.msg,
                type: 'warning'
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

