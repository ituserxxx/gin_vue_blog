var http = require('http')
// var querystring = require('querystring')
http.createServer(function (request, response) {
    response.setHeader('Access-Control-Allow-Origin', '*')
    // 允许的header类型
    response.setHeader('Access-Control-Allow-Headers', 'content-type')
    // 跨域允许的请求方式
    response.setHeader('Access-Control-Allow-Methods', 'DELETE,PUT,POST,GET,OPTIONS')
    const url = request.url
    let data = {}

    // 2.注册data事件接收数据（每当收到一段表单提交的数据，该方法会执行一次）
    request.on('data', function (chunk) {
        data += chunk
    })

    request.on('end', function () {
        data = decodeURI(data)
        // console.log(data)
        // var dataObject = querystring.parse(data)
        // console.log(dataObject)
    })

    console.log('请求url=>' + request.url)
    if (url === '/blog/article/list') {
        data = postList()//此处就是接口返回的数据
    }
    if (url === '/blog/tag/list') {
        data = tagList()//此处就是接口返回的数据
    }

    // 内容类型: text/plain
    response.writeHead(200, {'Content-Type': 'application/json; charset=UTF-8'})
    // 发送响应数据
    response.end(JSON.stringify({
        code: 20000,
        msg: true,
        data: data
    }))
}).listen(6008)

function postList() {
    return {
        "article_list": [{
            "id": 116,
            "title": "wsl2 原生 Linux 安装 Docker 方式  (非windows 桌面版 docker )"
        }, {
            "id": 109,
            "title": "nginx 部署网站"
        }, {
            "id": 108,
            "title": "linux 查看当前监听的所有端口"
        }, {
            "id": 107,
            "title": "mysql count(*)、count(1)、count(id)，count(name)比较"
        }],
        "total": 112
    }
}

function tagList() {
    return [{
        "id": 19,
        "tag_name": "音乐",
        "article_sum": 2
    }, {
        "id": 1,
        "tag_name": "Git",
        "article_sum": 2
    }]
}