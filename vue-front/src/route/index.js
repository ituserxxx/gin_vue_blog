
import { createRouter,createWebHistory } from 'vue-router'
import Home from '../view/Home.vue'
import About from '../view/About.vue'
import Search from '../view/Search.vue'
import Tag from '../view/Tag.vue'
import Detail from '../view/Detail.vue'
import LeaveMessage from '../view/LeaveMessage.vue'
const routerHistory = createWebHistory()
const routes = [
    {
        path:'/',
        name:'/',
        component:Home,
       
    },
    {
        path:'/Home',
        name:'Home',
        component:Home,
       
    },
    {
        path:'/About',
        name:'About',
        component:About
    },
    {
        path:'/Tag',
        name:'Tag',
        component:Tag
    },
    {
        path:'/Search',
        name:'Search',
        component:Search
    },
    {
        path:'/Detail',
        name:'Detail',
        component:Detail
    },
    {
        path:'/LeaveMessage',
        name:'LeaveMessage',
        component:LeaveMessage
    }
]
const router = createRouter({
    history: routerHistory,
    routes
})
export default router
