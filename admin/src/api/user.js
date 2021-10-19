import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/user/login',
    method: 'post',
    data
  })
}

export function getInfo(data) {
  return request({
    url: '/user/info',
    method: 'post',
    data
  })
}

export function logout(data) {
  return request({
    url: '/user/logout',
    method: 'post',
    data
  })
}
// 管理员列表
export function adminList(data) {
  return request({
    url: '/list',
    method: 'post',
    data
  })
}

// 访客列表
export function visitorList(data) {
  return request({
    url: '/visitor/list',
    method: 'post',
    data
  })
}

//访客编辑
export function editUser(data) {
  return request({
    url: '/visitor/visitorEdit',
    method: 'post',
    data
  })
}

//访客删除
export function deleteUser(data) {
  return request({
    url: '/visitor/delete',
    method: 'post',
    data
  })
}
