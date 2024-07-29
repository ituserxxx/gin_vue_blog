<template>
  <el-main class="Detail">
    <h1>{{ article.title }}</h1>
    <p>
       <span v-if="article.create_time">
        创建时间：{{ article.create_time }}
      </span>
      <span v-if="article.update_time">
        | 更新时间：{{ article.update_time }}
      </span>
    </p>
    <!-- 访问次数：<span>2002-9-9</span> |
    访问人数：<span>2002-9-9</span>  -->
    <div class="cont">
      <div class="markdown-body" v-html="article.content"></div>
    </div>
    <VueMarkdown v-model="article.content" ></VueMarkdown>
  </el-main>
</template>

<script>
import { articleDetailApi } from "../api/article.js";
import showdown from "showdown";
export default {
  name: "Detail",
  data() {
    return {
      html: "",
      value: "# test markdown",
      article: {
        title: "",
        create_time: "",
        update_time: "",
        content: "",
      },
    };
  },
  created() {
    this.initArticleDetail(this.$route.query.id);
  },
  methods: {
    initArticleDetail(id) {
      this.html = new showdown.Converter();
      articleDetailApi({
        id: Number(id),
      })
        .then((res) => {
          this.article.title = res.data.title;
          this.article.create_time = res.data.create_time;
          this.article.update_time = res.data.update_time;
          this.article.content = this.html.makeHtml(res.data.content);
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>

<style scoped>
.Detail {
  margin-top: 50px;
}
.cont {
  overflow-y: auto;
  box-sizing: border-box;
  overflow-x: hidden;
  text-align: center;
}
.markdown-body {
  text-align: left;
  padding: 0px 18% 15px 18%;
}

.Detail >>> pre {
  padding: 16px;
  overflow: auto;
  font-size: 20px;
  line-height: 1.45;
  background-color: #dadab0;
  border-radius: 3px;
  color: #952442;
  font-weight: bold;

}
</style>

