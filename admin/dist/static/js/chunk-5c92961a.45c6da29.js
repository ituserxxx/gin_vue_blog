(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5c92961a"],{2423:function(t,e,a){"use strict";a.d(e,"f",(function(){return r})),a.d(e,"a",(function(){return i})),a.d(e,"c",(function(){return u})),a.d(e,"e",(function(){return c})),a.d(e,"d",(function(){return l})),a.d(e,"g",(function(){return s})),a.d(e,"b",(function(){return o}));var n=a("b775");function r(t){return Object(n["a"])({url:"/article/list",method:"post",data:t})}function i(t){return Object(n["a"])({url:"/article/add",method:"post",data:t})}function u(t){return Object(n["a"])({url:"/article/detail",method:"post",data:t})}function c(t){return Object(n["a"])({url:"/article/update",method:"post",data:t})}function l(t){return Object(n["a"])({url:"/article/draft",method:"post",data:t})}function s(t){return Object(n["a"])({url:"/article/publish",method:"post",data:t})}function o(t){return Object(n["a"])({url:"/article/delete",method:"post",data:t})}},eb84:function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"app-container"},[a("el-row",[a("el-button",{attrs:{type:"primary"},on:{click:t.handleAdd}},[t._v("新增文章")])],1),a("el-table",{staticStyle:{width:"100%"},attrs:{data:t.tableData,border:""}},[a("el-table-column",{attrs:{label:"Id",prop:"id"}}),a("el-table-column",{attrs:{label:"标题",prop:"title"}}),a("el-table-column",{attrs:{label:"创建时间",prop:"create_time"}}),a("el-table-column",{attrs:{label:"更新时间",prop:"update_time"}}),a("el-table-column",{attrs:{label:"状态",prop:"status"},scopedSlots:t._u([{key:"default",fn:function(e){return[1===e.row.status?a("el-tag",{attrs:{type:"success"}},[t._v(" 已发布")]):t._e(),2===e.row.status?a("el-tag",{attrs:{type:"warning"}},[t._v(" 草稿 ")]):t._e()]}}])}),a("el-table-column",{attrs:{label:"操作",align:"right"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{size:"mini"},on:{click:function(a){return t.handleEdit(e.$index,e.row)}}},[t._v("编辑")]),1==e.row.status?a("el-button",{attrs:{size:"mini",type:"primary"},on:{click:function(a){return t.handleSaveDraft(e.$index,e.row)}}},[t._v("存草稿")]):t._e(),2==e.row.status?a("el-button",{attrs:{size:"medium",type:"success"},on:{click:function(a){return t.handlePublishArticle(e.$index,e.row)}}},[t._v("发布")]):t._e(),a("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(a){return t.handleDelete(e.$index,e.row)}}},[t._v("删除")])]}}])})],1),a("el-pagination",{attrs:{background:"",layout:"prev, pager, next",total:t.total},on:{"current-change":t.handleCurrentPageChange}})],1)},r=[],i=a("2423"),u={data:function(){return{tableData:[],total:100,search:""}},created:function(){this.fetchData()},methods:{fetchData:function(){var t=this;this.listLoading=!0,Object(i["f"])({page:this.page}).then((function(e){t.tableData=e.data,t.listLoading=!1}))},handleCurrentPageChange:function(t){var e=this;Object(i["f"])({page:t}).then((function(t){e.tableData=t.data,e.listLoading=!1}))},handleEdit:function(t,e){this.$router.push({path:"/article/edit",query:{id:e.id}})},handleAdd:function(){this.$router.push({path:"/article/add"})},handleSaveDraft:function(t,e){var a=this,n={id:e.id};Object(i["d"])(n).then((function(t){a.$message({type:"success",message:t.msg?t.msg:"成功"}),a.fetchData()}))},handlePublishArticle:function(t,e){var a=this,n={id:e.id};Object(i["g"])(n).then((function(t){a.$message({type:"success",message:t.msg?t.msg:"成功"}),a.fetchData()}))},handleDelete:function(t,e){var a=this,n={article_id:e.id};Object(i["b"])(n).then((function(t){a.$message({type:"success",message:t.msg?t.msg:"成功"}),a.fetchData()}))}}},c=u,l=a("2877"),s=Object(l["a"])(c,n,r,!1,null,null,null);e["default"]=s.exports}}]);