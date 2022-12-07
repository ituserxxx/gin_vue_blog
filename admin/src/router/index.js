import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '首页面板', icon: 'dashboard' }
    }]
  },

  {
    path: '/article',
    component: Layout,
    redirect: '/article/list',
    name: 'article',
    meta: { title: '', icon: 'el-icon-s-help' },
    children: [
      {
        path: 'list',
        name: 'list',
        component: () => import('@/views/article/list'),
        meta: { title: '文章列表', icon: 'table' }
      },
      {
        path: 'add',
        name: 'add',
        component: () => import('@/views/article/add'),
        meta: { title: '新增', icon: 'form' },
        hidden: true
      },
      {
        path: 'edit',
        name: 'edit',
        component: () => import('@/views/article/edit'),
        meta: { title: '编辑', icon: 'form' },
        hidden: true
      }
    ]
  },
  {
    path: '/tag',
    component: Layout,
    redirect: '/tag/list',
    name: 'tag',
    meta: { title: '标签管理', icon: 'el-icon-s-help' },
    children: [
      {
        path: 'list',
        name: 'Table',
        component: () => import('@/views/tag/list'),
        meta: { title: '标签列表', icon: 'table' }
      }
    ]
  },
  // {
  //   path: '/img',
  //   component: Layout,
  //   redirect: '/img/img',
  //   name: 'tag',
  //   meta: { title: '图片管理', icon: 'el-icon-s-help' },
  //   children: [
  //     {
  //       path: 'img',
  //       name: 'Img',
  //       component: () => import('@/views/img/add'),
  //       meta: { title: '添加图片', icon: 'table' }
  //     }
  //   ]
  // },
  {
    path: '/user',
    component: Layout,
    redirect: '/user/visitor_list',
    name: 'user',
    meta: { title: '用户管理', icon: 'el-icon-s-help' },
    children: [
      {
        path: 'visitor_list',
        name: 'Table',
        component: () => import('@/views/user/visitor_list'),
        meta: { title: '访客列表', icon: 'table' }
      },
      {
        path: 'admin_list',
        name: 'Table',
        component: () => import('@/views/user/admin_list'),
        meta: { title: '管理员列表', icon: 'table' }
      }
    ]
  },

  // {
  //   path: 'external-link',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'https://panjiachen.github.io/vue-element-admin-site/#/',
  //       meta: { title: 'External Link', icon: 'link' }
  //     }
  //   ]
  // },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
