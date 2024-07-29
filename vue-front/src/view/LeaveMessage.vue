<template>
  <div class="class-leave-message-list">
    <el-timeline>
      <el-timeline-item v-for="(msg_item, index) in msg_list" :key="index" :timestamp="msg_item.timestamp">
        <p style="font-size: larger;">{{ msg_item.content }}</p>
      </el-timeline-item>
    </el-timeline>

    <el-pagination background :pager-count="5" layout="prev, pager, next" :total="total" :hide-on-single-page="true"
      @current-change="handleCurrentChange">
    </el-pagination>
  </div>

</template>
<script>
  import { leaveMessageListApi } from "../api/leave_message.js";
  export default {
    name: "LeaveMessage",
    data() {
      return {
        current_page: 1,
        total: 0,
        msg_list: [
          // {
          //   content: 'asdfadfadf阿迪斯发斯蒂芬阿第三方阿斯蒂芬asdfadfadf阿迪斯发斯蒂芬阿第三方阿斯蒂芬阿凡达山东发斯蒂芬嘎嘎嘎和她和她人哈你打刚发您多发归纳代发给',
          //   timestamp: '2018-04-12 20:46',
          // },
          // {
          //   content: 'asdfadfadf阿迪斯发斯蒂芬阿第三方阿斯蒂芬阿凡达山东发斯蒂芬嘎嘎嘎和她和她人哈你打刚发您多发归纳代发给 ',
          //   timestamp: '2018-04-12 20:46',
          // },
        ],
      }
    },
    created() {
      this.initMsgList();
    },
    methods: {
      initMsgList() {
        this.randMsgList();
      },
      handleCurrentChange(val) {
        this.current_page = val;
        this.randMsgList()
      },
      randMsgList() {
        leaveMessageListApi({
          page: this.current_page,
        })
          .then((res) => {
            this.msg_list = res.data.msg_list;
            this.total = res.data.total;
          })
          .catch((err) => {
            console.log(err);
          });

      },
    },
  };
</script>
<style scoped>
  .Page-tool {
  text-align: center;
}
  .class-leave-message-list {
    margin-top: 80px;
    max-width: 80%;
    margin-left: 20%;
  }
</style>