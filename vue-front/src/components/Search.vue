<template>
  <div id="Search">
    <el-autocomplete
        v-model="state"
        :fetch-suggestions="querySearchAsync"
        placeholder="搜索"
        @select="handleSelect"
    ></el-autocomplete>
  </div>
</template>

<script>
export default {
  name: "Search",
  data() {
    return {
      restaurants: [],
      state: '',
      timeout: null
    };
  },
  methods: {
    loadAll() {
      return [
        {"value": "三全鲜食（北新泾店）", "address": "长宁区新渔路144号"},
        {"value": "Hot honey 首尔炸鸡（仙霞路）", "address": "上海市长宁区淞虹路661号"},
        {"value": "新旺角茶餐厅", "address": "上海市普陀区真北路988号创邑金沙谷6号楼113"},
        {"value": "泷千家(天山西路店)", "address": "天山西路438号"},
      ]
    }
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
  },
  mounted() {
    this.restaurants = this.loadAll();
  }
}
</script>

<style scoped>
#Search{
  /*display:flex;*/
  /*padding-left: 50px;*/
}
</style>