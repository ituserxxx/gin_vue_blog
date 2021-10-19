<template>
  <div class="app-container">
    <el-row>
      <el-button type="primary" @click="handleAdd">新增文章</el-button>
    </el-row>

    <el-table :data="tableData" border style="width: 100%">
      <el-table-column label="Id" prop="id" />
      <el-table-column label="标题" prop="title" />
      <el-table-column label="创建时间" prop="create_time" />
      <el-table-column label="更新时间" prop="update_time" />
      <el-table-column label="状态" prop="status">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 1" type="success">
            已发布</el-tag>
          <el-tag v-if="scope.row.status === 2" type="warning">
            草稿
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="right">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="handleEdit(scope.$index, scope.row)"
          >编辑</el-button>

          <el-button
            v-if="scope.row.status == 1"
            size="mini"
            type="primary"
            @click="handleSaveDraft(scope.$index, scope.row)"
          >存草稿</el-button>

          <el-button
            v-if="scope.row.status == 2"
            size="medium"
            type="success"
            @click="handlePublishArticle(scope.$index, scope.row)"
          >发布</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      background
      layout="prev, pager, next"
      :total="total"
      @current-change="handleCurrentPageChange"
    />
  </div>
</template>

<script>
import {
  getArticleList,
  draftArticle,
  publishArticle,
  deleteArticle
} from '@/api/article'

export default {
  data() {
    return {
      tableData: [],
      total: 100,
      search: ''
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getArticleList({ page: this.page }).then((response) => {
        this.tableData = response.data
        this.listLoading = false
      })
    },
    handleCurrentPageChange(page) {
      getArticleList({ page: page }).then((response) => {
        this.tableData = response.data
        this.listLoading = false
      })
    },
    // 编辑跳转
    handleEdit(index, row) {
      // 通过push进行跳转
      this.$router.push({
        path: '/article/edit',
        query: { id: row.id }
      })
    },
    // 新增跳转
    handleAdd() {
      // 通过push进行跳转
      this.$router.push({
        path: '/article/add'
      })
    },
    // 存入草稿
    handleSaveDraft(index, row) {
      const articleInfo = {
        id: row.id
      }
      draftArticle(articleInfo).then((response) => {
        this.$message({
          type: 'success',
          message: response.msg ? response.msg : '成功'
        })
        this.fetchData()
      })
    },
    handlePublishArticle(index, row) {
      const articleInfo = {
        id: row.id
      }
      publishArticle(articleInfo).then((response) => {
        this.$message({
          type: 'success',
          message: response.msg ? response.msg : '成功'
        })
        this.fetchData()
      })
    },
    handleDelete(index, row) {
      const articleInfo = {
        article_id: row.id
      }
      deleteArticle(articleInfo).then((response) => {
        this.$message({
          type: 'success',
          message: response.msg ? response.msg : '成功'
        })
        this.fetchData()
      })
    }
  }
}
</script>
