<template>
  <el-main class="Search">
    <div style="margin-top: 50px; width: 20% ;margin-left: 40%;">
      <el-input placeholder="搜索一下吧～～" v-model="input_search">
      </el-input>
    </div>
    <el-empty :image-size="200" v-if="article_length === 0"></el-empty>
    <div class="tag_article_list" v-if="article_length > 0">
      <el-timeline-item
          v-for="(activity, index) in article_list"
          :key="index"
          :hide-timestamp="true"
          @click="gotoDetail(activity.id)"
      >
        <el-card>
          <h4>{{ activity.title }}</h4>
        </el-card>
      </el-timeline-item>
    </div>

  </el-main>
</template>

<script>

import {searchArticleApi} from "../api/article";

export default {
  name: "Search",
  data() {
    return {
      input_search: "",
      article_list: [],
      article_page: 1,
      article_length: 0
    }
  },
  watch: {
    input_search(newQuestion, oldQuestion) {
      if (newQuestion.length > 0 && newQuestion.length !== oldQuestion.length) {
        this.searchArticle(newQuestion)
      }
    },
  },
  created() {

  },
  methods: {
    gotoDetail(id) {
      this.$router.push({
        path: "/Detail",
        query: {
          id: id,
        },
      });
    },
    searchArticle(inputValue) {
      searchArticleApi({content: inputValue,})
          .then((res) => {
              this.article_list = res.data.article_list;
              this.total = res.data.total;
              this.article_length = res.data.total;
          })
          .catch((err) => {
            console.log(err);
          });
    },
  },
};
</script>
<style>
.el-select .el-input {
  width: 130px;
}

.input-with-select .el-input-group__prepend {
  background-color: #fff;
}
</style>