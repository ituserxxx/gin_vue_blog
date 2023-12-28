<template>
  <el-main class="About">
    <div class="div_content" style="witdh: 600px">
      <el-timeline-item placement="top">
        <el-card>
          <h2>Vtian welcome to here~</h2>
          <h2>It is a coder blog</h2>
        </el-card>
      </el-timeline-item>
   
      <el-timeline-item placement="top">
        <el-card>
        <h2>个人技能～</h2>
          <el-progress :text-inside="true" :stroke-width="25"  
          v-for="(jnv, index) in jn"
          :key="index"
          :percentage="jnv.socre"  :color="jnv.color">{{ jnv.name }}</el-progress>
        </el-card>
      </el-timeline-item>

      <el-timeline-item placement="top">
        <el-card>
          <h2>开源项目</h2>
        <h4> <a href="https://github.com/ituserxxx/gin_vue_blog" target="_blank">当前博客源码</a></h4>
        <h4> <a href="https://github.com/ituserxxx/bun_api_frame" target="_blank">nodejs bun 基础框架</a></h4>
        <h4><a href="https://github.com/ituserxxx/client_test_tool" target="_blank">常用客户端工具</a></h4> 
        <h4><a href="https://github.com/ituserxxx/sse_demo" target="_blank">sse 通信 demo</a></h4> 
        <h4><a href="https://github.com/ituserxxx/gin-web-api-frame" target="_blank">基于gin基础api开发框架</a></h4> 
        <h4><a href="https://github.com/ituserxxx/rust_axum_web_api_demo" target="_blank">基于rust基础api开发框架</a></h4> 
        </el-card>
      </el-timeline-item>
    </div>
  </el-main>
</template>

<script>
export default {
  name: "About",
  data(){
    return{
        jn:[
          {
            name:"php",
            color:"#2D3540",
            socre:70,
          },
          {
            name:"golang",
            color:"#37404D",
            socre:80,
          },
          {
            name:"mysql",
            color:"#63526E",
            socre:75,
          },
          {
            name:"docker",
            color:"#9C86BA",
            socre:70,
          },
          {
            name:"redis",
            color:"#AE684E",
            socre:65,
          },
          {
            name:"vue",
            color:"#716161",
            socre:60,
          },
          {
            name:"rust",
            color:"#3C3D3F",
            socre:50,
          },
          {
            name:"python",
            color:"#661F33",
            socre:70,
          }
        ]
    }
  },
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