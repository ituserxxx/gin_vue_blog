import http from '../utils/http'
export function homeApi (data) {
    return http.post('/home',data)
}

