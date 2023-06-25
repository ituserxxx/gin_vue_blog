<template>
  <el-main class="About">
    <div class="div_content" style="witdh: 600px">
      <el-timeline-item placement="top">
        <el-card>
          <h4>Vtian welcome to here~</h4>
        </el-card>
      </el-timeline-item>
      <el-timeline-item placement="top">
        <el-card>
          <h4>It is a phpgoer blog</h4>
        </el-card>
      </el-timeline-item>
      <el-card >
        <el-progress type="circle" :percentage="0" :stroke-width="10" ></el-progress>
        <el-progress type="circle" :percentage="25"></el-progress>
        <el-progress type="circle" :percentage="100" status="success"></el-progress>
        <el-progress type="circle" :percentage="70" status="warning"></el-progress>
        <el-progress type="circle" :percentage="50" status="exception"></el-progress>
      </el-card>
      <el-timeline-item hide-timestamp="true">
        <el-card>
          <el-progress :text-inside="true" :stroke-width="22" :percentage="70">php</el-progress>
          <el-progress
              :text-inside="true"
              :stroke-width="22"
              :percentage="80"
              status="success"
          >redis
          </el-progress
          >
          <el-progress
              :text-inside="true"
              :stroke-width="22"
              :percentage="80"
              status="warning"
          >mysql
          </el-progress
          >
          <el-progress
              :text-inside="true"
              :stroke-width="22"
              :percentage="50"
              status="exception"
          >golang
          </el-progress
          >
        </el-card>
      </el-timeline-item>
    </div>
  </el-main>
</template>

<script>
export default {
  name: "About",
};
</script>

<style scoped>
.div_content {
  margin-top: 50px;
}

.div_content li {
  list-style-type: none;
  margin: 0% 20% 0px 20%;
  text-align: center;
  width: 60%;
}

</style>