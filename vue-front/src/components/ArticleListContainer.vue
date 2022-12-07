<template>
  <div class="ArticleListContainer">
    <router-view />
    <div class="list">
      <div style="height:20px"></div>
      <el-timeline-item
        v-for="(val, index) in article_list"
        :key="index"
        :hide-timestamp="true"
        @click="gotoDetail(val.id)"
      >
        <el-card>
          <h1>{{ val.title }}</h1>
        </el-card>
      </el-timeline-item>
    </div>
    <div v-if="total > 0" class="Page-tool">
      <el-pagination
        background
        :pager-count="5"
        layout="prev, pager, next"
        :total="total"
        :hide-on-single-page="true"
         @current-change="handleCurrentChange"
      >
      </el-pagination>
    </div>
  </div>
</template>

<script>
import { articleListApi } from "../api/article.js";
export default {
  name: "ContentContainer",
  data() {
    return {
      current_page: 1,
      total: 0,
      article_list: [],
    };
  },
  created() {
    this.initArticleList();
  },
  methods: {
    initArticleList() {
      this.randArticleList();
    },
    handleCurrentChange(val) {
      this.current_page = val;
      this.randArticleList()
    },
    gotoDetail(id) {
      this.$router.push({
        path: "/Detail",
        query: {
          id: id,
        },
      });
    },

    randArticleList() {
      articleListApi({
        page: this.current_page,
      })
        .then((res) => {
          this.article_list = res.data.article_list;
          this.total = res.data.total;
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>


<style  scoped>
.Page-tool {
  text-align: center;
}
.ArticleListContainer {
  margin-top: 50px;
  background-color: #e9eef3;
  color: #333;
  
}
.ArticleListContainer li {
  list-style-type: none;
}
.list li {
  margin: 0% 20% 0px 20%;
  text-align: center;
  width: 60%;
}
</style>