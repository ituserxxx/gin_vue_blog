<template>
  <el-main class="Tag">
    <div class="tag_list">
      <el-badge
        class="item"
        v-for="(val, index) in tag_list"
        :key="index"
        :value="val.article_sum"
        :max="50"
      >
        <el-button type="success" @click="queryTagArticle(val.id)">
          {{ val.tag_name }}
        </el-button>
      </el-badge>
    </div>
    <br />
    <div class="tag_article_list">
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
import { tagListApi, tagArticleListApi } from "../api/tag.js";

export default {
  name: "Tag",

  data() {
    return {
      tag_list: [],
      article_list: [],
      article_page: 1,
    };
  },

  created() {
    this.initTagListApi();
    const id = this.$route.query.id;
    if (id != undefined) {
      this.initTagArticleList();
    }
  },
  watch: {
    $route() {
      this.initTagArticleList();
    },
  },
  methods: {
    initTagListApi() {
      tagListApi({ page: -1 })
        .then((res) => {
          this.tag_list = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    initTagArticleList() {
      const id = this.$route.query.id;
      if (id != undefined) {
        tagArticleListApi({ id: Number(id), page: this.article_page })
          .then((res) => {
            this.article_list = res.data.article_list;
          })
          .catch((err) => {
            console.log(err);
          });
      }
    },
    queryTagArticle(id) {
      this.$router.push({
        path: "/Tag",
        query: {
          id: id,
        },
      });
    },
    gotoDetail(id) {
      this.$router.push({
        path: "/Detail",
        query: {
          id: id,
        },
      });
    },
  },
};
</script>

<style scoped>
.Tag{
    margin-top: 50px;
}
.tag_list {
  margin: 0% 20% 10% 20%;
}
.tag_article_list {
  text-align: left;
}
.item {
  margin-top: 10px;
  margin-right: 40px;
}
.Tag >>> .el-card__body {
  padding-top: 0px;
  padding-bottom: 0px;
}

.tag_article_list li {
  list-style-type: none;
    margin: 0% 20% 0px 20%;
  text-align: center;
  width: 60%;
}
</style>
