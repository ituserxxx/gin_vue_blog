// module.exports = {
//     devServer: {
//         proxy: {
//             '/blog': {
//                 target: 'http://127.0.0.1:1002', // 后台端口地址
//                 changeOrigin: true, // 将基于名称的虚拟托管网站的选项，如果不配置，请求会报404  如果接口跨域，需要进行这个参数配置
//                 ws: true, // true / false，是否代理websockets
//                 secure: false, // 如果是https接口，需要配置这个参数
//                 pathRewrite: {
//                     '^/': 'blog' // pathRewrite是使用proxy进行代理时，对请求路径进行重定向以匹配到正确的请求地址
//                 }
//             }
//         }
//     },
// }
