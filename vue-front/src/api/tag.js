import http from '../utils/http'

export function tagListApi (data) {
    return http.post('/tag/list',data)
}

export function tagArticleListApi (data) {
    return http.post('/tag/article/list',data)
}