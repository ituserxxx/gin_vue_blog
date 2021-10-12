import request from '@/utils/request'

export function getArticleList(data) {
  return request({
    url: '/article/list',
    method: 'post',
    data
  })
}

export function createArticle(data) {
  return request({
    url: '/article/add',
    method: 'post',
    data
  })
}

export function detailArticle(data) {
  return request({
    url: '/article/detail',
    method: 'post',
    data: data
  })
}

export function editArticle(data) {
  return request({
    url: '/article/update',
    method: 'post',
    data: data
  })
}
// 存入草稿
export function draftArticle(data) {
  return request({
    url: '/article/draft',
    method: 'post',
    data: data
  })
}
// 发布
export function publishArticle(data) {
  return request({
    url: '/article/publish',
    method: 'post',
    data: data
  })
}
// 发布
export function deleteArticle(data) {
  return request({
    url: '/article/delete',
    method: 'post',
    data: data
  })
}
