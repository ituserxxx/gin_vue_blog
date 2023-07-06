<template>
  <div class="HeadContainer">
    <el-header>
      <el-menu
          class="el-menu-demo"
          mode="horizontal"
          active-text-color="#ffd04b"
          background-color="#545c64"
          text-color="#fff"
      >
        <el-menu-item
            target="_blank"
            v-for="(val, index) in headCategoryList"
            :key="index"
            index="index"
        >
          <router-link :to="val.url">
            <span class="menu-title">{{ val.title }}</span>
          </router-link>
        </el-menu-item>
      </el-menu>
    </el-header>
  </div>
</template>

<script>

export default {
  name: "HeadContainer",

  data() {
    return {
      headCategoryList: [
        {title: "首页", url: "/Home"},
        {title: "标签", url: "/Tag"},
        {title: "搜索", url: "/Search"},
        {title: "关于", url: "/About"},
      ],
    };
  },


};
</script>
<style scoped>

.el-header {
  background-color: #ffcc99;
  color: #ccffcc;
  text-align: center;
  line-height: 60px;
  display: flex;
  justify-content: center;
  position: fixed;
  z-index: 10;
  left: 8px;
  right: 8px;
  top: 0;
}

.menu-title {
  font-size: 25px;
}

</style>