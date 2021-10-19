import request from '@/utils/request'

export function getTagList(data) {
  return request({
    url: '/tag/list',
    method: 'post',
    data
  })
}

export function editTag(data) {
  return request({
    url: '/tag/update',
    method: 'post',
    data
  })
}
export function addTag(data) {
  return request({
    url: '/tag/add',
    method: 'post',
    data
  })
}

export function delTag(data) {
  return request({
    url: '/tag/del',
    method: 'post',
    data
  })
}
