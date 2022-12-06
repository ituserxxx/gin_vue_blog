<template>
  <div class="app-container">
    <el-row>
      <el-button type="primary" @click="handleAdd">添加新标签</el-button>
    </el-row>
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column label="Id" prop="id" />
      <el-table-column label="标签名称" prop="tag_name" />
      <el-table-column label="文章数量" prop="article_sum" />
      <el-table-column align="right" label="操作">
        <template slot-scope="scope">
          <el-button
            type="text"
            @click="handleEdit(scope.$index, scope.row)"
          >编辑</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
          >Delete</el-button>
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
import { getTagList, editTag, addTag, delTag } from '@/api/tag'

export default {
  data() {
    return {
      tableData: [],
      total: 100,
      search: '',
      editTagName: '',
      curr_page: 1
    }
  },
  created() {
    // 初始化数据
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getTagList({ page: this.curr_page }).then((response) => {
        this.tableData = response.data
        this.listLoading = false
      })
    },
    handleCurrentPageChange(page) {
      this.curr_page = page
      this.fetchData()
    },
    handleDelete(index, row) {
      const tagInfo = {
        id: row.id
      }
      delTag(tagInfo).then((res) => {
        this.tableData.splice(index, 1)
      })
    },
    // 新增
    handleAdd() {
      this.$prompt('', '请输入标签名称', {
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      })
        .then(({ value }) => {
          if (value === null || value === '') {
            this.$message({
              type: 'info',
              message: '不能为空'
            })
          }
          const tagInfo = {
            tag_name: value
          }
          addTag(tagInfo).then((response) => {
            this.$message({
              type: 'success',
              message: '成功'
            })
            // 刷新表格
            getTagList({ page: 1 }).then((response) => {
              this.tableData = response.data
              this.listLoading = false
            })
          })
        })
        .catch(() => {})
    },
    // 编辑
    handleEdit(index, row) {
      this.$prompt('', '请输入标签名称', {
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      })
        .then(({ value }) => {
          if (value === null || value === '') {
            this.$message({
              type: 'info',
              message: '不能为空'
            })
          }
          const tagInfo = {
            id: row.id,
            tag_name: value
          }
          editTag(tagInfo).then((response) => {
            this.$message({
              type: 'success',
              message: '成功'
            })
            this.tableData[index].tag_name = value
          })
        })
        .catch(() => {})
    }
  }
}
</script>
