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

    <el-autocomplete
        v-model="state"
        maxlength="30"
        :fetch-suggestions="querySearchAsync"
        placeholder="搜索一下"
        clearable="true">

<!--        <template slot-scope="{ item }">-->
<!--          <div class="name">{{ item.value }}</div>-->
<!--          <span class="addr">{{ item.address }}</span>-->
<!--        </template>-->
    </el-autocomplete>
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
        {title: "关于", url: "/About"},
      ],
      restaurants: [],
      state: '',
      timeout:  null
    };
  },
  methods: {
    loadAll() {
      return [
        { "value": "三全鲜食（北新泾店）", "address": "长宁区新渔路144号" },
        { "value": "Hot honey 首尔炸鸡（仙霞路）", "address": "上海市长宁区淞虹路661号" },
        { "value": "新旺角茶餐厅", "address": "上海市普陀区真北路988号创邑金沙谷6号楼113" },
        { "value": "泷千家(天山西路店)", "address": "天山西路438号" },
      ];
    },
    querySearchAsync(queryString, cb) {
      var restaurants = this.restaurants;
      var results = queryString ? restaurants.filter(this.createStateFilter(queryString)) : restaurants;

      clearTimeout(this.timeout);
      this.timeout = setTimeout(() => {
        cb(results);
      }, 3000 * Math.random());
    },
    createStateFilter(queryString) {
      return (state) => {
        return (state.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0);
      };
    },
    handleSelect(item) {
      console.log(item);
    }
  },
  mounted() {
    this.restaurants = this.loadAll();
  }

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

.HeadContainer >>> .el-autocomplete {
  position: relative;
  display: inline-block;
  left: 80%;
  z-index: 50;
}


</style>