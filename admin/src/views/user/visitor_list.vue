<template>
  <div class="app-container">
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column label="Ip" prop="ip" />
      <el-table-column label="路由" prop="uri" />
      <el-table-column label="地址" prop="ip_address" />
      <el-table-column label="访问时间" prop="visit_time" />
      <!-- <el-table-column align="right" label="操作">
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
      </el-table-column> -->
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
import { visitorList } from '@/api/user'

export default {
  data() {
    return {
      tableData: [],
      total: 100,
      search: '',
      editTagName: '',
      page: 1
    }
  },
  created() {
    // 初始化数据
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      visitorList({ page: this.page }).then((response) => {
        this.tableData = response.data
        this.listLoading = false
      })
    },
    handleCurrentPageChange(size) {
      console.log(size)
    },
    handleDelete(index, row) {
      const tagInfo = {
        id: row.id
      }
      console.log(tagInfo, '这个接口还没写')
    },
    // 编辑
    handleEdit(index, row) {
      console.log('这个接口还没写')
    }
  }
}
</script>
