import http from '../utils/http'
export function leaveMessageListApi (data) {
    return http.post('/leaveMessage/list',data)
}

