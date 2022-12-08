module.exports = {
    devServer: {
      proxy: {
        '/blog': {
          target: 'http://127.0.0.1:6008',
          ws: true,
          changeOrigin: true
        },
      
      }
    }
  }