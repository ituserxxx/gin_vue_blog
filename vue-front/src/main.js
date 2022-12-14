import { createApp } from 'vue'
import router from './route'
import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'
import App from './App.vue'
const app = createApp(App)



// 渲染markdown内容
import showdown from 'showdown';
app.use(showdown)

// import VueMarkdown from 'vue-markdown'
// app.use(VueMarkdown)


app.use(router)
app.use(ElementPlus)
app.mount('#app')

