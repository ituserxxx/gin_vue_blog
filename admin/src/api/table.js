import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admin/table/list',
    method: 'get',
    params
  })
}

export function getArticleList(page) {
  return request({
    url: '/admin/table/list',
    method: 'post',
    data: {
      page: page
    }
  })
}

export function createArticle(data) {
  return request({
    url: '/admin/article/create',
    method: 'post',
    data
  })
}
