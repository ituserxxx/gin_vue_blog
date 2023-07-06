import http from '../utils/http'
export function articleDetailApi (data) {
    return http.post('/article/detail',data)
}

export function articleListApi (data) {
    return http.post('/article/list',data)
}

export function searchArticleApi (data) {
    return http.post('/search/article',data)
}
