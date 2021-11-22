import request from '@/utils/request'

export function getUploadToken(data) {
  return request({
    url: '/upload/token',
    method: 'post',
    data
  })
}
