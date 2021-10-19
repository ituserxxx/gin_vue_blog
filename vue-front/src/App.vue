<template>
  <div id="app">
    <HeadContainer/>
    <router-view/>
    <BottomContainer/>
  </div>
</template>
<script>
import HeadContainer from "./components/HeadContainer";
import BottomContainer from "./components/BottomContainer";
export default {
  name: 'App',
  components: {
    HeadContainer,
    BottomContainer
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.el-main {
  background-color: #e9eef3;
  color: #333;
  text-align: center;
}
</style>
