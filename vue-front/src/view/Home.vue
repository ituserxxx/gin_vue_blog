<template>
  <el-main class="Home">
    <ArticleListContainer></ArticleListContainer>
  </el-main>
</template>
<script>
import ArticleListContainer from "../components/ArticleListContainer";
export default {
  name: "Home",
  components: {
    ArticleListContainer,
  }
  
};
</script>
<style scoped>
.Home{
  background-color: #CCFF99
}
</style>
