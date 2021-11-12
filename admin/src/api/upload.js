import request from '@/utils/request'

export function getUploadToken(data) {
  return request({
    url: '/upload/toekn',
    method: 'post',
    data
  })
}