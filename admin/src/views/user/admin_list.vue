<template>
  <div class="app-container">
    <el-row>
      <el-button type="primary" @click="handleAdd">添加管理员</el-button>
    </el-row>
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column label="Id" prop="id" />
      <el-table-column label="登录账户" prop="username" />
      <el-table-column label="显示昵称" prop="nickname" />
      <el-table-column label="介绍" prop="introduction" />
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
  </div>
</template>

<script>
import { adminList } from '@/api/user'
export default {
  data() {
    return {
      tableData: [],
      total: 100,
      search: '',
      editTagName: ''
    }
  },
  created() {
    // 初始化数据
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      adminList({ page: 1 }).then((response) => {
        this.tableData = response.data
        this.listLoading = false
      })
    },
    // 删除
    handleDelete(index, row) {
      const tagInfo = {
        id: row.id
      }
      console.log(tagInfo, '这个接口还没写')
    },
    // 新增
    handleAdd() {
      console.log('这个接口还没写')
    },
    // 编辑
    handleEdit(index, row) {
      console.log('这个接口还没写')
    }
  }
}
</script>
